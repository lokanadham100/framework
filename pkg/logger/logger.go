package logger

import (
	"context"
	"github.com/voonik/framework/pkg/config"
)

type Logger struct{
	loggerInterface	
	ctx context.Context
}

var log loggerInterface

func Init(){	
	logg := config.LogConfigLogger()
	log, _ = Get(logg)
}

func getLoggerWithContext(ctx context.Context) (loggerInterface) {		
	//TODO : Need to implement tracer as config.
	tid,sid,_ := getTraceAndSpanID(ctx)
	return log.WithField("TraceID", tid).WithField("SpanID", sid)
}