// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/filter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

type metricDiskioEbsEc2InstancePerformanceExceededIops struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills diskio_ebs_ec2_instance_performance_exceeded_iops metric with initial data.
func (m *metricDiskioEbsEc2InstancePerformanceExceededIops) init() {
	m.data.SetName("diskio_ebs_ec2_instance_performance_exceeded_iops")
	m.data.SetDescription("The total time, in microseconds, that the EBS volume exceeded the attached Amazon EC2 instance's maximum IOPS performance")
	m.data.SetUnit("us")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricDiskioEbsEc2InstancePerformanceExceededIops) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiskioEbsEc2InstancePerformanceExceededIops) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiskioEbsEc2InstancePerformanceExceededIops) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiskioEbsEc2InstancePerformanceExceededIops(cfg MetricConfig) metricDiskioEbsEc2InstancePerformanceExceededIops {
	m := metricDiskioEbsEc2InstancePerformanceExceededIops{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiskioEbsEc2InstancePerformanceExceededTp struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills diskio_ebs_ec2_instance_performance_exceeded_tp metric with initial data.
func (m *metricDiskioEbsEc2InstancePerformanceExceededTp) init() {
	m.data.SetName("diskio_ebs_ec2_instance_performance_exceeded_tp")
	m.data.SetDescription("The total time, in microseconds, that the EBS volume exceeded the attached Amazon EC2 instance's maximum throughput performance")
	m.data.SetUnit("us")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricDiskioEbsEc2InstancePerformanceExceededTp) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiskioEbsEc2InstancePerformanceExceededTp) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiskioEbsEc2InstancePerformanceExceededTp) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiskioEbsEc2InstancePerformanceExceededTp(cfg MetricConfig) metricDiskioEbsEc2InstancePerformanceExceededTp {
	m := metricDiskioEbsEc2InstancePerformanceExceededTp{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiskioEbsTotalReadBytes struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills diskio_ebs_total_read_bytes metric with initial data.
func (m *metricDiskioEbsTotalReadBytes) init() {
	m.data.SetName("diskio_ebs_total_read_bytes")
	m.data.SetDescription("The total number of read bytes transferred")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricDiskioEbsTotalReadBytes) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiskioEbsTotalReadBytes) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiskioEbsTotalReadBytes) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiskioEbsTotalReadBytes(cfg MetricConfig) metricDiskioEbsTotalReadBytes {
	m := metricDiskioEbsTotalReadBytes{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiskioEbsTotalReadOps struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills diskio_ebs_total_read_ops metric with initial data.
func (m *metricDiskioEbsTotalReadOps) init() {
	m.data.SetName("diskio_ebs_total_read_ops")
	m.data.SetDescription("The total number of completed read operations")
	m.data.SetUnit("1")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricDiskioEbsTotalReadOps) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiskioEbsTotalReadOps) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiskioEbsTotalReadOps) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiskioEbsTotalReadOps(cfg MetricConfig) metricDiskioEbsTotalReadOps {
	m := metricDiskioEbsTotalReadOps{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiskioEbsTotalReadTime struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills diskio_ebs_total_read_time metric with initial data.
func (m *metricDiskioEbsTotalReadTime) init() {
	m.data.SetName("diskio_ebs_total_read_time")
	m.data.SetDescription("The total time spent, in microseconds, by all completed read operations")
	m.data.SetUnit("us")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricDiskioEbsTotalReadTime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiskioEbsTotalReadTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiskioEbsTotalReadTime) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiskioEbsTotalReadTime(cfg MetricConfig) metricDiskioEbsTotalReadTime {
	m := metricDiskioEbsTotalReadTime{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiskioEbsTotalWriteBytes struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills diskio_ebs_total_write_bytes metric with initial data.
func (m *metricDiskioEbsTotalWriteBytes) init() {
	m.data.SetName("diskio_ebs_total_write_bytes")
	m.data.SetDescription("The total number of write bytes transferred")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricDiskioEbsTotalWriteBytes) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiskioEbsTotalWriteBytes) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiskioEbsTotalWriteBytes) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiskioEbsTotalWriteBytes(cfg MetricConfig) metricDiskioEbsTotalWriteBytes {
	m := metricDiskioEbsTotalWriteBytes{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiskioEbsTotalWriteOps struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills diskio_ebs_total_write_ops metric with initial data.
func (m *metricDiskioEbsTotalWriteOps) init() {
	m.data.SetName("diskio_ebs_total_write_ops")
	m.data.SetDescription("The total number of completed write operations")
	m.data.SetUnit("1")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricDiskioEbsTotalWriteOps) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiskioEbsTotalWriteOps) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiskioEbsTotalWriteOps) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiskioEbsTotalWriteOps(cfg MetricConfig) metricDiskioEbsTotalWriteOps {
	m := metricDiskioEbsTotalWriteOps{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiskioEbsTotalWriteTime struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills diskio_ebs_total_write_time metric with initial data.
func (m *metricDiskioEbsTotalWriteTime) init() {
	m.data.SetName("diskio_ebs_total_write_time")
	m.data.SetDescription("The total time spent, in microseconds, by all completed write operations")
	m.data.SetUnit("us")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricDiskioEbsTotalWriteTime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiskioEbsTotalWriteTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiskioEbsTotalWriteTime) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiskioEbsTotalWriteTime(cfg MetricConfig) metricDiskioEbsTotalWriteTime {
	m := metricDiskioEbsTotalWriteTime{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiskioEbsVolumePerformanceExceededIops struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills diskio_ebs_volume_performance_exceeded_iops metric with initial data.
func (m *metricDiskioEbsVolumePerformanceExceededIops) init() {
	m.data.SetName("diskio_ebs_volume_performance_exceeded_iops")
	m.data.SetDescription("The total time, in microseconds, that IOPS demand exceeded the volume's provisioned IOPS performance")
	m.data.SetUnit("us")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricDiskioEbsVolumePerformanceExceededIops) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiskioEbsVolumePerformanceExceededIops) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiskioEbsVolumePerformanceExceededIops) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiskioEbsVolumePerformanceExceededIops(cfg MetricConfig) metricDiskioEbsVolumePerformanceExceededIops {
	m := metricDiskioEbsVolumePerformanceExceededIops{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiskioEbsVolumePerformanceExceededTp struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills diskio_ebs_volume_performance_exceeded_tp metric with initial data.
func (m *metricDiskioEbsVolumePerformanceExceededTp) init() {
	m.data.SetName("diskio_ebs_volume_performance_exceeded_tp")
	m.data.SetDescription("The total time, in microseconds, that throughput demand exceeded the volume's provisioned throughput performance")
	m.data.SetUnit("us")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricDiskioEbsVolumePerformanceExceededTp) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiskioEbsVolumePerformanceExceededTp) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiskioEbsVolumePerformanceExceededTp) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiskioEbsVolumePerformanceExceededTp(cfg MetricConfig) metricDiskioEbsVolumePerformanceExceededTp {
	m := metricDiskioEbsVolumePerformanceExceededTp{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiskioEbsVolumeQueueLength struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills diskio_ebs_volume_queue_length metric with initial data.
func (m *metricDiskioEbsVolumeQueueLength) init() {
	m.data.SetName("diskio_ebs_volume_queue_length")
	m.data.SetDescription("The number of read and write operations waiting to be completed")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricDiskioEbsVolumeQueueLength) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiskioEbsVolumeQueueLength) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiskioEbsVolumeQueueLength) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiskioEbsVolumeQueueLength(cfg MetricConfig) metricDiskioEbsVolumeQueueLength {
	m := metricDiskioEbsVolumeQueueLength{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                                            MetricsBuilderConfig // config of the metrics builder.
	startTime                                         pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                                   int                  // maximum observed number of metrics per resource.
	metricsBuffer                                     pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                                         component.BuildInfo  // contains version information.
	resourceAttributeIncludeFilter                    map[string]filter.Filter
	resourceAttributeExcludeFilter                    map[string]filter.Filter
	metricDiskioEbsEc2InstancePerformanceExceededIops metricDiskioEbsEc2InstancePerformanceExceededIops
	metricDiskioEbsEc2InstancePerformanceExceededTp   metricDiskioEbsEc2InstancePerformanceExceededTp
	metricDiskioEbsTotalReadBytes                     metricDiskioEbsTotalReadBytes
	metricDiskioEbsTotalReadOps                       metricDiskioEbsTotalReadOps
	metricDiskioEbsTotalReadTime                      metricDiskioEbsTotalReadTime
	metricDiskioEbsTotalWriteBytes                    metricDiskioEbsTotalWriteBytes
	metricDiskioEbsTotalWriteOps                      metricDiskioEbsTotalWriteOps
	metricDiskioEbsTotalWriteTime                     metricDiskioEbsTotalWriteTime
	metricDiskioEbsVolumePerformanceExceededIops      metricDiskioEbsVolumePerformanceExceededIops
	metricDiskioEbsVolumePerformanceExceededTp        metricDiskioEbsVolumePerformanceExceededTp
	metricDiskioEbsVolumeQueueLength                  metricDiskioEbsVolumeQueueLength
}

// MetricBuilderOption applies changes to default metrics builder.
type MetricBuilderOption interface {
	apply(*MetricsBuilder)
}

type metricBuilderOptionFunc func(mb *MetricsBuilder)

func (mbof metricBuilderOptionFunc) apply(mb *MetricsBuilder) {
	mbof(mb)
}

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) MetricBuilderOption {
	return metricBuilderOptionFunc(func(mb *MetricsBuilder) {
		mb.startTime = startTime
	})
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.Settings, options ...MetricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:        mbc,
		startTime:     pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer: pmetric.NewMetrics(),
		buildInfo:     settings.BuildInfo,
		metricDiskioEbsEc2InstancePerformanceExceededIops: newMetricDiskioEbsEc2InstancePerformanceExceededIops(mbc.Metrics.DiskioEbsEc2InstancePerformanceExceededIops),
		metricDiskioEbsEc2InstancePerformanceExceededTp:   newMetricDiskioEbsEc2InstancePerformanceExceededTp(mbc.Metrics.DiskioEbsEc2InstancePerformanceExceededTp),
		metricDiskioEbsTotalReadBytes:                     newMetricDiskioEbsTotalReadBytes(mbc.Metrics.DiskioEbsTotalReadBytes),
		metricDiskioEbsTotalReadOps:                       newMetricDiskioEbsTotalReadOps(mbc.Metrics.DiskioEbsTotalReadOps),
		metricDiskioEbsTotalReadTime:                      newMetricDiskioEbsTotalReadTime(mbc.Metrics.DiskioEbsTotalReadTime),
		metricDiskioEbsTotalWriteBytes:                    newMetricDiskioEbsTotalWriteBytes(mbc.Metrics.DiskioEbsTotalWriteBytes),
		metricDiskioEbsTotalWriteOps:                      newMetricDiskioEbsTotalWriteOps(mbc.Metrics.DiskioEbsTotalWriteOps),
		metricDiskioEbsTotalWriteTime:                     newMetricDiskioEbsTotalWriteTime(mbc.Metrics.DiskioEbsTotalWriteTime),
		metricDiskioEbsVolumePerformanceExceededIops:      newMetricDiskioEbsVolumePerformanceExceededIops(mbc.Metrics.DiskioEbsVolumePerformanceExceededIops),
		metricDiskioEbsVolumePerformanceExceededTp:        newMetricDiskioEbsVolumePerformanceExceededTp(mbc.Metrics.DiskioEbsVolumePerformanceExceededTp),
		metricDiskioEbsVolumeQueueLength:                  newMetricDiskioEbsVolumeQueueLength(mbc.Metrics.DiskioEbsVolumeQueueLength),
		resourceAttributeIncludeFilter:                    make(map[string]filter.Filter),
		resourceAttributeExcludeFilter:                    make(map[string]filter.Filter),
	}
	if mbc.ResourceAttributes.VolumeID.MetricsInclude != nil {
		mb.resourceAttributeIncludeFilter["VolumeId"] = filter.CreateFilter(mbc.ResourceAttributes.VolumeID.MetricsInclude)
	}
	if mbc.ResourceAttributes.VolumeID.MetricsExclude != nil {
		mb.resourceAttributeExcludeFilter["VolumeId"] = filter.CreateFilter(mbc.ResourceAttributes.VolumeID.MetricsExclude)
	}

	for _, op := range options {
		op.apply(mb)
	}
	return mb
}

// NewResourceBuilder returns a new resource builder that should be used to build a resource associated with for the emitted metrics.
func (mb *MetricsBuilder) NewResourceBuilder() *ResourceBuilder {
	return NewResourceBuilder(mb.config.ResourceAttributes)
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption interface {
	apply(pmetric.ResourceMetrics)
}

type resourceMetricsOptionFunc func(pmetric.ResourceMetrics)

func (rmof resourceMetricsOptionFunc) apply(rm pmetric.ResourceMetrics) {
	rmof(rm)
}

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	})
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	})
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(options ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("github.com/aws/amazon-cloudwatch-agent/receiver/awsebsnvmereceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricDiskioEbsEc2InstancePerformanceExceededIops.emit(ils.Metrics())
	mb.metricDiskioEbsEc2InstancePerformanceExceededTp.emit(ils.Metrics())
	mb.metricDiskioEbsTotalReadBytes.emit(ils.Metrics())
	mb.metricDiskioEbsTotalReadOps.emit(ils.Metrics())
	mb.metricDiskioEbsTotalReadTime.emit(ils.Metrics())
	mb.metricDiskioEbsTotalWriteBytes.emit(ils.Metrics())
	mb.metricDiskioEbsTotalWriteOps.emit(ils.Metrics())
	mb.metricDiskioEbsTotalWriteTime.emit(ils.Metrics())
	mb.metricDiskioEbsVolumePerformanceExceededIops.emit(ils.Metrics())
	mb.metricDiskioEbsVolumePerformanceExceededTp.emit(ils.Metrics())
	mb.metricDiskioEbsVolumeQueueLength.emit(ils.Metrics())

	for _, op := range options {
		op.apply(rm)
	}
	for attr, filter := range mb.resourceAttributeIncludeFilter {
		if val, ok := rm.Resource().Attributes().Get(attr); ok && !filter.Matches(val.AsString()) {
			return
		}
	}
	for attr, filter := range mb.resourceAttributeExcludeFilter {
		if val, ok := rm.Resource().Attributes().Get(attr); ok && filter.Matches(val.AsString()) {
			return
		}
	}

	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(options ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(options...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordDiskioEbsEc2InstancePerformanceExceededIopsDataPoint adds a data point to diskio_ebs_ec2_instance_performance_exceeded_iops metric.
func (mb *MetricsBuilder) RecordDiskioEbsEc2InstancePerformanceExceededIopsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricDiskioEbsEc2InstancePerformanceExceededIops.recordDataPoint(mb.startTime, ts, val)
}

// RecordDiskioEbsEc2InstancePerformanceExceededTpDataPoint adds a data point to diskio_ebs_ec2_instance_performance_exceeded_tp metric.
func (mb *MetricsBuilder) RecordDiskioEbsEc2InstancePerformanceExceededTpDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricDiskioEbsEc2InstancePerformanceExceededTp.recordDataPoint(mb.startTime, ts, val)
}

// RecordDiskioEbsTotalReadBytesDataPoint adds a data point to diskio_ebs_total_read_bytes metric.
func (mb *MetricsBuilder) RecordDiskioEbsTotalReadBytesDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricDiskioEbsTotalReadBytes.recordDataPoint(mb.startTime, ts, val)
}

// RecordDiskioEbsTotalReadOpsDataPoint adds a data point to diskio_ebs_total_read_ops metric.
func (mb *MetricsBuilder) RecordDiskioEbsTotalReadOpsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricDiskioEbsTotalReadOps.recordDataPoint(mb.startTime, ts, val)
}

// RecordDiskioEbsTotalReadTimeDataPoint adds a data point to diskio_ebs_total_read_time metric.
func (mb *MetricsBuilder) RecordDiskioEbsTotalReadTimeDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricDiskioEbsTotalReadTime.recordDataPoint(mb.startTime, ts, val)
}

// RecordDiskioEbsTotalWriteBytesDataPoint adds a data point to diskio_ebs_total_write_bytes metric.
func (mb *MetricsBuilder) RecordDiskioEbsTotalWriteBytesDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricDiskioEbsTotalWriteBytes.recordDataPoint(mb.startTime, ts, val)
}

// RecordDiskioEbsTotalWriteOpsDataPoint adds a data point to diskio_ebs_total_write_ops metric.
func (mb *MetricsBuilder) RecordDiskioEbsTotalWriteOpsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricDiskioEbsTotalWriteOps.recordDataPoint(mb.startTime, ts, val)
}

// RecordDiskioEbsTotalWriteTimeDataPoint adds a data point to diskio_ebs_total_write_time metric.
func (mb *MetricsBuilder) RecordDiskioEbsTotalWriteTimeDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricDiskioEbsTotalWriteTime.recordDataPoint(mb.startTime, ts, val)
}

// RecordDiskioEbsVolumePerformanceExceededIopsDataPoint adds a data point to diskio_ebs_volume_performance_exceeded_iops metric.
func (mb *MetricsBuilder) RecordDiskioEbsVolumePerformanceExceededIopsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricDiskioEbsVolumePerformanceExceededIops.recordDataPoint(mb.startTime, ts, val)
}

// RecordDiskioEbsVolumePerformanceExceededTpDataPoint adds a data point to diskio_ebs_volume_performance_exceeded_tp metric.
func (mb *MetricsBuilder) RecordDiskioEbsVolumePerformanceExceededTpDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricDiskioEbsVolumePerformanceExceededTp.recordDataPoint(mb.startTime, ts, val)
}

// RecordDiskioEbsVolumeQueueLengthDataPoint adds a data point to diskio_ebs_volume_queue_length metric.
func (mb *MetricsBuilder) RecordDiskioEbsVolumeQueueLengthDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricDiskioEbsVolumeQueueLength.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...MetricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op.apply(mb)
	}
}
