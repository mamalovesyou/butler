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

var DefaultJaegerConfig = JaegerConfig{
	ServiceName: "butler",
	Addr:        "jaeger-agent:6831",
	LogSpans:    true,
	Disabled:    true,
}

type JaegerConfig struct {
	ServiceName string `env:"SERVICE_NAME"`
	Addr        string `env:"ADDR"`
	LogSpans    bool   `env:"LOG_SPANS"`
	Disabled    bool   `env:"DISABLED"`
}

// NewJaegerTracer set up a global opentracing.Tracer
func NewJaegerTracer(cfg *JaegerConfig) (opentracing.Tracer, io.Closer, error) {
	jaegerCfgInstance := jaegercfg.Configuration{
		Disabled:    cfg.Disabled,
		ServiceName: cfg.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           cfg.LogSpans,
			LocalAgentHostPort: cfg.Addr,
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
