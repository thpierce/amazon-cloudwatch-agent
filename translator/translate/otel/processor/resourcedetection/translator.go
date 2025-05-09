// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package resourcedetection

import (
	_ "embed"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/processor"

	"github.com/aws/amazon-cloudwatch-agent/translator/config"
	"github.com/aws/amazon-cloudwatch-agent/translator/context"
	"github.com/aws/amazon-cloudwatch-agent/translator/translate/otel/common"
	"github.com/aws/amazon-cloudwatch-agent/translator/translate/otel/extension/agenthealth"
	"github.com/aws/amazon-cloudwatch-agent/translator/util/ecsutil"
)

//go:embed configs/config.yaml
var appSignalsDefaultResourceDetectionConfig string

//go:embed configs/ecs_config.yaml
var appSignalsECSResourceDetectionConfig string

type translator struct {
	name    string
	signal  pipeline.Signal
	factory processor.Factory
}

type Option interface {
	apply(t *translator)
}

type optionFunc func(t *translator)

func (o optionFunc) apply(t *translator) {
	o(t)
}

// WithSignal determines where the translator should look to find
// the configuration.
func WithSignal(signal pipeline.Signal) Option {
	return optionFunc(func(t *translator) {
		t.signal = signal
	})
}

var _ common.ComponentTranslator = (*translator)(nil)

func NewTranslator(opts ...Option) common.ComponentTranslator {
	t := &translator{factory: resourcedetectionprocessor.NewFactory()}
	for _, opt := range opts {
		opt.apply(t)
	}
	return t
}

func (t *translator) ID() component.ID {
	return component.NewIDWithName(t.factory.Type(), t.name)
}

func (t *translator) Translate(conf *confmap.Conf) (component.Config, error) {
	cfg := t.factory.CreateDefaultConfig().(*resourcedetectionprocessor.Config)
	cfg.MiddlewareID = &agenthealth.StatusCodeID
	mode := context.CurrentContext().KubernetesMode()
	if mode == "" {
		mode = context.CurrentContext().Mode()
	}
	if mode == config.ModeEC2 {
		if ecsutil.GetECSUtilSingleton().IsECS() {
			mode = config.ModeECS
		}
	}

	switch mode {
	case config.ModeECS:
		return common.GetYamlFileToYamlConfig(cfg, appSignalsECSResourceDetectionConfig)
	default:
		return common.GetYamlFileToYamlConfig(cfg, appSignalsDefaultResourceDetectionConfig)
	}

}
