package logger

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

type JaegerConfig struct {
	ServiceName string
	Host        string
	LogSpans    bool
}

// New Jaeger client
func NewJaegerTracer(cfg *JaegerConfig) (opentracing.Tracer, io.Closer, error) {
	jaegerCfgInstance := jaegercfg.Configuration{
		ServiceName: cfg.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           cfg.LogSpans,
			LocalAgentHostPort: cfg.Host,
		},
	}
	jLogger := log.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := jaegerCfgInstance.NewTracer(config.Logger(jLogger), config.Metrics(jMetricsFactory))
	if err == nil {
		opentracing.SetGlobalTracer(tracer)
	}

	return tracer, closer, err
}
