// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package otel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pipeline"

	"github.com/aws/amazon-cloudwatch-agent/tool/testutil"
	"github.com/aws/amazon-cloudwatch-agent/translator"
	_ "github.com/aws/amazon-cloudwatch-agent/translator/registerrules"
	"github.com/aws/amazon-cloudwatch-agent/translator/translate/agent"
	"github.com/aws/amazon-cloudwatch-agent/translator/translate/otel/common"
	"github.com/aws/amazon-cloudwatch-agent/translator/translate/otel/pipeline/prometheus"
	"github.com/aws/amazon-cloudwatch-agent/translator/util/eksdetector"
)

func TestTranslator(t *testing.T) {
	agent.Global_Config.Region = "us-east-1"
	testutil.SetPrometheusRemoteWriteTestingEnv(t)
	testCases := map[string]struct {
		input           interface{}
		wantErrContains string
		detector        func() (eksdetector.Detector, error)
		isEKSDataStore  func() eksdetector.IsEKSCache
	}{
		"WithValidConfig": {
			input: map[string]interface{}{
				"agent": map[string]interface{}{
					"debug": true,
				},
				"logs": map[string]interface{}{
					"metrics_collected": map[string]interface{}{
						"kubernetes": map[string]interface{}{
							"cluster_name":           "TestCluster",
							"jmx_container_insights": true,
						},
					},
				},
			},
		},
		"WithEmptyConfig": {
			input:           map[string]interface{}{},
			wantErrContains: "no valid pipelines",
		},
		"WithoutReceivers": {
			input: map[string]interface{}{
				"metrics": map[string]interface{}{},
			},
			wantErrContains: "no valid pipelines",
		},
		"WithMinimalConfig": {
			input: map[string]interface{}{
				"metrics": map[string]interface{}{
					"metrics_collected": map[string]interface{}{
						"cpu": map[string]interface{}{},
					},
				},
			},
		},
		"WithAppSignalsMetricsEnabled": {
			input: map[string]interface{}{
				"logs": map[string]interface{}{
					"metrics_collected": map[string]interface{}{
						"application_signals": map[string]interface{}{},
					},
				},
			},
			detector:       eksdetector.TestEKSDetector,
			isEKSDataStore: eksdetector.TestIsEKSCacheEKS,
		},
		"WithAppSignalsTracesEnabled": {
			input: map[string]interface{}{
				"traces": map[string]interface{}{
					"traces_collected": map[string]interface{}{
						"application_signals": map[string]interface{}{},
					},
				},
			},
			detector:       eksdetector.TestEKSDetector,
			isEKSDataStore: eksdetector.TestIsEKSCacheEKS,
		},
		"WithAppSignalsMetricsAndTracesEnabled": {
			input: map[string]interface{}{
				"logs": map[string]interface{}{
					"metrics_collected": map[string]interface{}{
						"application_signals": map[string]interface{}{},
					},
				},
				"traces": map[string]interface{}{
					"traces_collected": map[string]interface{}{
						"application_signals": map[string]interface{}{},
					},
				},
			},
			detector:       eksdetector.TestEKSDetector,
			isEKSDataStore: eksdetector.TestIsEKSCacheEKS,
		},
		"WithAppSignalsMultipleMetricsReceiversConfig": {
			input: map[string]interface{}{
				"logs": map[string]interface{}{
					"metrics_collected": map[string]interface{}{
						"application_signals": map[string]interface{}{},
						"cpu":                 map[string]interface{}{},
					},
				},
				"traces": map[string]interface{}{
					"traces_collected": map[string]interface{}{
						"application_signals": map[string]interface{}{},
						"otlp":                map[string]interface{}{},
						"otlp2":               map[string]interface{}{},
					},
				},
			},
			detector:       eksdetector.TestEKSDetector,
			isEKSDataStore: eksdetector.TestIsEKSCacheEKS,
		},
		"WithAppSignalsFallbackMetricsEnabled": {
			input: map[string]interface{}{
				"logs": map[string]interface{}{
					"metrics_collected": map[string]interface{}{
						"app_signals": map[string]interface{}{},
					},
				},
			},
			detector:       eksdetector.TestEKSDetector,
			isEKSDataStore: eksdetector.TestIsEKSCacheEKS,
		},
		"WithAppSignalsFallbackTracesEnabled": {
			input: map[string]interface{}{
				"traces": map[string]interface{}{
					"traces_collected": map[string]interface{}{
						"app_signals": map[string]interface{}{},
					},
				},
			},
			detector:       eksdetector.TestEKSDetector,
			isEKSDataStore: eksdetector.TestIsEKSCacheEKS,
		},
		"WithAppSignalsFallbackMetricsAndTracesEnabled": {
			input: map[string]interface{}{
				"logs": map[string]interface{}{
					"metrics_collected": map[string]interface{}{
						"app_signals": map[string]interface{}{},
					},
				},
				"traces": map[string]interface{}{
					"traces_collected": map[string]interface{}{
						"app_signals": map[string]interface{}{},
					},
				},
			},
			detector:       eksdetector.TestEKSDetector,
			isEKSDataStore: eksdetector.TestIsEKSCacheEKS,
		},
		"WithAppSignalsFallbackMultipleMetricsReceiversConfig": {
			input: map[string]interface{}{
				"logs": map[string]interface{}{
					"metrics_collected": map[string]interface{}{
						"app_signals": map[string]interface{}{},
						"cpu":         map[string]interface{}{},
					},
				},
				"traces": map[string]interface{}{
					"traces_collected": map[string]interface{}{
						"app_signals": map[string]interface{}{},
						"otlp":        map[string]interface{}{},
						"otlp2":       map[string]interface{}{},
					},
				},
			},
			detector:       eksdetector.TestEKSDetector,
			isEKSDataStore: eksdetector.TestIsEKSCacheEKS,
		},
		"WithAMPDestinationConfig": {
			input: map[string]interface{}{
				"metrics": map[string]interface{}{
					"metrics_destinations": map[string]interface{}{
						"amp": map[string]interface{}{
							"workspace_id": "ws-12345",
						},
					},
					"metrics_collected": map[string]interface{}{
						"cpu": map[string]interface{}{},
					},
				},
			},
		},
		"WithOutValidatePrometheusConfig": {
			input: map[string]interface{}{
				"metrics": map[string]interface{}{
					"metrics_destinations": map[string]interface{}{
						"amp": map[string]interface{}{
							"workspace_id": "ws-12345",
						},
					},
					"metrics_collected": map[string]interface{}{
						"prometheus": map[string]interface{}{
							"prometheus_config": "missing.yaml",
						},
					},
				},
			},
			wantErrContains: common.ConfigKey(prometheus.MetricsKey, common.PrometheusConfigPathKey),
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			eksdetector.NewDetector = testCase.detector
			eksdetector.IsEKS = testCase.isEKSDataStore
			translator.SetTargetPlatform("linux")
			got, err := Translate(testCase.input, "linux")
			if testCase.wantErrContains != "" {
				require.Error(t, err)
				assert.Nil(t, got)
				t.Log(err)
				assert.ErrorContains(t, err, testCase.wantErrContains)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

type testTranslator struct {
	id      pipeline.ID
	version int
}

func (t testTranslator) Translate(_ *confmap.Conf) (*common.ComponentTranslators, error) {
	return nil, nil
}

func (t testTranslator) ID() pipeline.ID {
	return t.id
}

var _ common.PipelineTranslator = (*testTranslator)(nil)

func TestRegisterPipeline(t *testing.T) {

	original := &testTranslator{id: pipeline.NewID(pipeline.SignalLogs), version: 1}
	tm := common.NewTranslatorMap[*common.ComponentTranslators](original)
	assert.Equal(t, 0, registry.Len())

	first := &testTranslator{id: pipeline.NewID(pipeline.SignalLogs), version: 2}
	second := &testTranslator{id: pipeline.NewID(pipeline.SignalLogs), version: 3}
	RegisterPipeline(first, second)
	assert.Equal(t, 1, registry.Len())

	tm.Merge(registry)
	got, ok := tm.Get(pipeline.NewID(pipeline.SignalLogs))
	assert.True(t, ok)
	assert.Equal(t, second.version, got.(*testTranslator).version)
	assert.NotEqual(t, first.version, got.(*testTranslator).version)
	assert.NotEqual(t, original.version, got.(*testTranslator).version)
}
