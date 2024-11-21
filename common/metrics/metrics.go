package metrics

import "time"

// Metrics provides a convenient interface for reporting metrics.
type Metrics interface {
	// Start starts the metrics server.
	Start() error

	// Stop stops the metrics server.
	Stop() error

	// NewLatencyMetric creates a new LatencyMetric instance. Useful for reporting the latency of an operation.
	// Metric name and label may only contain alphanumeric characters and underscores.
	NewLatencyMetric(name string, label string, quantiles ...*Quantile) (LatencyMetric, error)

	// NewCountMetric creates a new CountMetric instance. Useful for tracking the count of a type of event.
	// Metric name and label may only contain alphanumeric characters and underscores.
	NewCountMetric(name string, label string) (CountMetric, error)

	// NewGaugeMetric creates a new GaugeMetric instance. Useful for reporting specific values.
	// Metric name and label may only contain alphanumeric characters and underscores.
	NewGaugeMetric(name string, label string) (GaugeMetric, error)

	// NewAutoGauge creates a new GaugeMetric instance that is automatically updated by the given source function.
	// The function is polled at the given period. This produces a gauge type metric internally.
	// Metric name and label may only contain alphanumeric characters and underscores.
	NewAutoGauge(name string, label string, pollPeriod time.Duration, source func() float64) error
}

// Metric represents a metric that can be reported.
type Metric interface {

	// Name returns the name of the metric.
	Name() string

	// Label returns the label of the metric. Metrics without a label will return an empty string.
	Label() string

	// Enabled returns true if the metric is enabled.
	Enabled() bool
}

// TODO can we require units for gauges?

// GaugeMetric allows specific values to be reported.
type GaugeMetric interface {
	Metric

	// Set sets the value of a gauge metric.
	Set(value float64)
}

// CountMetric allows the count of a type of event to be tracked.
type CountMetric interface {
	Metric

	// Increment increments the count by 1.
	Increment()

	// Add increments the count by the given value.
	Add(value float64)
}

// Quantile describes a quantile of a latency metric that should be reported. For a description of how
// to interpret a quantile, see the prometheus documentation
// https://github.com/prometheus/client_golang/blob/v1.20.5/prometheus/summary.go#L126
type Quantile struct {
	Quantile float64
	Error    float64
}

// NewQuantile creates a new Quantile instance. Error is set to 1% of the quantile.
func NewQuantile(quantile float64) *Quantile {
	return &Quantile{
		Quantile: quantile,
		Error:    quantile / 100.0,
	}
}

// LatencyMetric allows the latency of an operation to be tracked. Similar to a gauge metric, but specialized for time.
type LatencyMetric interface {
	Metric

	// ReportLatency reports a latency value.
	ReportLatency(latency time.Duration)
}