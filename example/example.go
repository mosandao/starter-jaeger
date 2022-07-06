package main

import (
	"context"
	"fmt"
	"github.com/go-spring/spring-base/conf"
	_ "github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/gs"
	_ "github.com/mosandao/starter-jaeger"
	"github.com/opentracing/opentracing-go"
	"time"
)

var (
	config = `
jaeger:
  serviceName: "test_app"
  sampler:
    type: "const"
    param: 1
  reporter:
    logSpans: false
    addr: "localhost:6831"`
)

type runner struct {
	Tracer opentracing.Tracer `inject:"?"`
}

func (r *runner) Run(ctx gs.Context) {
	gCtx := context.WithValue(ctx.Context(), "key", "key1")
	tracing(gCtx, r.Tracer)
}

func tracing(ctx context.Context, tracer opentracing.Tracer) context.Context {
	span, ctx2 := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "example")
	span.SetTag("test", ctx.Value("key").(string))
	span.Finish()
	return ctx2
}

func main() {
	p, _ := conf.Bytes([]byte(config), ".yaml")
	for _, key := range p.Keys() {
		gs.Property(key, p.Get(key))
	}
	gs.Object(&runner{}).Export((*gs.AppRunner)(nil))
	go func() {
		time.Sleep(time.Second * 5)
		gs.ShutDown()
	}()
	fmt.Printf("program exited %v\n", gs.Web(false).Run())

}
