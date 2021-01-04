//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package inverted

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/adapters/repos/db/docid"
	"github.com/semi-technologies/weaviate/adapters/repos/db/helpers"
	"github.com/semi-technologies/weaviate/adapters/repos/db/notimplemented"
	"github.com/semi-technologies/weaviate/adapters/repos/db/propertyspecific"
	"github.com/semi-technologies/weaviate/adapters/repos/db/storobj"
	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/usecases/traverser"
	bolt "go.etcd.io/bbolt"
)

type Searcher struct {
	db            *bolt.DB
	schema        schema.Schema
	rowCache      *RowCacher
	classSearcher ClassSearcher // to allow recursive searches on ref-props
	propIndices   propertyspecific.Indices
	deletedDocIDs DeletedDocIDChecker
}

type DeletedDocIDChecker interface {
	Contains(id uint64) bool
}

func NewSearcher(db *bolt.DB, schema schema.Schema,
	rowCache *RowCacher, propIndices propertyspecific.Indices,
	classSearcher ClassSearcher, deletedDocIDs DeletedDocIDChecker) *Searcher {
	return &Searcher{
		db:            db,
		schema:        schema,
		rowCache:      rowCache,
		propIndices:   propIndices,
		classSearcher: classSearcher,
		deletedDocIDs: deletedDocIDs,
	}
}

// Object returns a list of full objects
func (f *Searcher) Object(ctx context.Context, limit int,
	filter *filters.LocalFilter, underscore traverser.UnderscoreProperties,
	className schema.ClassName) ([]*storobj.Object, error) {
	pv, err := f.extractPropValuePair(filter.Root, className)
	if err != nil {
		return nil, err
	}

	var out []*storobj.Object
	if err := f.db.View(func(tx *bolt.Tx) error {
		if err := pv.fetchDocIDs(tx, f, limit); err != nil {
			return errors.Wrap(err, "fetch doc ids for prop/value pair")
		}

		pointers, err := pv.mergeDocIDs()
		if err != nil {
			return errors.Wrap(err, "merge doc ids by operator")
		}

		// cutoff if required, e.g. after merging unlimted filters
		if len(pointers.docIDs) > limit {
			pointers.docIDs = pointers.docIDs[:limit]
		}

		res, err := docid.ObjectsInTx(tx, pointers.IDs())
		if err != nil {
			return errors.Wrap(err, "resolve doc ids to objects")
		}

		out = res
		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "object filter search bolt view tx")
	}

	return out, nil
}

// DocIDs is similar to Objects, but does not actually resolve the docIDs to
// full objects. Instead it returns the pure object id pointers. They can then
// be used in a secondary index (e.g. vector index)
//
// DocID queries does not contain a limit by design, as we won't know if the limit
// wouldn't remove the item that is most important for the follow up query.
// Imagine the user sets the limit to 1 and the follow-up is a vector search.
// If we already limited the allowList to 1, the vector search would be
// pointless, as only the first element would be allowed, regardless of which
// had the shortest distance
func (f *Searcher) DocIDs(ctx context.Context, filter *filters.LocalFilter,
	underscore traverser.UnderscoreProperties, className schema.ClassName) (helpers.AllowList, error) {
	pv, err := f.extractPropValuePair(filter.Root, className)
	if err != nil {
		return nil, err
	}

	if err := f.db.View(func(tx *bolt.Tx) error {
		if err := pv.fetchDocIDs(tx, f, -1); err != nil {
			return errors.Wrap(err, "fetch doc ids for prop/value pair")
		}

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "doc id filter search bolt view tx")
	}

	pointers, err := pv.mergeDocIDs()
	if err != nil {
		return nil, errors.Wrap(err, "merge doc ids by operator")
	}

	out := make(helpers.AllowList, len(pointers.docIDs))
	for _, p := range pointers.docIDs {
		out.Insert(p.id)
	}

	return out, nil
}

func (fs *Searcher) parseInvertedIndexRow(id, in []byte, limit int,
	hasFrequency bool) (docPointers, error) {
	out := docPointers{
		checksum: make([]byte, 8),
	}

	// 0 is a non-existing row, 8 is one that only contains a checksum, but no content
	if len(in) == 0 || len(in) == 8 {
		return out, nil
	}

	r := bytes.NewReader(in)
	if _, err := r.Read(out.checksum); err != nil {
		return out, errors.Wrap(err, "read checksum")
	}

	// only use cache on unlimited searches, e.g. when building allow lists
	if limit < 0 {
		cached, ok := fs.rowCache.Load(id, out.checksum)
		if ok {
			return *cached, nil
		}
	}

	if err := binary.Read(r, binary.LittleEndian, &out.count); err != nil {
		return out, errors.Wrap(err, "read doc count")
	}

	read := 0
	for {
		// limit >0 allows us to specify -1 to mean unlimited
		if limit > 0 && read >= limit {
			// we are done because the user specified limit is reached
			break
		}

		var docID uint64
		if err := binary.Read(r, binary.LittleEndian, &docID); err != nil {
			if err == io.EOF {
				// we are done, because all entries are read
				break
			}

			return out, errors.Wrap(err, "read doc id")
		}

		var frequency float64

		if hasFrequency {
			if err := binary.Read(r, binary.LittleEndian, &frequency); err != nil {
				// EOF would be unexpected here, so any error including EOF is an error
				return out, errors.Wrap(err, "read doc frequency")
			}
		}

		if fs.deletedDocIDs.Contains(docID) {
			// make sure a deleted docID does not count into the limit, otherwise we
			// will return 0 results with a limit of n if the first n doc ids are
			// marked as deleted (gh-1308)
			continue
		}

		out.docIDs = append(out.docIDs, docPointer{id: docID, frequency: &frequency})
		read++
	}

	// only write into cache on unlimited requests of a certain length
	if limit < 0 && read > 500 { // TODO: what's a realistic cutoff?
		fs.rowCache.Store(id, &out)
	}

	return out, nil
}

func (fs *Searcher) extractPropValuePair(filter *filters.Clause,
	className schema.ClassName) (*propValuePair, error) {
	var out propValuePair
	if filter.Operands != nil {
		// nested filter
		out.children = make([]*propValuePair, len(filter.Operands))

		for i, clause := range filter.Operands {
			child, err := fs.extractPropValuePair(&clause, className)
			if err != nil {
				return nil, errors.Wrapf(err, "nested clause at pos %d", i)
			}
			out.children[i] = child
		}
		out.operator = filter.Operator
		return &out, nil
	}

	// on value or non-nested filter
	props := filter.On.Slice()
	if len(props) != 1 {
		return fs.extractReferenceFilter(filter, className)
	}
	// we are on a value element

	if fs.onRefProp(className, props[0]) && filter.Value.Type == schema.DataTypeInt {
		// ref prop and int type is a special case, the user is looking for the
		// reference count as opposed to the content
		return fs.extractReferenceCount(props[0], filter.Value.Value, filter.Operator)
	}

	if fs.onGeoProp(className, props[0]) {
		return fs.extractGeoFilter(props[0], filter.Value.Value, filter.Value.Type,
			filter.Operator)
	}

	if fs.onUUIDProp(props[0]) {
		return fs.extractUUIDProp(filter.Value.Value, filter.Operator)
	}

	if fs.onMultiWordPropValue(filter.Operator, filter.Value.Value, filter.Value.Type) {
		return fs.extractMultiWordProp(props[0], filter.Value.Type, filter.Value.Value,
			filter.Operator)
	}

	return fs.extractPrimitiveProp(props[0], filter.Value.Type, filter.Value.Value,
		filter.Operator)
}

func (fs *Searcher) extractReferenceFilter(filter *filters.Clause,
	className schema.ClassName) (*propValuePair, error) {
	ctx := context.TODO()
	return newRefFilterExtractor(fs.classSearcher, filter, className, fs.schema).Do(ctx)
}

func (fs *Searcher) extractPrimitiveProp(propName string, dt schema.DataType,
	value interface{}, operator filters.Operator) (*propValuePair, error) {
	var extractValueFn func(in interface{}) ([]byte, error)
	var hasFrequency bool
	switch dt {
	case schema.DataTypeText:
		if operator == filters.OperatorLike {
			// if the operator is like, we cannot apply the regular text-splitting
			// logic as it would remove all wildcard symbols
			extractValueFn = fs.extractTextValueKeepWildcards
		} else {
			extractValueFn = fs.extractTextValue
		}
		hasFrequency = true
	case schema.DataTypeString:
		extractValueFn = fs.extractStringValue
		hasFrequency = true
	case schema.DataTypeBoolean:
		extractValueFn = fs.extractBoolValue
		hasFrequency = false
	case schema.DataTypeInt:
		extractValueFn = fs.extractIntValue
		hasFrequency = false
	case schema.DataTypeNumber:
		extractValueFn = fs.extractNumberValue
		hasFrequency = false
	case schema.DataTypeDate:
		extractValueFn = fs.extractDateValue
		hasFrequency = false
	case "":
		return nil, fmt.Errorf("data type cannot be empty")
	default:
		return nil, fmt.Errorf("data type %q not supported yet in standalone mode, "+
			"see %s for details", dt, notimplemented.Link)
	}

	byteValue, err := extractValueFn(value)
	if err != nil {
		return nil, err
	}

	return &propValuePair{
		value:        byteValue,
		hasFrequency: hasFrequency,
		prop:         propName,
		operator:     operator,
	}, nil
}

func (fs *Searcher) extractReferenceCount(propName string, value interface{},
	operator filters.Operator) (*propValuePair, error) {
	byteValue, err := fs.extractIntCountValue(value)
	if err != nil {
		return nil, err
	}

	return &propValuePair{
		value:        byteValue,
		hasFrequency: false,
		prop:         helpers.MetaCountProp(propName),
		operator:     operator,
	}, nil
}

func (fs *Searcher) extractGeoFilter(propName string, value interface{},
	valueType schema.DataType, operator filters.Operator) (*propValuePair, error) {
	if valueType != schema.DataTypeGeoCoordinates {
		return nil, fmt.Errorf("prop %q is of type geoCoordinates, it can only"+
			"be used with geoRange filters", propName)
	}

	parsed := value.(filters.GeoRange)

	return &propValuePair{
		value:         nil, // not going to be served by an inverted index
		valueGeoRange: &parsed,
		hasFrequency:  false,
		prop:          propName,
		operator:      operator,
	}, nil
}

func (fs *Searcher) extractUUIDProp(value interface{},
	operator filters.Operator) (*propValuePair, error) {
	v, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("expected value to be string, got %T", value)
	}

	return &propValuePair{
		value:        []byte(v),
		hasFrequency: false,
		prop:         helpers.PropertyNameUUID,
		operator:     operator,
	}, nil
}

func (fs *Searcher) extractMultiWordProp(propName string, dt schema.DataType,
	value interface{}, operator filters.Operator) (*propValuePair, error) {
	var out propValuePair
	var parts []string
	switch dt {
	case schema.DataTypeString:
		parts = helpers.TokenizeString(value.(string))
	case schema.DataTypeText:
		parts = helpers.TokenizeText(value.(string))
	default:
		return nil, fmt.Errorf("expected value type to be string or text, got %T", dt)
	}

	out.children = make([]*propValuePair, len(parts))

	for i, part := range parts {
		child, err := fs.extractPrimitiveProp(propName, dt, part, operator)
		if err != nil {
			return nil, errors.Wrapf(err, "multi word at pos %d", i)
		}
		out.children[i] = child
	}
	out.operator = filters.OperatorAnd

	return &out, nil
}

// TODO: repeated calls to on... aren't too efficient because we iterate over
// the schema each time, might be smarter to have a single method that
// determines the type and then we switch based on the result. However, the
// effect of that should be very small unless the schema is absolutely massive.
func (fs *Searcher) onRefProp(className schema.ClassName, propName string) bool {
	c := fs.schema.FindClassByName(className)
	if c == nil {
		return false
	}

	for _, prop := range c.Properties {
		if prop.Name != propName {
			continue
		}

		if schema.IsRefDataType(prop.DataType) {
			return true
		}
	}

	return false
}

// TODO: repeated calls to on... aren't too efficient because we iterate over
// the schema each time, might be smarter to have a single method that
// determines the type and then we switch based on the result. However, the
// effect of that should be very small unless the schema is absolutely massive.
func (fs *Searcher) onGeoProp(className schema.ClassName, propName string) bool {
	c := fs.schema.FindClassByName(className)
	if c == nil {
		return false
	}

	for _, prop := range c.Properties {
		if prop.Name != propName {
			continue
		}

		return schema.DataType(prop.DataType[0]) == schema.DataTypeGeoCoordinates
	}

	return false
}

func (fs *Searcher) onUUIDProp(propName string) bool {
	return propName == helpers.PropertyNameUUID
}

func (fs *Searcher) onMultiWordPropValue(operator filters.Operator,
	value interface{}, valueType schema.DataType) bool {
	switch valueType {
	case schema.DataTypeString:
		parts := helpers.TokenizeString(value.(string))
		return len(parts) > 1
	case schema.DataTypeText:
		var parts []string
		if operator == filters.OperatorLike {
			// if the operator is like, we cannot apply the regular text-splitting
			// logic as it would remove all wildcard symbols
			parts = helpers.TokenizeTextKeepWildcards(value.(string))
		} else {
			parts = helpers.TokenizeText(value.(string))
		}
		return len(parts) > 1
	default:
		return false
	}
}

type docPointers struct {
	count    uint64
	docIDs   []docPointer
	checksum []byte // helps us judge if a cached read is still fresh
}

type docPointer struct {
	id        uint64
	frequency *float64
}

func (d docPointers) IDs() []uint64 {
	out := make([]uint64, len(d.docIDs))
	for i, elem := range d.docIDs {
		out[i] = elem.id
	}
	return out
}
