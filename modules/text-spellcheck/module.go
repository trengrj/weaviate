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

package modspellcheck

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/entities/modulecapabilities"
	"github.com/semi-technologies/weaviate/entities/moduletools"
	spellcheckadditional "github.com/semi-technologies/weaviate/modules/text-spellcheck/additional"
	spellcheckadditionalspellcheck "github.com/semi-technologies/weaviate/modules/text-spellcheck/additional/spellcheck"
	"github.com/semi-technologies/weaviate/modules/text-spellcheck/clients"
	"github.com/semi-technologies/weaviate/modules/text-spellcheck/ent"
	"github.com/sirupsen/logrus"
)

func New() *SpellCheckModule {
	return &SpellCheckModule{}
}

type SpellCheckModule struct {
	spellCheck                   spellCheckClient
	additionalPropertiesProvider modulecapabilities.AdditionalProperties
}

type spellCheckClient interface {
	Check(ctx context.Context, text []string) (*ent.SpellCheckResult, error)
	MetaInfo() (map[string]interface{}, error)
}

func (m *SpellCheckModule) Name() string {
	return "text-spellcheck"
}

func (m *SpellCheckModule) Init(ctx context.Context,
	params moduletools.ModuleInitParams) error {
	if err := m.initAdditional(ctx, params.GetLogger()); err != nil {
		return errors.Wrap(err, "init additional")
	}
	return nil
}

func (m *SpellCheckModule) initAdditional(ctx context.Context,
	logger logrus.FieldLogger) error {
	uri := os.Getenv("SPELLCHECK_INFERENCE_API")
	if uri == "" {
		return errors.Errorf("required variable SPELLCHECK_INFERENCE_API is not set")
	}

	client := clients.New(uri, logger)
	if err := client.WaitForStartup(ctx, 1*time.Second); err != nil {
		return errors.Wrap(err, "init remote spell check module")
	}

	m.spellCheck = client

	spellCheckProvider := spellcheckadditionalspellcheck.New(m.spellCheck)
	m.additionalPropertiesProvider = spellcheckadditional.New(spellCheckProvider)

	return nil
}

func (m *SpellCheckModule) RootHandler() http.Handler {
	// TODO: remove once this is a capability interface
	return nil
}

func (m *SpellCheckModule) MetaInfo() (map[string]interface{}, error) {
	return m.spellCheck.MetaInfo()
}

func (m *SpellCheckModule) AdditionalProperties() map[string]modulecapabilities.AdditionalProperty {
	return m.additionalPropertiesProvider.AdditionalProperties()
}

// verify we implement the modules.Module interface
var (
	_ = modulecapabilities.Module(New())
	_ = modulecapabilities.AdditionalProperties(New())
	_ = modulecapabilities.MetaProvider(New())
)
