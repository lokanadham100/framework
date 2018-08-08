package event

import (
	"context"
	"time"

	"github.com/voonik/framework/pkg/metrics"
	"github.com/voonik/framework/pkg/tracer"
)

type functionEvent struct{
	startTime time.Time
	functionName string
	packageName string
	extra map[string]interface{}
}

func init(){
	RegisterEventWrapper("function", newFunctionEvent)
}

func newFunctionEvent(ctx context.Context, args ...interface{})(WrapInterface, context.Context){
	return &functionEvent{extra: make(map[string]interface{})}, nil
}

func (fe *functionEvent)Start(ctx context.Context, args ...interface{})(WrapInterface, context.Context){
	fe.parseArguments(args...) // TODO : Use of goroutine for this call
	fe.startTime = time.Now()	
	return fe.startSpan(ctx)
}

func (fe *functionEvent)Push(ctx context.Context, args ...interface{}){
	metrics.FunctionEventHistogram(fe.packageName, fe.functionName, time.Since(fe.startTime).Seconds())
}

func (fe *functionEvent)Finish(ctx context.Context, args ...interface{}){
	fe.stopSpan()
	fe.Push(ctx, args...)
}

func (fe *functionEvent)startSpan(ctx context.Context)(*functionEvent, context.Context){
	span, ctx := tracer.StartSpanFromContext(ctx, fe.functionName)
	span.LogFields(
		tracer.String("functionName", fe.functionName),
		tracer.String("packageName", fe.packageName),
	)
	fe.extra["span"] = span
	return fe, ctx
}

func (fe *functionEvent)stopSpan(){
	span := fe.extra["span"].(tracer.SpanInterface)
	span.Finish()
}

func (fe *functionEvent)parseArguments(args ...interface{})(){	
		if v := args[0].(string); len(v) > 0 {
			fe.packageName = v
		}
		if v := args[1].(string); len(v) > 0 {
			fe.functionName = v
		}
}