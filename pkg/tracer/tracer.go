package tracer

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	// "github.com/voonik/framework/pkg/logger"
)

type SpanInterface interface{
	opentracing.Span
}

func Init(){	
    opentracing.SetGlobalTracer(JaegerTracer())
}

func StartSpanFromContext(ctx context.Context, name string) (SpanInterface, context.Context){
	return opentracing.StartSpanFromContext(ctx, name)
}

func String(key, val string) log.Field {
	return log.String(key, val)
}

func Bool(key string, val bool) log.Field {
	return log.Bool(key, val)
}

func Int(key string, val int) log.Field {
	return log.Int(key, val)
}

func Int32(key string, val int32) log.Field {
	return log.Int32(key, val)
}

func Int64(key string, val int64) log.Field {
	return log.Int64(key, val)
}

func Uint32(key string, val uint32) log.Field {
	return log.Uint32(key, val)
}

func Uint64(key string, val uint64) log.Field {
	return log.Uint64(key, val)
}

func Float32(key string, val float32) log.Field {
	return log.Float32(key, val)
}

func Float64(key string, val float64) log.Field {
	return log.Float64(key, val)
}

func Error(err error) log.Field {
	return log.Error(err)
}

func Object(key string, obj interface{}) log.Field {
	return log.Object(key, obj)
}