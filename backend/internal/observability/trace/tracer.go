package trace

import "time"

type SpanKind string

const (
	SpanKindInternal SpanKind = "internal"
	SpanKindServer   SpanKind = "server"
	SpanKindClient   SpanKind = "client"
	SpanKindProducer SpanKind = "producer"
	SpanKindConsumer SpanKind = "consumer"
)

type Span struct {
	TraceID   string            `json:"trace_id"`
	SpanID    string            `json:"span_id"`
	ParentID  string            `json:"parent_id,omitempty"`
	Name      string            `json:"name"`
	Kind      SpanKind          `json:"kind"`
	StartTime time.Time         `json:"start_time"`
	EndTime   time.Time         `json:"end_time"`
	Duration  int64             `json:"duration_ms"`
	Attrs     map[string]string `json:"attributes"`
	Status    string            `json:"status"`
	Events    []SpanEvent       `json:"events,omitempty"`
}

type SpanEvent struct {
	Name      string            `json:"name"`
	Timestamp time.Time         `json:"timestamp"`
	Attrs     map[string]string `json:"attributes,omitempty"`
}

type TraceResult struct {
	TraceID    string    `json:"trace_id"`
	RootSpan   *Span     `json:"root_span"`
	Spans      []Span    `json:"spans"`
	Duration   int64     `json:"total_duration_ms"`
	TokenCount int       `json:"token_count"`
	Status     string    `json:"status"`
}

type Tracer struct {
	config interface{}
	logger interface{}
}

func NewTracer(config, logger interface{}) *Tracer {
	return &Tracer{config: config, logger: logger}
}

func (t *Tracer) StartSpan(name string, kind SpanKind, parent *Span) *Span {
	span := &Span{
		Name:      name,
		Kind:      kind,
		StartTime: time.Now(),
		Attrs:     make(map[string]string),
	}
	if parent != nil {
		span.ParentID = parent.SpanID
		span.TraceID = parent.TraceID
	}
	return span
}

func (t *Tracer) EndSpan(span *Span) {
	span.EndTime = time.Now()
	span.Duration = span.EndTime.Sub(span.StartTime).Milliseconds()
}