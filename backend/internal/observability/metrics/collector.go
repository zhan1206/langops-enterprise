package metrics

import "sync/atomic"

type MetricType string

const (
	MetricCounter  MetricType = "counter"
	MetricGauge    MetricType = "gauge"
	MetricHistogram MetricType = "histogram"
)

type Metric struct {
	Name  string     json:"name"
	Type  MetricType json:"type"
	Value float64    json:"value"
	Labels map[string]string json:"labels,omitempty"
}

type Collector struct {
	config  interface{}
	logger  interface{}
	counters map[string]*atomic.Int64
	gauges   map[string]*atomic.Float64
}

func NewCollector(config, logger interface{}) *Collector {
	return &Collector{
		config:   config,
		logger:   logger,
		counters: make(map[string]*atomic.Int64),
		gauges:   make(map[string]*atomic.Float64),
	}
}

func (c *Collector) Inc(name string, labels map[string]string) {
	key := name
	if v, ok := c.counters[key]; ok {
		v.Add(1)
	} else {
		var val atomic.Int64
		val.Add(1)
		c.counters[key] = &val
	}
}

func (c *Collector) Set(name string, value float64, labels map[string]string) {
	key := name
	if v, ok := c.gauges[key]; ok {
		v.Store(value)
	} else {
		var val atomic.Float64
		val.Store(value)
		c.gauges[key] = &val
	}
}

func (c *Collector) GetAll() []Metric {
	var metrics []Metric
	for k, v := range c.counters {
		metrics = append(metrics, Metric{Name: k, Type: MetricCounter, Value: float64(v.Load())})
	}
	for k, v := range c.gauges {
		metrics = append(metrics, Metric{Name: k, Type: MetricGauge, Value: v.Load()})
	}
	return metrics
}