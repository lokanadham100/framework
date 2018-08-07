package tracer

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/voonik/framework/pkg/logger"
)

func Init(){	
    opentracing.SetGlobalTracer(tracer)
}