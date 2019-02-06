/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */

// Package network_get provides the network get graphql endpoint for Weaviate
package network_get

import (
	"fmt"
	"strings"

	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/descriptions"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/local/common_filters"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/graphql-go/graphql"
)

// Build the dynamically generated Get Actions part of the schema
func ActionClassFieldsFromSchema(dbSchema *schema.Schema, getActionsAndThings *map[string]*graphql.Object, weaviate string) (*graphql.Object, error) {
	actionClassFields := graphql.Fields{}

	for _, class := range dbSchema.Actions.Classes {
		singleActionClassField, singleActionClassObject := actionClassField(class, getActionsAndThings, weaviate)
		actionClassFields[class.Class] = singleActionClassField
		// this line assigns the created class to a Hashmap which is used in thunks to handle cyclical relationships (Classes with other Classes as properties)
		(*getActionsAndThings)[class.Class] = singleActionClassObject
	}

	getActions := graphql.ObjectConfig{
		Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGet", weaviate, "ActionsObj"),
		Fields:      actionClassFields,
		Description: descriptions.NetworkGetWeaviateActionsObj,
	}

	return graphql.NewObject(getActions), nil
}

func actionClassField(class *models.SemanticSchemaClass, getActionsAndThings *map[string]*graphql.Object, weaviate string) (*graphql.Field, *graphql.Object) {
	actionClassPropertyFields := graphql.ObjectConfig{
		Name: fmt.Sprintf("%s%s", weaviate, class.Class),
		Fields: (graphql.FieldsThunk)(func() graphql.Fields {
			singleactionClassPropertyFields, err := actionClassPropertyFields(class, getActionsAndThings, weaviate)

			if err != nil {
				panic("Failed to generate single Network Action Class property fields")
			}

			return singleactionClassPropertyFields
		}),
		Description: class.Description,
	}

	actionClassPropertyFieldsObj := graphql.NewObject(actionClassPropertyFields)

	actionClassPropertyFieldsField := &graphql.Field{
		Type:        graphql.NewList(actionClassPropertyFieldsObj),
		Description: class.Description,
		Args: graphql.FieldConfigArgument{
			"first": &graphql.ArgumentConfig{
				Description: descriptions.First,
				Type:        graphql.Int,
			},
			"after": &graphql.ArgumentConfig{
				Description: descriptions.After,
				Type:        graphql.Int,
			},
			"where": &graphql.ArgumentConfig{
				Description: descriptions.NetworkGetWhere,
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name:        fmt.Sprintf("WeaviateNetworkGet%sActions%sWhereInpObj", weaviate, class.Class),
						Fields:      common_filters.BuildNew(fmt.Sprintf("WeaviateNetworkGet%sActions%s", weaviate, class.Class)),
						Description: descriptions.NetworkGetWhereInpObj,
					},
				),
			},
		},
		Resolve: ResolveAction,
	}
	return actionClassPropertyFieldsField, actionClassPropertyFieldsObj
}

func actionClassPropertyFields(class *models.SemanticSchemaClass, getActionsAndThings *map[string]*graphql.Object, weaviate string) (graphql.Fields, error) {
	actionClassPropertyFields := graphql.Fields{}

	for _, property := range class.Properties {
		propertyType, err := schema.GetPropertyDataType(class, property.Name)

		if err != nil {
			return nil, err
		}

		if *propertyType == schema.DataTypeCRef {
			capitalizedPropertyName := strings.Title(property.Name)
			numberOfDataTypes := len(property.AtDataType)
			dataTypeClasses := make([]*graphql.Object, numberOfDataTypes)

			for index, dataType := range property.AtDataType {
				thingOrActionType, ok := (*getActionsAndThings)[dataType]

				if !ok {
					return nil, fmt.Errorf("no such thing/action class '%s'", property.AtDataType[index])
				}

				dataTypeClasses[index] = thingOrActionType
			}

			dataTypeUnionConf := graphql.UnionConfig{
				Name:  fmt.Sprintf("%s%s%s%s", weaviate, class.Class, capitalizedPropertyName, "Obj"),
				Types: dataTypeClasses,
				ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
					return nil
				},
				Description: property.Description,
			}

			multipleClassDataTypesUnion := graphql.NewUnion(dataTypeUnionConf)

			actionClassPropertyFields[capitalizedPropertyName] = &graphql.Field{
				Type:        multipleClassDataTypesUnion,
				Description: property.Description,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, fmt.Errorf("not supported")
				},
			}
		} else {
			convertedDataType, err := handleNetworkGetNonObjectDataTypes(*propertyType, property)

			if err != nil {
				return nil, err
			}

			actionClassPropertyFields[property.Name] = convertedDataType
		}
	}

	actionClassPropertyFields["uuid"] = &graphql.Field{
		Description: descriptions.NetworkGetClassUUID,
		Type:        graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return nil, fmt.Errorf("not supported")
		},
	}

	return actionClassPropertyFields, nil
}

// Build the dynamically generated Get Things part of the schema
func ThingClassFieldsFromSchema(dbSchema *schema.Schema, actionsAndThings *map[string]*graphql.Object, weaviate string) (*graphql.Object, error) {
	thingClassFields := graphql.Fields{}

	for _, class := range dbSchema.Things.Classes {
		singleThingClassField, singleThingClassObject := thingClassField(class, actionsAndThings, weaviate)
		thingClassFields[class.Class] = singleThingClassField
		// this line assigns the created class to a Hashmap which is used in thunks to handle cyclical relationships (Classes with other Classes as properties)
		(*actionsAndThings)[class.Class] = singleThingClassObject
	}

	getThings := graphql.ObjectConfig{
		Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGet", weaviate, "ThingsObj"),
		Fields:      thingClassFields,
		Description: descriptions.NetworkGetWeaviateThingsObj,
	}

	return graphql.NewObject(getThings), nil
}

func thingClassField(class *models.SemanticSchemaClass, getActionsAndThings *map[string]*graphql.Object, weaviate string) (*graphql.Field, *graphql.Object) {
	singleThingClassPropertyFieldsObj := graphql.ObjectConfig{
		Name: fmt.Sprintf("%s%s", weaviate, class.Class),
		Fields: (graphql.FieldsThunk)(func() graphql.Fields {
			singleThingClassPropertyFields, err := thingClassPropertyFields(class, getActionsAndThings, weaviate)
			if err != nil {
				panic(fmt.Errorf("failed to assemble single Network Thing Class field for Class %s", class.Class))
			}
			return singleThingClassPropertyFields
		}),
		Description: class.Description,
	}

	thingClassPropertyFieldsObject := graphql.NewObject(singleThingClassPropertyFieldsObj)
	thingClassPropertyFieldsField := &graphql.Field{
		Type:        graphql.NewList(thingClassPropertyFieldsObject),
		Description: class.Description,
		Args: graphql.FieldConfigArgument{
			"first": &graphql.ArgumentConfig{
				Description: descriptions.First,
				Type:        graphql.Int,
			},
			"after": &graphql.ArgumentConfig{
				Description: descriptions.After,
				Type:        graphql.Int,
			},
			"where": &graphql.ArgumentConfig{
				Description: descriptions.NetworkGetWhere,
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name:        fmt.Sprintf("WeaviateNetworkGet%sThings%sWhereInpObj", weaviate, class.Class),
						Fields:      common_filters.BuildNew(fmt.Sprintf("WeaviateNetworkGet%sThings%s", weaviate, class.Class)),
						Description: descriptions.NetworkGetWhereInpObj,
					},
				),
			},
		},
		Resolve: ResolveThing,
	}
	return thingClassPropertyFieldsField, thingClassPropertyFieldsObject
}

func thingClassPropertyFields(class *models.SemanticSchemaClass, actionsAndThings *map[string]*graphql.Object, weaviate string) (graphql.Fields, error) {
	fields := graphql.Fields{}

	for _, property := range class.Properties {

		propertyType, err := schema.GetPropertyDataType(class, property.Name)

		if err != nil {
			return nil, err
		}

		if *propertyType == schema.DataTypeCRef {
			capitalizedPropertyName := strings.Title(property.Name)
			numberOfDataTypes := len(property.AtDataType)
			dataTypeClasses := make([]*graphql.Object, numberOfDataTypes)

			for index, dataType := range property.AtDataType {
				thingOrActionType, ok := (*actionsAndThings)[dataType]

				if !ok {
					return nil, fmt.Errorf("no such thing/action class '%s'", property.AtDataType[index])
				}

				dataTypeClasses[index] = thingOrActionType
			}

			dataTypeUnionConf := graphql.UnionConfig{
				Name:  fmt.Sprintf("%s%s%s%s", weaviate, class.Class, capitalizedPropertyName, "Obj"),
				Types: dataTypeClasses,
				ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
					return nil
				},
				Description: property.Description,
			}

			multipleClassDataTypesUnion := graphql.NewUnion(dataTypeUnionConf)

			fields[capitalizedPropertyName] = &graphql.Field{
				Type:        multipleClassDataTypesUnion,
				Description: property.Description,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, fmt.Errorf("resolving single network thing class property field not supported")
				},
			}
		} else {
			convertedDataType, err := handleNetworkGetNonObjectDataTypes(*propertyType, property)

			if err != nil {
				return nil, err
			}

			fields[property.Name] = convertedDataType
		}
	}

	fields["uuid"] = &graphql.Field{
		Description: descriptions.NetworkGetClassUUID,
		Type:        graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return nil, fmt.Errorf("not supported")
		},
	}

	return fields, nil
}

func handleNetworkGetNonObjectDataTypes(dataType schema.DataType, property *models.SemanticSchemaClassProperty) (*graphql.Field, error) {

	switch dataType {

	case schema.DataTypeString:
		return &graphql.Field{
			Description: property.Description,
			Type:        graphql.String,
		}, nil

	case schema.DataTypeText:
		return &graphql.Field{
			Description: property.Description,
			Type:        graphql.String,
		}, nil

	case schema.DataTypeInt:
		return &graphql.Field{
			Description: property.Description,
			Type:        graphql.Int,
		}, nil

	case schema.DataTypeNumber:
		return &graphql.Field{
			Description: property.Description,
			Type:        graphql.Float,
		}, nil

	case schema.DataTypeBoolean:
		return &graphql.Field{
			Description: property.Description,
			Type:        graphql.Boolean,
		}, nil

	case schema.DataTypeDate:
		return &graphql.Field{
			Description: property.Description,
			Type:        graphql.String,
		}, nil

	default:
		return nil, fmt.Errorf("%s", schema.ErrorNoSuchDatatype)
	}
}
