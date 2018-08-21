package event

import (
	"fmt"
	"time"
	"context"

	// "github.com/opentracing/opentracing-go"
	// "github.com/opentracing/opentracing-go/log"

	"github.com/voonik/framework/pkg/metrics"
	"github.com/voonik/framework/pkg/tracer"
)

type databaseEvent struct{
	startTime time.Time
	qtype string
	query string	
	extra map[string]interface{}
}

func init(){
	RegisterEventWrapper("database", newDatabaseEvent)
}

func newDatabaseEvent(ctx context.Context, args ...interface{})(WrapInterface, context.Context){
	return &databaseEvent{extra:make(map[string]interface{})}, nil
}

func (de *databaseEvent)Start(ctx context.Context, args ...interface{})(WrapInterface, context.Context){
	// commenting for now as wwe are not using
	// de.parseArguments(arg...) 
	// TODO : Use of goroutine for this call
	de.startTime = time.Now()	
	return de.startSpan(ctx)
}

func (de *databaseEvent)Push(ctx context.Context, args ...interface{}){
	metrics.DatabaseEventHistogram(de.qtype, de.query, time.Since(de.startTime).Seconds())
}

func (de *databaseEvent)Finish(ctx context.Context, args ...interface{}){
	de.parseArguments(args...)
	de.stopSpan()
	de.Push(ctx, args...)
}

func (de *databaseEvent)startSpan(ctx context.Context)(*databaseEvent, context.Context){
	span, ctx := tracer.StartSpanFromContext(ctx, "mysql")	
	de.extra["span"] = span
	return de, ctx
}

func (de *databaseEvent)stopSpan(){
	span := de.extra["span"].(tracer.SpanInterface)
	span.SetTag("qtype", de.qtype)
	span.SetTag("query", de.query)
	span.SetOperationName(fmt.Sprintf("mysql-%s",de.qtype))
	span.Finish()
}

func (de *databaseEvent)parseArguments(args ...interface{})(){	
	if len(args) > 0 {
		pargs := args[0].(map[string]string)
		if v := pargs["qtype"]; len(v) > 0 {
			de.qtype = pargs["qtype"]
		}
		if v := pargs["query"]; len(v) > 0 {
			de.query = pargs["query"]
		}
	}
}