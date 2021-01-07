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

package storobj

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/entities/search"
	"github.com/semi-technologies/weaviate/usecases/traverser"
)

type Object struct {
	MarshallerVersion uint8
	Kind              kind.Kind     `json:"kind"`
	Object            models.Object `json:"object"`
	Vector            []float32     `json:"vector"`
	docID             uint64
}

func New(k kind.Kind, docID uint64) *Object {
	return &Object{
		MarshallerVersion: 1,
		Kind:              k,
		docID:             docID,
	}
}

func FromObject(object *models.Object, vector []float32) *Object {
	return &Object{
		Kind:              kind.Object,
		Object:            *object,
		Vector:            vector,
		MarshallerVersion: 1,
	}
}

func FromBinary(data []byte) (*Object, error) {
	ko := &Object{}
	if err := ko.UnmarshalBinary(data); err != nil {
		return nil, err
	}

	return ko, nil
}

func (ko *Object) Class() schema.ClassName {
	switch ko.Kind {
	case kind.Object:
		return schema.ClassName(ko.Object.Class)
	default:
		panic(fmt.Sprintf("impossible kind: %s", ko.Kind.Name()))
	}
}

func (ko *Object) SetDocID(id uint64) {
	ko.docID = id
}

func (ko *Object) DocID() uint64 {
	return ko.docID
}

func (ko *Object) CreationTimeUnix() int64 {
	switch ko.Kind {
	case kind.Object:
		return ko.Object.CreationTimeUnix
	default:
		panic("impossible kind")
	}
}

func (ko *Object) ID() strfmt.UUID {
	switch ko.Kind {
	case kind.Object:
		return ko.Object.ID
	default:
		panic(fmt.Sprintf("impossible kind: %q", ko.Kind))
	}
}

func (ko *Object) SetID(id strfmt.UUID) {
	switch ko.Kind {
	case kind.Object:
		ko.Object.ID = id
	default:
		panic(fmt.Sprintf("impossible kind: %q", ko.Kind))
	}
}

func (ko *Object) SetClass(class string) {
	switch ko.Kind {
	case kind.Object:
		ko.Object.Class = class
	default:
		panic(fmt.Sprintf("impossible kind: %q", ko.Kind))
	}
}

func (ko *Object) LastUpdateTimeUnix() int64 {
	switch ko.Kind {
	case kind.Object:
		return ko.Object.LastUpdateTimeUnix
	default:
		panic("impossible kind")
	}
}

// AdditionalProperties groups all properties which are stored with the
// object and not generated at runtime
func (ko *Object) AdditionalProperties() *models.AdditionalProperties {
	switch ko.Kind {
	case kind.Object:
		return ko.Object.Additional
	default:
		panic("impossible kind")
	}
}

func (ko *Object) Schema() models.PropertySchema {
	switch ko.Kind {
	case kind.Object:
		return ko.Object.Properties
	default:
		panic("impossible kind")
	}
}

func (ko *Object) SchemaWithAdditional(
	additional traverser.AdditionalProperties) models.PropertySchema {
	schema := ko.Schema()

	if additional.RefMeta {
		// nothing to remove
		return schema
	}

	asMap, ok := schema.(map[string]interface{})
	if !ok || asMap == nil {
		return schema
	}

	for propName, value := range asMap {
		asRefs, ok := value.(models.MultipleRef)
		if !ok {
			// not a ref, we can skip
			continue
		}

		for i := range asRefs {
			asRefs[i].Classification = nil
		}

		asMap[propName] = asRefs
	}

	return asMap
}

func (ko *Object) SetSchema(schema models.PropertySchema) {
	switch ko.Kind {
	case kind.Object:
		ko.Object.Properties = schema
	default:
		panic("impossible kind")
	}
}

func (ko *Object) VectorWeights() models.VectorWeights {
	switch ko.Kind {
	case kind.Object:
		return ko.Object.VectorWeights
	default:
		panic("impossible kind")
	}
}

func (ko *Object) SearchResult(additional traverser.AdditionalProperties) *search.Result {
	schemaMap, ok := ko.SchemaWithAdditional(additional).(map[string]interface{})
	if !ok || schemaMap == nil {
		schemaMap = map[string]interface{}{}
	}
	schemaMap["id"] = ko.ID()
	ko.SetSchema(schemaMap)

	additionalProperties := &models.AdditionalProperties{}
	if ko.AdditionalProperties() != nil {
		if additional.Interpretation {
			additionalProperties.Interpretation = ko.AdditionalProperties().Interpretation
		}
		if additional.Classification {
			additionalProperties.Classification = ko.AdditionalProperties().Classification
		}
	}

	return &search.Result{
		Kind:      ko.Kind,
		ID:        ko.ID(),
		ClassName: ko.Class().String(),
		Schema:    ko.Schema(),
		Vector:    ko.Vector,
		// VectorWeights: ko.VectorWeights(), // TODO: add vector weights
		Created:              ko.CreationTimeUnix(),
		Updated:              ko.LastUpdateTimeUnix(),
		AdditionalProperties: additionalProperties,
		Score:                1, // TODO: actuallly score
		// TODO: Beacon?
	}
}

func (ko *Object) Valid() bool {
	return ko.Kind != "" &&
		ko.ID() != "" &&
		ko.Class().String() != ""
}

func SearchResults(in []*Object, additional traverser.AdditionalProperties) search.Results {
	out := make(search.Results, len(in))

	for i, elem := range in {
		out[i] = *(elem.SearchResult(additional))
	}

	return out
}

func DocIDFromBinary(in []byte) (uint64, error) {
	var version uint8
	r := bytes.NewReader(in)
	le := binary.LittleEndian
	if err := binary.Read(r, le, &version); err != nil {
		return 0, err
	}

	if version != 1 {
		return 0, fmt.Errorf("unsupported binary marshaller version %d", version)
	}

	var docID uint64
	err := binary.Read(r, le, &docID)
	return docID, err
}

// MarshalBinary creates the binary representation of a kind object. Regardless
// of the marshaller version the first byte is a uint8 indicating the version
// followed by the payload which depends on the specific version
//
// Version 1
// No. of B   | Type      | Content
// ------------------------------------------------
// 1          | uint8     | MarshallerVersion = 1
// 8          | uint64    | index id, keep early so id-only lookups are maximum efficient
// 1          | uint8     | kind, 0=action, 1=thing
// 16         | uint128   | uuid
// 8          | int64     | create time
// 8          | int64     | update time
// 2          | uint16    | VectorLength
// n*4        | []float32 | vector of length n
// 2          | uint16    | length of class name
// n          | []byte    | className
// 4          | uint32    | length of schema json
// n          | []byte    | schema as json
// 2          | uint32    | length of meta json
// n          | []byte    | meta as json
// 2          | uint32    | length of vectorweights json
// n          | []byte    | vectorweights as json
func (ko *Object) MarshalBinary() ([]byte, error) {
	if ko.MarshallerVersion != 1 {
		return nil, fmt.Errorf("unsupported marshaller version %d", ko.MarshallerVersion)
	}

	kindByte := uint8(0)
	if ko.Kind == kind.Object {
		kindByte = 1
	}

	idParsed, err := uuid.Parse(ko.ID().String())
	if err != nil {
		return nil, err
	}
	idBytes, err := idParsed.MarshalBinary()
	if err != nil {
		return nil, err
	}
	vectorLength := uint16(len(ko.Vector))
	className := []byte(ko.Class())
	classNameLength := uint16(len(className))
	schema, err := json.Marshal(ko.Schema())
	if err != nil {
		return nil, err
	}
	schemaLength := uint32(len(schema))
	meta, err := json.Marshal(ko.AdditionalProperties())
	if err != nil {
		return nil, err
	}
	metaLength := uint32(len(meta))
	vectorWeights, err := json.Marshal(ko.VectorWeights())
	if err != nil {
		return nil, err
	}
	vectorWeightsLength := uint32(len(vectorWeights))

	ec := &errorCompounder{}
	buf := bytes.NewBuffer(nil)
	le := binary.LittleEndian
	ec.add(binary.Write(buf, le, &ko.MarshallerVersion))
	ec.add(binary.Write(buf, le, &ko.docID))
	ec.add(binary.Write(buf, le, kindByte))
	_, err = buf.Write(idBytes)
	ec.add(err)
	ec.add(binary.Write(buf, le, ko.CreationTimeUnix()))
	ec.add(binary.Write(buf, le, ko.LastUpdateTimeUnix()))
	ec.add(binary.Write(buf, le, vectorLength))
	ec.add(binary.Write(buf, le, ko.Vector))
	ec.add(binary.Write(buf, le, classNameLength))
	_, err = buf.Write(className)
	ec.add(err)
	ec.add(binary.Write(buf, le, schemaLength))
	_, err = buf.Write(schema)
	ec.add(err)
	ec.add(binary.Write(buf, le, metaLength))
	_, err = buf.Write(meta)
	ec.add(err)
	ec.add(binary.Write(buf, le, vectorWeightsLength))
	_, err = buf.Write(vectorWeights)
	ec.add(err)

	return buf.Bytes(), ec.toError()
}

// UnmarshalBinary is the versioned way to unmarshal a kind object from binary,
// see MarshalBinary for the exact contents of each version
func (ko *Object) UnmarshalBinary(data []byte) error {
	var version uint8
	r := bytes.NewReader(data)
	le := binary.LittleEndian
	if err := binary.Read(r, le, &version); err != nil {
		return err
	}

	if version != 1 {
		return fmt.Errorf("unsupported binary marshaller version %d", version)
	}

	ko.MarshallerVersion = version

	var (
		kindByte            uint8
		uuidBytes           = make([]byte, 16)
		createTime          int64
		updateTime          int64
		vectorLength        uint16
		classNameLength     uint16
		schemaLength        uint32
		metaLength          uint32
		vectorWeightsLength uint32
	)

	ec := &errorCompounder{}
	ec.add(binary.Read(r, le, &ko.docID))
	ec.add(binary.Read(r, le, &kindByte))
	_, err := r.Read(uuidBytes)
	ec.add(err)
	ec.add(binary.Read(r, le, &createTime))
	ec.add(binary.Read(r, le, &updateTime))
	ec.add(binary.Read(r, le, &vectorLength))
	ko.Vector = make([]float32, vectorLength)
	ec.add(binary.Read(r, le, &ko.Vector))
	ec.add(binary.Read(r, le, &classNameLength))
	className := make([]byte, classNameLength)
	_, err = r.Read(className)
	ec.add(err)
	ec.add(binary.Read(r, le, &schemaLength))
	schema := make([]byte, schemaLength)
	_, err = r.Read(schema)
	ec.add(err)
	ec.add(binary.Read(r, le, &metaLength))
	meta := make([]byte, metaLength)
	_, err = r.Read(meta)
	ec.add(err)
	ec.add(binary.Read(r, le, &vectorWeightsLength))
	vectorWeights := make([]byte, vectorWeightsLength)
	_, err = r.Read(vectorWeights)
	ec.add(err)

	if ec.toError() != nil {
		return err
	}

	uuidParsed, err := uuid.FromBytes(uuidBytes)
	if err != nil {
		return err
	}

	if kindByte == 1 {
		ko.Kind = kind.Object
	} else {
		ko.Kind = kind.Object
	}

	return ko.parseKind(
		strfmt.UUID(uuidParsed.String()),
		createTime,
		updateTime,
		string(className),
		schema,
		meta,
		vectorWeights,
	)
}

func (ko *Object) parseKind(uuid strfmt.UUID, create, update int64, className string,
	schemaB []byte, additionalB []byte, vectorWeightsB []byte) error {
	var schema map[string]interface{}
	if err := json.Unmarshal(schemaB, &schema); err != nil {
		return err
	}

	if err := ko.enrichSchemaTypes(schema); err != nil {
		return errors.Wrap(err, "enrich schema datatypes")
	}

	var additional *models.AdditionalProperties
	if err := json.Unmarshal(additionalB, &additional); err != nil {
		return err
	}

	var vectorWeights interface{}
	if err := json.Unmarshal(vectorWeightsB, &vectorWeights); err != nil {
		return err
	}

	if ko.Kind == kind.Object {
		ko.Object = models.Object{
			Class:              className,
			CreationTimeUnix:   create,
			LastUpdateTimeUnix: update,
			ID:                 uuid,
			Properties:         schema,
			VectorWeights:      vectorWeights,
			Additional:         additional,
		}
	}

	return nil
}

type errorCompounder struct {
	errors []error
}

func (ec *errorCompounder) add(err error) {
	if err != nil {
		ec.errors = append(ec.errors, err)
	}
}

func (ec *errorCompounder) toError() error {
	if len(ec.errors) == 0 {
		return nil
	}

	var msg strings.Builder
	for i, err := range ec.errors {
		if i != 0 {
			msg.WriteString(", ")
		}

		msg.WriteString(err.Error())
	}

	return errors.New(msg.String())
}
