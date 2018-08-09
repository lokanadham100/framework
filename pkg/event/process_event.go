package event

import (
	"os"
	"context"

	"github.com/voonik/framework/pkg/logger"
	"github.com/voonik/framework/pkg/tracer"
	"github.com/voonik/framework/pkg/metrics"
)


type processEvent struct{
	name string
	extra map[string]interface{}
}

func init(){
	RegisterEventWrapper("process", newProcessEvent)
}

func newProcessEvent(ctx context.Context, args ...interface{})(WrapInterface, context.Context){
	callInits()
	checkAndSetEnv()
	return &processEvent{extra: make(map[string]interface{})}, nil
}

func (pe *processEvent)Start(ctx context.Context, args ...interface{})(WrapInterface, context.Context){
	return pe, ctx
}

func (fe *processEvent)Push(ctx context.Context, args ...interface{}){
	
}

func (fe *processEvent)Finish(ctx context.Context, args ...interface{}){
}

func (fe *processEvent)parseArguments(args ...interface{})(){	
		
}

func callInits(){
	logger.Init()
	metrics.Init()
	tracer.Init()
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