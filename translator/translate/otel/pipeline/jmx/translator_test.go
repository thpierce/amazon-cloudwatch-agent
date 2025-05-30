// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package jmx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pipeline"

	"github.com/aws/amazon-cloudwatch-agent/internal/util/collections"
	"github.com/aws/amazon-cloudwatch-agent/translator/context"
	"github.com/aws/amazon-cloudwatch-agent/translator/translate/otel/common"
)

func TestTranslator(t *testing.T) {
	type want struct {
		pipelineID string
		receivers  []string
		processors []string
		exporters  []string
		extensions []string
	}
	testCases := map[string]struct {
		input       map[string]any
		index       int
		destination string
		isContainer bool
		want        *want
		wantErr     error
	}{
		"WithoutJMX": {
			input: map[string]any{},
			index: -1,
			wantErr: &common.MissingKeyError{
				ID:      pipeline.NewIDWithName(pipeline.SignalMetrics, "jmx"),
				JsonKey: "metrics::metrics_collected::jmx",
			},
		},
		"WithEmpty": {
			input: map[string]any{
				"metrics": map[string]any{
					"metrics_collected": map[string]any{
						"jmx": map[string]any{},
					},
				},
			},
			index: -1,
			wantErr: &common.MissingKeyError{
				ID:      pipeline.NewIDWithName(pipeline.SignalMetrics, "jmx"),
				JsonKey: "metrics::metrics_collected::jmx::<target-system>::measurement",
			},
		},
		"WithEmpty/Target": {
			input: map[string]any{
				"metrics": map[string]any{
					"metrics_collected": map[string]any{
						"jmx": map[string]any{
							"jvm": map[string]any{},
						},
					},
				},
			},
			index: -1,
			wantErr: &common.MissingKeyError{
				ID:      pipeline.NewIDWithName(pipeline.SignalMetrics, "jmx"),
				JsonKey: "metrics::metrics_collected::jmx::<target-system>::measurement",
			},
		},
		"WithEmpty/Measurement": {
			input: map[string]any{
				"metrics": map[string]any{
					"metrics_collected": map[string]any{
						"jmx": map[string]any{
							"jvm": map[string]any{
								"measurement": []any{
									"jvm.memory.heap.init",
								},
							},
							"tomcat": map[string]any{
								"measurement": []any{},
							},
						},
					},
				},
			},
			index: -1,
			wantErr: &common.MissingKeyError{
				ID:      pipeline.NewIDWithName(pipeline.SignalMetrics, "jmx"),
				JsonKey: "metrics::metrics_collected::jmx::<target-system>::measurement",
			},
		},
		"WithInvalidIndex": {
			input: map[string]any{
				"metrics": map[string]any{
					"metrics_collected": map[string]any{
						"jmx": []any{},
					},
				},
			},
			index: 1,
			wantErr: &common.MissingKeyError{
				ID:      pipeline.NewIDWithName(pipeline.SignalMetrics, "jmx/1"),
				JsonKey: "metrics::metrics_collected::jmx[1]::<target-system>::measurement",
			},
		},
		"WithValidJMX/Object": {
			input: map[string]any{
				"metrics": map[string]any{
					"metrics_collected": map[string]any{
						"jmx": map[string]any{
							"endpoint": "localhost:8080",
							"jvm": map[string]any{
								"measurement": []any{
									"jvm.memory.heap.init",
								},
							},
						},
					},
				},
			},
			index: -1,
			want: &want{
				pipelineID: "metrics/jmx",
				receivers:  []string{"jmx"},
				processors: []string{"filter/jmx", "resource/jmx", "cumulativetodelta/jmx"},
				exporters:  []string{"awscloudwatch"},
				extensions: []string{"agenthealth/metrics"},
			},
		},
		"WithValidJMX/Object/EKS": {
			input: map[string]any{
				"metrics": map[string]any{
					"metrics_collected": map[string]any{
						"jmx": map[string]any{
							"jvm": map[string]any{
								"measurement": []any{
									"jvm.memory.heap.init",
								},
							},
						},
					},
				},
			},
			index:       -1,
			isContainer: true,
			want: &want{
				pipelineID: "metrics/jmx",
				receivers:  []string{"otlp/jmx"},
				processors: []string{"filter/jmx", "metricstransform/jmx", "transform/jmx/drop", "cumulativetodelta/jmx"},
				exporters:  []string{"awscloudwatch"},
				extensions: []string{"agenthealth/metrics"},
			},
		},
		"WithValidJMX/Object/EKS/AppendDimensions": {
			input: map[string]any{
				"metrics": map[string]any{
					"metrics_collected": map[string]any{
						"jmx": map[string]any{
							"jvm": map[string]any{
								"measurement": []any{
									"jvm.memory.heap.init",
								},
							},
							"append_dimensions": map[string]any{
								"key": "value",
							},
						},
					},
				},
			},
			index:       -1,
			isContainer: true,
			want: &want{
				pipelineID: "metrics/jmx",
				receivers:  []string{"otlp/jmx"},
				processors: []string{"filter/jmx", "metricstransform/jmx", "resource/jmx", "transform/jmx/drop", "cumulativetodelta/jmx"},
				exporters:  []string{"awscloudwatch"},
				extensions: []string{"agenthealth/metrics"},
			},
		},
		"WithValidJMX/Object/AMP": {
			input: map[string]any{
				"metrics": map[string]any{
					"metrics_destinations": map[string]any{
						"amp": map[string]any{
							"workspace_id": "ws-12345",
						},
					},
					"metrics_collected": map[string]any{
						"jmx": map[string]any{
							"endpoint": "localhost:8080",
							"jvm": map[string]any{
								"measurement": []any{
									"jvm.memory.heap.init",
								},
							},
						},
					},
				},
			},
			index:       -1,
			destination: "amp",
			want: &want{
				pipelineID: "metrics/jmx/amp",
				receivers:  []string{"jmx"},
				processors: []string{"filter/jmx", "resource/jmx", "batch/jmx/amp", "deltatocumulative/jmx/amp"},
				exporters:  []string{"prometheusremotewrite/amp"},
				extensions: []string{"sigv4auth"},
			},
		},
		"WithValidJMX/Object/AMP/EKS": {
			input: map[string]any{
				"metrics": map[string]any{
					"metrics_destinations": map[string]any{
						"amp": map[string]any{
							"workspace_id": "ws-12345",
						},
					},
					"metrics_collected": map[string]any{
						"jmx": map[string]any{
							"jvm": map[string]any{
								"measurement": []any{
									"jvm.memory.heap.init",
								},
							},
						},
					},
				},
			},
			index:       -1,
			destination: "amp",
			isContainer: true,
			want: &want{
				pipelineID: "metrics/jmx/amp",
				receivers:  []string{"otlp/jmx"},
				processors: []string{"filter/jmx", "metricstransform/jmx", "transform/jmx/drop", "batch/jmx/amp", "deltatocumulative/jmx/amp"},
				exporters:  []string{"prometheusremotewrite/amp"},
				extensions: []string{"sigv4auth"},
			},
		},
		"WithValidJMX/Object/Decoration": {
			input: map[string]any{
				"metrics": map[string]any{
					"metrics_collected": map[string]any{
						"jmx": map[string]any{
							"endpoint": "localhost:8080",
							"jvm": map[string]any{
								"measurement": []any{
									map[string]any{
										"name":   "jvm.classes.loaded",
										"rename": "JVM.CLASSES.LOADED",
										"unit":   "Count",
									},
								},
							},
						},
					},
				},
			},
			index:       -1,
			destination: "cloudwatch",
			want: &want{
				pipelineID: "metrics/jmx/cloudwatch",
				receivers:  []string{"jmx"},
				processors: []string{"filter/jmx", "resource/jmx", "transform/jmx", "cumulativetodelta/jmx"},
				exporters:  []string{"awscloudwatch"},
				extensions: []string{"agenthealth/metrics"},
			},
		},
		"WithValidJMX/Array": {
			input: map[string]any{
				"metrics": map[string]any{
					"append_dimensions": map[string]any{
						"InstanceId": "${aws:InstanceId}",
					},
					"metrics_collected": map[string]any{
						"jmx": []any{
							map[string]any{
								"endpoint": "localhost:8080",
								"jvm": map[string]any{
									"measurement": []any{
										"jvm.memory.heap.init",
										map[string]any{
											"name":   "jvm.classes.loaded",
											"rename": "JVM.CLASSES.LOADED",
											"unit":   "Count",
										},
									},
								},
							},
						},
					},
				},
			},
			index: 0,
			want: &want{
				pipelineID: "metrics/jmx/0",
				receivers:  []string{"jmx/0"},
				processors: []string{"filter/jmx/0", "resource/jmx", "transform/jmx/0", "ec2tagger", "cumulativetodelta/jmx"},
				exporters:  []string{"awscloudwatch"},
				extensions: []string{"agenthealth/metrics"},
			},
		},
		"WithValidJMX/Array/AMP/EKS": {
			input: map[string]any{
				"metrics": map[string]any{
					"metrics_destinations": map[string]any{
						"amp": map[string]any{
							"workspace_id": "ws-12345",
						},
					},
					"metrics_collected": map[string]any{
						"jmx": []any{
							map[string]any{
								"jvm": map[string]any{
									"measurement": []any{
										"jvm.memory.heap.init",
										map[string]any{
											"name":   "jvm.classes.loaded",
											"rename": "JVM.CLASSES.LOADED",
											"unit":   "Count",
										},
									},
								},
							},
						},
					},
				},
			},
			index:       0,
			destination: "amp",
			isContainer: true,
			want: &want{
				pipelineID: "metrics/jmx/amp/0",
				receivers:  []string{"otlp/jmx"},
				processors: []string{"filter/jmx/0", "metricstransform/jmx", "transform/jmx/drop", "transform/jmx/0", "batch/jmx/amp/0", "deltatocumulative/jmx/amp/0"},
				exporters:  []string{"prometheusremotewrite/amp"},
				extensions: []string{"sigv4auth"},
			},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			context.CurrentContext().SetRunInContainer(testCase.isContainer)
			tt := NewTranslator(common.WithIndex(testCase.index), common.WithDestination(testCase.destination))
			conf := confmap.NewFromStringMap(testCase.input)
			got, err := tt.Translate(conf)
			require.Equal(t, testCase.wantErr, err)
			if testCase.want == nil {
				require.Nil(t, got)
			} else {
				require.NotNil(t, got)
				require.EqualValues(t, testCase.want.pipelineID, tt.ID().String())
				assert.Equal(t, testCase.want.receivers, collections.MapSlice(got.Receivers.Keys(), component.ID.String))
				assert.Equal(t, testCase.want.processors, collections.MapSlice(got.Processors.Keys(), component.ID.String))
				assert.Equal(t, testCase.want.exporters, collections.MapSlice(got.Exporters.Keys(), component.ID.String))
				assert.Equal(t, testCase.want.extensions, collections.MapSlice(got.Extensions.Keys(), component.ID.String))
			}
		})
	}
}
