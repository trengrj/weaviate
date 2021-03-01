//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package objects

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/modulecapabilities"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/search"
	modcontextionaryadditional "github.com/semi-technologies/weaviate/modules/text2vec-contextionary/additional"
	modcontextionaryadditionalprojector "github.com/semi-technologies/weaviate/modules/text2vec-contextionary/additional/projector"
	modcontextionaryadditionalsempath "github.com/semi-technologies/weaviate/modules/text2vec-contextionary/additional/sempath"
	"github.com/semi-technologies/weaviate/usecases/traverser"
	"github.com/stretchr/testify/mock"
)

type fakeSchemaManager struct {
	CalledWith struct {
		fromClass string
		property  string
		toClass   string
	}
	GetSchemaResponse schema.Schema
}

func (f *fakeSchemaManager) UpdatePropertyAddDataType(ctx context.Context, principal *models.Principal,
	fromClass, property, toClass string) error {
	f.CalledWith = struct {
		fromClass string
		property  string
		toClass   string
	}{
		fromClass: fromClass,
		property:  property,
		toClass:   toClass,
	}
	return nil
}

func (f *fakeSchemaManager) GetSchema(principal *models.Principal) (schema.Schema, error) {
	return f.GetSchemaResponse, nil
}

type fakeLocks struct{}

func (f *fakeLocks) LockConnector() (func() error, error) {
	return func() error { return nil }, nil
}

func (f *fakeLocks) LockSchema() (func() error, error) {
	return func() error { return nil }, nil
}

type fakeVectorizerProvider struct {
	vectorizer *fakeVectorizer
}

func (f *fakeVectorizerProvider) Vectorizer(modName, className string) (Vectorizer, error) {
	return f.vectorizer, nil
}

type fakeVectorizer struct {
	mock.Mock
}

func (f *fakeVectorizer) UpdateObject(ctx context.Context, object *models.Object) error {
	args := f.Called(object)
	object.Vector = args.Get(0).([]float32)
	return args.Error(1)
}

func (f *fakeVectorizer) Corpi(ctx context.Context, corpi []string) ([]float32, error) {
	panic("not implemented")
}

type fakeAuthorizer struct{}

func (f *fakeAuthorizer) Authorize(principal *models.Principal, verb, resource string) error {
	return nil
}

type fakeVectorRepo struct {
	mock.Mock
}

func (f *fakeVectorRepo) Exists(ctx context.Context,
	id strfmt.UUID) (bool, error) {
	args := f.Called(id)
	return args.Bool(0), args.Error(1)
}

func (f *fakeVectorRepo) ObjectByID(ctx context.Context,
	id strfmt.UUID, props traverser.SelectProperties, additional traverser.AdditionalProperties) (*search.Result, error) {
	args := f.Called(id, props, additional)
	return args.Get(0).(*search.Result), args.Error(1)
}

func (f *fakeVectorRepo) ObjectSearch(ctx context.Context, limit int,
	filters *filters.LocalFilter, additional traverser.AdditionalProperties) (search.Results, error) {
	args := f.Called(limit, filters, additional)
	return args.Get(0).([]search.Result), args.Error(1)
}

func (f *fakeVectorRepo) PutObject(ctx context.Context,
	concept *models.Object, vector []float32) error {
	args := f.Called(concept, vector)
	return args.Error(0)
}

func (f *fakeVectorRepo) BatchPutObjects(ctx context.Context, batch BatchObjects) (BatchObjects, error) {
	args := f.Called(batch)
	return batch, args.Error(0)
}

func (f *fakeVectorRepo) AddBatchReferences(ctx context.Context, batch BatchReferences) (BatchReferences, error) {
	args := f.Called(batch)
	return batch, args.Error(0)
}

func (f *fakeVectorRepo) Merge(ctx context.Context, merge MergeDocument) error {
	args := f.Called(merge)
	return args.Error(0)
}

func (f *fakeVectorRepo) DeleteObject(ctx context.Context,
	className string, id strfmt.UUID) error {
	args := f.Called(className, id)
	return args.Error(0)
}

func (f *fakeVectorRepo) AddReference(ctx context.Context,
	class string, source strfmt.UUID, prop string,
	ref *models.SingleRef) error {
	args := f.Called(source, prop, ref)
	return args.Error(0)
}

type fakeExtender struct {
	multi []search.Result
}

func (f *fakeExtender) AdditionalPropertyFn(ctx context.Context,
	in []search.Result, params interface{}, limit *int) ([]search.Result, error) {
	return f.multi, nil
}

func (f *fakeExtender) ExtractAdditionalFn(param []*ast.Argument) interface{} {
	return nil
}

func (f *fakeExtender) DefaultValueFn() interface{} {
	return getDefaultParam("nearestNeighbors")
}

type fakeProjector struct {
	multi []search.Result
}

func (f *fakeProjector) AdditionalPropertyFn(ctx context.Context,
	in []search.Result, params interface{}, limit *int) ([]search.Result, error) {
	return f.multi, nil
}

func (f *fakeProjector) ExtractAdditionalFn(param []*ast.Argument) interface{} {
	return nil
}

func (f *fakeProjector) DefaultValueFn() interface{} {
	return getDefaultParam("featureProjection")
}

type fakePathBuilder struct {
	multi []search.Result
}

func (f *fakePathBuilder) AdditionalPropertyFn(ctx context.Context,
	in []search.Result, params interface{}, limit *int) ([]search.Result, error) {
	return f.multi, nil
}

func (f *fakePathBuilder) ExtractAdditionalFn(param []*ast.Argument) interface{} {
	return nil
}

func (f *fakePathBuilder) DefaultValueFn() interface{} {
	return getDefaultParam("semanticPath")
}

type fakeModulesProvider struct {
	customExtender  *fakeExtender
	customProjector *fakeProjector
}

func (p *fakeModulesProvider) AdditionalPropertyFunction(name string) modulecapabilities.AdditionalPropertyFn {
	fns := modcontextionaryadditional.New(p.getExtender(), p.getProjector(), &fakePathBuilder{}).AdditionalPropertiesFunctions()
	if fns != nil {
		return fns[name]
	}
	return nil
}

func (p *fakeModulesProvider) GetObjectAdditionalExtend(ctx context.Context,
	in *search.Result, moduleParams map[string]interface{}) (*search.Result, error) {
	res, err := p.additionalExtend(ctx, search.Results{*in}, moduleParams, "ObjectGet")
	if err != nil {
		return nil, err
	}
	return &res[0], nil
}

func (p *fakeModulesProvider) ListObjectsAdditionalExtend(ctx context.Context,
	in search.Results, moduleParams map[string]interface{}) (search.Results, error) {
	return p.additionalExtend(ctx, in, moduleParams, "ObjectList")
}

func (p *fakeModulesProvider) additionalExtend(ctx context.Context,
	in search.Results, moduleParams map[string]interface{}, capability string) (search.Results, error) {
	txt2vec := modcontextionaryadditional.New(p.getExtender(), p.getProjector(), &fakePathBuilder{})
	fns := txt2vec.SearchAdditionalFunctions()
	if err := p.checkCapabilities(fns, moduleParams, capability); err != nil {
		return nil, err
	}
	for name, value := range moduleParams {
		additionalPropertyFn := p.getAdditionalPropertyFn(fns[name], capability)
		if additionalPropertyFn != nil && value != nil {
			resArray, err := additionalPropertyFn(ctx, in, nil, nil)
			if err != nil {
				return nil, err
			}
			in = resArray
		}
	}
	return in, nil
}

func (p *fakeModulesProvider) checkCapabilities(fns map[string]modulecapabilities.AdditionalSearch,
	moduleParams map[string]interface{}, capability string) error {
	for name := range moduleParams {
		additionalPropertyFn := p.getAdditionalPropertyFn(fns[name], capability)
		if additionalPropertyFn == nil {
			return errors.Errorf("unknown capability: %s", name)
		}
	}
	return nil
}

func (p *fakeModulesProvider) getAdditionalPropertyFn(searchAdditionalFns modulecapabilities.AdditionalSearch,
	capability string) modulecapabilities.AdditionalPropertyFn {
	switch capability {
	case "ObjectGet":
		return searchAdditionalFns.ObjectGet
	case "ObjectList":
		return searchAdditionalFns.ObjectList
	case "ExploreGet":
		return searchAdditionalFns.ExploreGet
	case "ExploreList":
		return searchAdditionalFns.ExploreList
	default:
		return nil
	}
}

func (p *fakeModulesProvider) getExtender() *fakeExtender {
	if p.customExtender != nil {
		return p.customExtender
	}
	return &fakeExtender{}
}

func (p *fakeModulesProvider) getProjector() *fakeProjector {
	if p.customProjector != nil {
		return p.customProjector
	}
	return &fakeProjector{}
}

func getDefaultParam(name string) interface{} {
	switch name {
	case "featureProjection":
		return &modcontextionaryadditionalprojector.Params{}
	case "semanticPath":
		return &modcontextionaryadditionalsempath.Params{}
	case "nearestNeighbors":
		return true
	default:
		return nil
	}
}

func getFakeModulesProvider() *fakeModulesProvider {
	return &fakeModulesProvider{}
}

func getFakeModulesProviderWithCustomExtenders(
	customExtender *fakeExtender,
	customProjector *fakeProjector,
) *fakeModulesProvider {
	return &fakeModulesProvider{customExtender, customProjector}
}
