// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package metricstransformprocessor

import (
	_ "embed"
	"fmt"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/processor"

	"github.com/aws/amazon-cloudwatch-agent/internal/containerinsightscommon"
	"github.com/aws/amazon-cloudwatch-agent/translator/translate/otel/common"
	"github.com/aws/amazon-cloudwatch-agent/translator/translate/otel/receiver/awscontainerinsight"
)

//go:embed metricstransform_jmx_config.yaml
var metricTransformJmxConfig string

//go:embed appsignals_runtime_config.yaml
var appSignalsRuntimeConfig string

var metricDuplicateTypes = []string{
	containerinsightscommon.TypeGpuContainer,
	containerinsightscommon.TypeGpuPod,
	containerinsightscommon.TypeGpuNode,
}

var renameMapForDcgm = map[string]string{
	"DCGM_FI_DEV_GPU_UTIL":        containerinsightscommon.GpuUtilization,
	"DCGM_FI_DEV_FB_USED_PERCENT": containerinsightscommon.GpuMemUtilization,
	"DCGM_FI_DEV_FB_USED":         containerinsightscommon.GpuMemUsed,
	"DCGM_FI_DEV_FB_TOTAL":        containerinsightscommon.GpuMemTotal,
	"DCGM_FI_DEV_GPU_TEMP":        containerinsightscommon.GpuTemperature,
	"DCGM_FI_DEV_POWER_USAGE":     containerinsightscommon.GpuPowerDraw,
}

var renameMapForNeuronMonitor = map[string]string{
	"execution_errors_total":                          containerinsightscommon.NeuronExecutionErrors,
	"execution_status_total":                          containerinsightscommon.NeuronExecutionStatus,
	"neuron_runtime_memory_used_bytes":                containerinsightscommon.NeuronRuntimeMemoryUsage,
	"neuroncore_memory_usage_constants":               containerinsightscommon.NeuronCoreMemoryUtilizationConstants,
	"neuroncore_memory_usage_model_code":              containerinsightscommon.NeuronCoreMemoryUtilizationModelCode,
	"neuroncore_memory_usage_model_shared_scratchpad": containerinsightscommon.NeuronCoreMemoryUtilizationSharedScratchpad,
	"neuroncore_memory_usage_runtime_memory":          containerinsightscommon.NeuronCoreMemoryUtilizationRuntimeMemory,
	"neuroncore_memory_usage_tensors":                 containerinsightscommon.NeuronCoreMemoryUtilizationTensors,
	"neuroncore_utilization_ratio":                    containerinsightscommon.NeuronCoreUtilization,
	"instance_info":                                   containerinsightscommon.NeuronInstanceInfo,
	"neuron_hardware":                                 containerinsightscommon.NeuronHardware,
	"hardware_ecc_events_total":                       containerinsightscommon.NeuronDeviceHardwareEccEvents,
	"execution_latency_seconds":                       containerinsightscommon.NeuronExecutionLatency,
}

var renameMapForNvme = map[string]string{
	"aws_ebs_csi_read_ops_total":                  containerinsightscommon.NvmeReadOpsTotal,
	"aws_ebs_csi_write_ops_total":                 containerinsightscommon.NvmeWriteOpsTotal,
	"aws_ebs_csi_read_bytes_total":                containerinsightscommon.NvmeReadBytesTotal,
	"aws_ebs_csi_write_bytes_total":               containerinsightscommon.NvmeWriteBytesTotal,
	"aws_ebs_csi_read_seconds_total":              containerinsightscommon.NvmeReadTime,
	"aws_ebs_csi_write_seconds_total":             containerinsightscommon.NvmeWriteTime,
	"aws_ebs_csi_exceeded_iops_seconds_total":     containerinsightscommon.NvmeExceededIOPSTime,
	"aws_ebs_csi_exceeded_tp_seconds_total":       containerinsightscommon.NvmeExceededTPTime,
	"aws_ebs_csi_ec2_exceeded_iops_seconds_total": containerinsightscommon.NvmeExceededEC2IOPSTime,
	"aws_ebs_csi_ec2_exceeded_tp_seconds_total":   containerinsightscommon.NvmeExceededEC2TPTime,
	"aws_ebs_csi_volume_queue_length":             containerinsightscommon.NvmeVolumeQueueLength,
}

type translator struct {
	name    string
	factory processor.Factory
}

var _ common.ComponentTranslator = (*translator)(nil)

func NewTranslatorWithName(name string) common.ComponentTranslator {
	return &translator{name, metricstransformprocessor.NewFactory()}
}

func (t *translator) ID() component.ID {
	return component.NewIDWithName(t.factory.Type(), t.name)
}

func (t *translator) Translate(conf *confmap.Conf) (component.Config, error) {
	cfg := t.factory.CreateDefaultConfig().(*metricstransformprocessor.Config)
	if t.name == common.PipelineNameContainerInsightsJmx {
		return common.GetYamlFileToYamlConfig(cfg, metricTransformJmxConfig)
	} else if t.name == common.AppSignals {
		return common.GetYamlFileToYamlConfig(cfg, appSignalsRuntimeConfig)
	}

	var transformRules []map[string]interface{}
	if t.name == common.PipelineNameContainerInsights {
		transformRules = []map[string]interface{}{
			{
				"include":                   "apiserver_request_total",
				"match_type":                "regexp",
				"experimental_match_labels": map[string]string{"code": "^5.*"},
				"action":                    "insert",
				"new_name":                  "apiserver_request_total_5xx",
			},
		}

		if awscontainerinsight.EnhancedContainerInsightsEnabled(conf) {
			for oldNvmeMetric, newNvmeMetric := range renameMapForNvme {
				transformRules = append(transformRules, map[string]interface{}{
					"include":  oldNvmeMetric,
					"action":   "update",
					"new_name": containerinsightscommon.MetricName(containerinsightscommon.TypeNode, newNvmeMetric),
					"operations": []map[string]interface{}{{
						"action":    "add_label",
						"new_label": containerinsightscommon.MetricType,
						"new_value": containerinsightscommon.TypeNodeEBS,
					}},
				})
			}

		}

		if awscontainerinsight.AcceleratedComputeMetricsEnabled(conf) {
			// appends DCGM metric transform rules for each metric type (container/pod/node) with following format:
			// {
			//		"include":  "DCGM_FI_DEV_GPU_UTIL",
			//		"action":   "insert",
			//		"new_name": "container_gpu_utilization",
			//		"operations": [
			//     		{
			//				"action":   "add_label",
			//				"new_label": "Type",
			//				"new_value": "ContainerGPU",
			//			},
			//			<additional operations>...
			//      ]
			//	},
			for old, new := range renameMapForDcgm {
				var operations []map[string]interface{}
				// convert decimals to percent
				if new == containerinsightscommon.GpuMemUtilization {
					operations = append(operations, map[string]interface{}{
						"action":             "experimental_scale_value",
						"experimental_scale": 100,
					})
				} else if new == containerinsightscommon.GpuMemTotal || new == containerinsightscommon.GpuMemUsed {
					operations = append(operations, map[string]interface{}{
						"action":             "experimental_scale_value",
						"experimental_scale": 1024 * 1024,
					})
				}
				for _, t := range metricDuplicateTypes {
					transformRules = append(transformRules, map[string]interface{}{
						"include":  old,
						"action":   "insert",
						"new_name": containerinsightscommon.MetricName(t, new),
						"operations": append([]map[string]interface{}{
							{
								"action":    "add_label",
								"new_label": containerinsightscommon.MetricType,
								"new_value": t,
							},
						}, operations...),
					})
				}
			}

			for oldName, newName := range renameMapForNeuronMonitor {
				var operations []map[string]interface{}
				if newName == containerinsightscommon.NeuronCoreUtilization {
					operations = append(operations, map[string]interface{}{
						"action":             "experimental_scale_value",
						"experimental_scale": 100,
					})
				}

				transformRules = append(transformRules, map[string]interface{}{
					"include":  oldName,
					"action":   "update",
					"new_name": newName,
					"operations": append([]map[string]interface{}{},
						operations...),
				})
			}
		}
	} else if t.name == common.PipelineNameJmx {
		transformRules = []map[string]interface{}{
			{
				"include": "tomcat.sessions",
				"action":  "update",
				"operations": []map[string]interface{}{
					{
						"action":           "aggregate_labels",
						"aggregation_type": "sum",
					},
					{
						"action": "delete_label_value",
						"label":  "context",
					},
				},
			},
			{
				"include": "tomcat.rejected_sessions",
				"action":  "update",
				"operations": []map[string]interface{}{
					{
						"action":           "aggregate_labels",
						"aggregation_type": "sum",
					},
					{
						"action": "delete_label_value",
						"label":  "context",
					},
				},
			},
		}
	}

	if len(transformRules) == 0 {
		return nil, fmt.Errorf("no transform rules for %s", t.name)
	}

	c := confmap.NewFromStringMap(map[string]interface{}{
		"transforms": transformRules,
	})
	if err := c.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to unmarshal into metricstransform config: %w", err)
	}

	return cfg, nil
}
