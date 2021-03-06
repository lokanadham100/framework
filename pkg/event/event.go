package event

import (
	"fmt"
	"context"

)

type WrapInterface interface{
	PushInterface
	Start(context.Context, ...interface{}) (WrapInterface, context.Context)	
	Finish(context.Context, ...interface{})
}

type PushInterface interface{
	Push(context.Context, ...interface{})
}

// For Wrapper
var wrapRegistry = make(map[string]wrapFunc)

type wrapFunc func(context.Context, ...interface{}) (WrapInterface, context.Context)

func RegisterEventWrapper(name string, wFunc wrapFunc){
	if _, ok := wrapRegistry[name]; ok {
		panic(fmt.Sprintf("%s is already registered", name))
	}
	wrapRegistry[name] = wFunc
}

func GetWrapEvent(name string, ctx context.Context, args ...interface{})(WrapInterface, context.Context){
	f, ok := wrapRegistry[name]
	if !ok {
		return nil, nil//fmt.Errorf("WrapInterface %q not found", name)
	}
	return f(ctx, args...)
}

// For Pusher
var pushRegistry = make(map[string]pushFunc)

type pushFunc func() (PushInterface)

func RegisterPushWrapper(name string, pFunc pushFunc){
	if _, ok := pushRegistry[name]; ok {
		panic(fmt.Sprintf("%s is already registered", name))
	}
	pushRegistry[name] = pFunc
}

func GetPushEvent(name string)(PushInterface){
	f, ok := pushRegistry[name]
	if !ok {
		return nil
	}
	return f()
}

//Developer friendly push API :)
//Usage: event.Push("error", ctx, "some issue")
func Push(name string, ctx context.Context, args ...interface{}){
	GetPushEvent(name).Push(ctx, args...)
}

//Developer friendly event API :)
//Usage: event.Start("function", ctx, "method name", "arg1")
func Start(name string, ctx context.Context, args ...interface{}) (WrapInterface, context.Context){
	if ev, ok := GetWrapEvent(name, ctx, args...); ok == nil{
		return ev.Start(ctx, args...)
	}else{
		return nil,ok
	}
}