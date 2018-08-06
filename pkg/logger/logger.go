package logger

import (
	"context"
	"github.com/voonik/framework/pkg/config"
)

var log Logger

type Logger struct{
	LoggerInterface	
	ctx context.Context
}

func Init(){
	if Config.logger == "" {
		Config.logger = "logrus"
	}
	log, _ = Get(Config.logger)
}

func getLoggerWithContext(ctx context.Context) (loggerInterface, error) {		
		
	// tid,_ := getTraceID(ctx)
	// lg.WithField("TraceID", tid)
	return log
}