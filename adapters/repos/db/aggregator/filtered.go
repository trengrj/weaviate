package aggregator

import (
	"context"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/adapters/repos/db/inverted"
	"github.com/semi-technologies/weaviate/adapters/repos/db/storobj"
	"github.com/semi-technologies/weaviate/entities/aggregation"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/usecases/traverser"
)

type filteredAggregator struct {
	*Aggregator
}

func newFilteredAggregator(agg *Aggregator) *filteredAggregator {
	return &filteredAggregator{Aggregator: agg}
}

func (fa *filteredAggregator) Do(ctx context.Context) (*aggregation.Result, error) {
	out := aggregation.Result{}

	// without grouping there is always exactly one group
	out.Groups = make([]aggregation.Group, 1)

	s := fa.getSchema.GetSchemaSkipAuth()
	ids, err := inverted.NewSearcher(fa.db, s, fa.invertedRowCache).
		DocIDs(ctx, fa.params.Filters, false, fa.params.ClassName)
	if err != nil {
		return nil, errors.Wrap(err, "retrieve doc IDs from searcher")
	}

	if fa.params.IncludeMetaCount {
		out.Groups[0].Count = len(ids)
	}

	props, err := fa.properties(ctx, ids)
	if err != nil {
		return nil, errors.Wrap(err, "aggregate properties")
	}

	out.Groups[0].Properties = props

	return &out, nil
}

func (fa *filteredAggregator) properties(ctx context.Context,
	ids inverted.AllowList) (map[string]aggregation.Property, error) {
	pointers := flattenAllowList(ids)

	propAggs, err := fa.prepareAggregatorsForProps()
	if err != nil {
		return nil, errors.Wrap(err, "prepare aggregators for props")
	}

	scan := func(obj *storobj.Object) (bool, error) {
		if err := fa.analyzeObject(ctx, obj, propAggs); err != nil {
			return false, errors.Wrapf(err, "analyze object %s", obj.ID())
		}
		return true, nil
	}

	if err := fa.db.View(func(tx *bolt.Tx) error {
		return inverted.ScanObjectsFromDocIDsInTx(tx, pointers, scan)
	}); err != nil {
		return nil, errors.Wrap(err, "properties view tx")
	}

	return propAggs.results()
}

func (fa *filteredAggregator) analyzeObject(ctx context.Context,
	obj *storobj.Object, propAggs map[string]propAgg) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if obj.Schema() == nil {
		return nil
	}

	for propName, prop := range propAggs {
		value, ok := obj.Schema().(map[string]interface{})[propName]
		if !ok {
			continue
		}

		fa.addPropValue(prop, value)
	}

	return nil
}

func (fa *filteredAggregator) addPropValue(prop propAgg, value interface{}) {
	switch prop.aggType {
	case aggregation.PropertyTypeBoolean:
		asBool, ok := value.(bool)
		if !ok {
			return
		}
		prop.boolAgg.AddBool(asBool)
	case aggregation.PropertyTypeNumerical:
		asFloat, ok := value.(float64)
		if !ok {
			return
		}
		prop.numericalAgg.AddFloat64(asFloat)
	default:
	}
}

// a helper type to select the right aggreagtor for a prop
type propAgg struct {
	name schema.PropertyName

	// the user is interested in those specific aggregations
	specifiedAggregators []traverser.Aggregator

	// underlying data type of prop
	dataType schema.DataType

	// use aggType to chose with agg to use
	aggType aggregation.PropertyType

	// only one of the following three would ever best
	boolAgg      *boolAggregator
	textAgg      *textAggregator
	numericalAgg *numericalAggregator
}

// propAggs groups propAgg helpers by prop name
type propAggs map[string]propAgg

func (pa *propAgg) initAggregator() {
	switch pa.aggType {
	case aggregation.PropertyTypeText:
		limit := extractLimitFromTopOccs(pa.specifiedAggregators)
		pa.textAgg = newTextAggregator(limit)
	case aggregation.PropertyTypeBoolean:
		pa.boolAgg = newBoolAggregator()
	case aggregation.PropertyTypeNumerical:
		pa.numericalAgg = newNumericalAggregator()
	default:
	}
}

func (pa propAggs) results() (map[string]aggregation.Property, error) {
	out := map[string]aggregation.Property{}

	for _, prop := range pa {
		aggProp := aggregation.Property{
			Type: prop.aggType,
		}

		switch prop.aggType {
		case aggregation.PropertyTypeBoolean:
			aggProp.BooleanAggregation = prop.boolAgg.Res()
			out[prop.name.String()] = aggProp

		case aggregation.PropertyTypeNumerical:
			prop.numericalAgg.buildPairsFromCounts()
			addNumericalAggregations(&aggProp, prop.specifiedAggregators,
				prop.numericalAgg)
			out[prop.name.String()] = aggProp

		case aggregation.PropertyTypeText:
		default:
		}
	}

	return out, nil
}

func (fa *filteredAggregator) prepareAggregatorsForProps() (propAggs, error) {
	out := propAggs{}

	for _, prop := range fa.params.Properties {
		pa := propAgg{
			name:                 prop.Name,
			specifiedAggregators: prop.Aggregators,
		}

		at, dt, err := fa.aggTypeOfProperty(prop.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "property %s", prop.Name)
		}

		pa.aggType = at
		pa.dataType = dt
		pa.initAggregator()
		out[prop.Name.String()] = pa
	}

	return out, nil
}

func flattenAllowList(list inverted.AllowList) []uint32 {
	out := make([]uint32, len(list))
	i := 0
	for id := range list {
		out[i] = id
		i++
	}

	return out
}
