package event

import (
	"os"
	"context"
)


type processEvent struct{
	name string
	extra map[string]interface{}
}

func init(){
	RegisterEventWrapper("process", newProcessEvent)
}

func newProcessEvent(ctx context.Context, args ...interface{})(WrapInterface, context.Context){
	checkAndSetEnv()
	return &processEvent{extra: make(map[string]interface{})}, nil
}

func (pe *processEvent)Start(ctx context.Context, args ...interface{})(WrapInterface, context.Context){
	// return pe.startSpan(ctx)
	return pe, ctx
}

func (fe *processEvent)Push(ctx context.Context, args ...interface{}){
	// metrics.FunctionEventHistogram(pe.packageName, pe.functionName, time.Since(pe.startTime).Seconds())
}

func (fe *processEvent)Finish(ctx context.Context, args ...interface{}){
	// pe.stopSpan()
	// pe.Push(ctx, args...)
}

func (fe *processEvent)parseArguments(args ...interface{})(){	
		
}

func checkAndSetEnv(){
	if env := os.Getenv("ENV"); env == ""{
		if env := os.Getenv("ENVIRONMENT"); env == ""{
			setEnv("development")
		}else{
			setEnv(env)
		}
	}else{
		setEnv(env)
	}
}

func setEnv(env string){
	os.Setenv("ENV", env)
	os.Setenv("ENVIRONMENT", env)
}