package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jConfig "github.com/uber/jaeger-client-go/config"
	"time"
)

func main() {
	cfg := jConfig.Configuration{
		Sampler: &jConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
		ServiceName: "go-maxms",
	}
	tracer, closer, err := cfg.NewTracer(jConfig.Logger(jaeger.StdLogger))
	defer closer.Close()
	if err != nil {
		panic(err)
	}

	parentSpan := tracer.StartSpan("order")

	cart := tracer.StartSpan("cart", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(time.Second * 2)
	cart.Finish()

	product := tracer.StartSpan("product", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(time.Second * 3)
	product.Finish()

	stock := tracer.StartSpan("product", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(time.Second * 1)
	stock.Finish()

	parentSpan.Finish()

}
