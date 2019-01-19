package meta

import (
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/local/getmeta"
)

// TypeInspector can process the types of each specified props
type TypeInspector struct {
	typeSource typeSource
}

// NewTypeInspector to process the types of MetaPoperties
func NewTypeInspector(typeSource typeSource) *TypeInspector {
	return &TypeInspector{typeSource: typeSource}
}

// Process returns a simple map where each property is the key, the value
// contains the analysis prop that the user asked for through the graphQL API
func (t *TypeInspector) Process(params *getmeta.Params) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	for _, prop := range params.Properties {
		if prop.Name == MetaProp {
			// no typing is possible on the generic "meta" prop, skip! If we didn't
			// skip we might incorrectly error later on trying to look up the type of
			// this prop (which doesn't exist on the schema as it's a helper
			// construct that can be applied to any class.
			continue
		}

		propResult, err := t.analyzeAll(params, prop)
		if err != nil {
			return nil, err
		}

		if propResult != nil {
			result[string(prop.Name)] = propResult
		}
	}

	return result, nil
}

func (t *TypeInspector) analyzeAll(params *getmeta.Params,
	prop getmeta.MetaProperty) (map[string]interface{}, error) {
	results := []map[string]interface{}{}
	for _, analysis := range prop.StatisticalAnalyses {
		result, err := t.analyze(params, prop, analysis)
		if err != nil {
			return nil, err
		}

		if result == nil {
			continue
		}

		results = append(results, result)
	}

	switch len(results) {
	case 0:
		return nil, nil
	case 1:
		return results[0], nil
	case 2:
		return mergeMaps(results[0], results[1]), nil
	default:
		// there is no viable scenario where this is the case, the most the type
		// inspector can ever return is two results, which is only possible on a
		// CRef prop if the user asked for both 'type' and 'pointingTo'. All Other
		// propertys only have the statistical property 'type' which the type
		// inspector cares about. So in most cases the length will be 0 or 1 and 2
		// for ref-props. More than 2 indicates somethign went wrong.
		return nil, fmt.Errorf("got more than two results per property: %#v", results)
	}
}

func (t *TypeInspector) analyze(params *getmeta.Params, prop getmeta.MetaProperty,
	analysis getmeta.StatisticalAnalysis) (map[string]interface{}, error) {
	err, schemaProp := t.typeSource.GetProperty(params.Kind, params.ClassName, untitle(prop.Name))
	if err != nil {
		return nil, fmt.Errorf(
			"type inspector cannot get property %s.%s: %s", params.ClassName, prop.Name, err)
	}

	propType, err := t.typeSource.FindPropertyDataType(schemaProp.AtDataType)
	if err != nil {
		return nil, fmt.Errorf(
			"type inspector cannot get data type of property %s.%s: %s", params.ClassName, prop.Name, err)
	}

	switch analysis {
	case getmeta.PointingTo:
		return t.analyzeRefProp(params, propType)
	case getmeta.Type:
		return t.analyzePrimitiveProp(params, propType)
	default:
		return nil, nil
	}
}
func (t *TypeInspector) analyzePrimitiveProp(params *getmeta.Params,
	propType schema.PropertyDataType) (map[string]interface{}, error) {
	var typeName string
	if propType.IsPrimitive() {
		typeName = string(propType.AsPrimitive())
	} else {
		typeName = string(schema.DataTypeCRef)
	}

	return map[string]interface{}{
		string(getmeta.Type): typeName,
	}, nil
}

func (t *TypeInspector) analyzeRefProp(params *getmeta.Params,
	propType schema.PropertyDataType) (map[string]interface{}, error) {

	if !propType.IsReference() {
		return nil, fmt.Errorf("asked for pointingTo, but don't have a ref prop: %#v", propType)
	}

	return map[string]interface{}{
		string(getmeta.PointingTo): classSliceToInterfaceSlice(propType.Classes()),
	}, nil
}

func classSliceToInterfaceSlice(classes []schema.ClassName) []interface{} {
	result := make([]interface{}, len(classes), len(classes))
	for i, className := range classes {
		result[i] = string(className)
	}

	return result
}
