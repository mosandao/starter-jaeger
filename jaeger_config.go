package starterJaeger

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"time"
)

// JaegerConfig 注释 config.Configuration
type JaegerConfig struct {
	ServiceName         string                    `value:"${serviceName:=test_service}"`
	Disabled            bool                      `value:"${disabled:=false}"`
	RPCMetrics          bool                      `value:"${rpc_metrics:=false}"`
	Gen128Bit           bool                      `value:"${traceid_128bit:=false}"`
	Tags                []opentracing.Tag         `value:"${tags:=}"`
	Sampler             SamplerConfig             `value:"${sampler}"`
	Reporter            ReporterConfig            `value:"${reporter}"`
	Headers             HeadersConfig             `value:"${headers}"`
	BaggageRestrictions BaggageRestrictionsConfig `value:"${baggageRestrictions}"`
	Throttler           ThrottlerConfig           `value:"${throttler}"`
}

// SamplerConfig 注释 config.SamplerConfig
type SamplerConfig struct {
	Type                     string        `value:"${type:=}"`
	Param                    float64       `value:"${param:=0}"`
	SamplingServerURL        string        `value:"${samplingServerURL:=}"`
	SamplingRefreshInterval  time.Duration `value:"${samplingRefreshInterval:=0}"`
	MaxOperations            int           `value:"${maxOperations:=0}"`
	OperationNameLateBinding bool          `value:"${operationNameLateBinding:=false}"`
	Options                  []jaeger.SamplerOption
}

// ReporterConfig 注释 config.ReporterConfig
type ReporterConfig struct {
	QueueSize                  int               `value:"${queueSize:=0}"`
	BufferFlushInterval        time.Duration     `value:"${bufferFlushInterval:=0}"`
	LogSpans                   bool              `value:"${logSpans:=false}"`
	LocalAgentHostPort         string            `value:"${addr:=}"`
	DisableAttemptReconnecting bool              `value:"${disableAttemptReconnecting:=false}"`
	AttemptReconnectInterval   time.Duration     `value:"${attemptReconnectInterval:=0}"`
	CollectorEndpoint          string            `value:"${collectorEndpoint:=}"`
	User                       string            `value:"${user:=}"`
	Password                   string            `value:"${password:=}"`
	HTTPHeaders                map[string]string `value:"${http_headers:=}"`
}

// HeadersConfig 注释 config.HeadersConfig
type HeadersConfig struct {
	JaegerDebugHeader        string `value:"${jaegerDebugHeader:=}"`
	JaegerBaggageHeader      string `value:"${jaegerBaggageHeader:=}"`
	TraceContextHeaderName   string `value:"${TraceContextHeaderName:=}"`
	TraceBaggageHeaderPrefix string `value:"${traceBaggageHeaderPrefix:=}"`
}

// BaggageRestrictionsConfig 注释 config.BaggageRestrictionsConfig
type BaggageRestrictionsConfig struct {
	DenyBaggageOnInitializationFailure bool          `value:"${denyBaggageOnInitializationFailure:=false}"`
	HostPort                           string        `value:"${hostPort:=}"`
	RefreshInterval                    time.Duration `value:"${refreshInterval:=0}"`
}

type ThrottlerConfig struct {
	HostPort                  string        `value:"${hostPort:=}"`
	RefreshInterval           time.Duration `value:"${refreshInterval:=0}"`
	SynchronousInitialization bool          `value:"${synchronousInitialization:=false}"`
}
