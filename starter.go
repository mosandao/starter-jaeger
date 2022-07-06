package starterJaeger

import (
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/gs/cond"
	"github.com/opentracing/opentracing-go"
)

func init() {
	gs.Provide(NewClient, "${jaeger}").
		Name("opentracing").
		On(cond.OnMissingBean(gs.BeanID((*opentracing.Tracer)(nil), "opentracing")))
}
