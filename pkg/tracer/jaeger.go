package trace

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
    jconfig "github.com/uber/jaeger-client-go/config"
    "github.com/voonik/framework/pkg/config"
)

func JaegerTracer() (opentracing.Tracer){
	cfg := jconfig.Configuration{
    Sampler: &jconfig.SamplerConfig{
        Type:  config.TraceConfigType(),
        Param: config.TraceConfigParam(),
    },
    Reporter: &jconfig.ReporterConfig{
        LogSpans:            true,
        BufferFlushInterval: 1 * time.Second,
        LocalAgentHostPort: getJaegerHostPort(),
    },
    }
    tracer, _, _ := cfg.New(
        config.TraceConfigServiceName(),
        jconfig.Logger(logger.getLoggerWithName("opentracing")),
    )
    return tracer
}

func getJaegerHostPort() string{
	return fmt.Sprintf("%s:%s",config.TraceConfigHost(),config.TraceConfigPort())
}