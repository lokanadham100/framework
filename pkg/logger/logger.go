package logger

import (
	"context"
	"fmt"
)

type Logger struct{
	logInterface LoggerInterface	
	ctx context.Context
}

func getLoggerWithContext(ctx context.Context) LoggerInterface{		
	lg := getOrCreate()
	lg.WithField("TraceID", getTraceID(ctx))
	return lg
}

func getLogger() LoggerInterface{
	return getOrCreate()
}