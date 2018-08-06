package logger

import (
	"context"
	_ "github.com/voonik/framework/pkg/config"
)

type Logger struct{
	loggerInterface	
	ctx context.Context
}

var log loggerInterface

func Init(){	
	log, _ = Get("logrus")
}

func getLoggerWithContext(ctx context.Context) (loggerInterface) {		
	tid,sid,_ := getTraceAndSpanID(ctx)
	return log.WithField("TraceID", tid).WithField("SpanID", sid)
}