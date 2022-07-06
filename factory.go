package starterJaeger

import (
	"github.com/go-spring/spring-base/log"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	jaegerMetric "github.com/uber/jaeger-lib/metrics"
	"os"
	"os/signal"
	"syscall"
)

func newConfig(config JaegerConfig) *jaegerCfg.Configuration {
	cfg := &jaegerCfg.Configuration{
		ServiceName: config.ServiceName,
		Disabled:    config.Disabled,
		RPCMetrics:  config.RPCMetrics,
		Gen128Bit:   config.Gen128Bit,
		Tags:        config.Tags,
		Sampler: &jaegerCfg.SamplerConfig{
			Type:                     config.Sampler.Type,
			Param:                    config.Sampler.Param,
			SamplingServerURL:        config.Sampler.SamplingServerURL,
			SamplingRefreshInterval:  config.Sampler.SamplingRefreshInterval,
			MaxOperations:            config.Sampler.MaxOperations,
			OperationNameLateBinding: config.Sampler.OperationNameLateBinding,
			Options:                  config.Sampler.Options,
		},
		Reporter: &jaegerCfg.ReporterConfig{
			QueueSize:                  config.Reporter.QueueSize,
			BufferFlushInterval:        config.Reporter.BufferFlushInterval,
			LogSpans:                   config.Reporter.LogSpans,
			LocalAgentHostPort:         config.Reporter.LocalAgentHostPort,
			DisableAttemptReconnecting: config.Reporter.DisableAttemptReconnecting,
			AttemptReconnectInterval:   config.Reporter.AttemptReconnectInterval,
			CollectorEndpoint:          config.Reporter.CollectorEndpoint,
			User:                       config.Reporter.User,
			Password:                   config.Reporter.Password,
			HTTPHeaders:                config.Reporter.HTTPHeaders,
		},
		Headers: &jaeger.HeadersConfig{
			JaegerDebugHeader:        config.Headers.JaegerDebugHeader,
			JaegerBaggageHeader:      config.Headers.JaegerBaggageHeader,
			TraceContextHeaderName:   config.Headers.TraceContextHeaderName,
			TraceBaggageHeaderPrefix: config.Headers.TraceBaggageHeaderPrefix,
		},
	}

	if len(config.BaggageRestrictions.HostPort) > 0 {
		cfg.BaggageRestrictions = &jaegerCfg.BaggageRestrictionsConfig{
			DenyBaggageOnInitializationFailure: config.BaggageRestrictions.DenyBaggageOnInitializationFailure,
			HostPort:                           config.BaggageRestrictions.HostPort,
			RefreshInterval:                    config.BaggageRestrictions.RefreshInterval,
		}
	}

	if len(config.Throttler.HostPort) > 0 {
		cfg.Throttler = &jaegerCfg.ThrottlerConfig{
			HostPort:                  config.Throttler.HostPort,
			RefreshInterval:           config.Throttler.RefreshInterval,
			SynchronousInitialization: config.Throttler.SynchronousInitialization,
		}
	}
	return cfg
}

func NewClient(config JaegerConfig) (opentracing.Tracer, error) {
	cfg := newConfig(config)
	jLogger := jaeger.StdLogger
	jMetricsFactory := jaegerMetric.NullFactory
	tracer, cls, err := cfg.NewTracer(jaegerCfg.Logger(jLogger), jaegerCfg.Metrics(jMetricsFactory))
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)
	// 监听信号量
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-c // 阻塞等待
		log.Info("jaeger tracer will exit")
		_ = cls.Close() // 关闭
	}()
	return tracer, nil
}
