package logger

import (
	"fmt"
	"context"	
)

// type LoggerContext interface{
// 	LoggerFromContext()
// 	ContextWithLogger()
// }

// func getTraceID(ctx context.Context) (string,error){
// 	return GetLogContext(ctx,"jaeger")
// }

// type LogTraceIdFunc func(context.Context) (string, error)

// var logContextRegistry = make(map[string]LogTraceIdFunc)

// func RegisterLogContext(name string, logTraceIdFunc LogTraceIdFunc){
// 	if _, ok := logContextRegistry[name]; ok {
// 		panic(fmt.Sprintf("%s is already registered", name))
// 	}
// 	logContextRegistry[name] = logTraceIdFunc
// }

// func GetLogContext(ctx context.Context, name string)(string, error){
// 	f, ok := logContextRegistry[name]
// 	if !ok {
// 		return "", fmt.Errorf("logger %q not found", name)
// 	}
// 	return f(ctx)
// }

func getTraceAndSpanID(ctx context.Context) (string,string,error){
	return "Trace","Span",nil
}