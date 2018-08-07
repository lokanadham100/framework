package Event

import (
	"context"

)

type WrapInterface interface{
	PushInterface
	Start(context.Context, args ...interface{}) (WrapInterface, context.Context)	
	Stop(context.Context, args ...interface{})
}

type PushInterface interface{
	Push(context.Context, args ...interface{})
}

// For Wrapper
var wrapRegistry = make(map[string]wrapFunc)

type wrapFunc func(context.Context, args ...interface{}) (WrapInterface, error)

func RegisterEventWrapper(name string, wFunc wrapFunc){
	if _, ok := wrapRegistry[name]; ok {
		panic(fmt.Sprintf("%s is already registered", name))
	}
	wrapRegistry[name] = wFunc
}

func GetWrapEvent(name string, ctx context.Context, args ...interface{})(WrapInterface, error){
	f, ok := wrapRegistry[name]
	if !ok {
		return nil, fmt.Errorf("WrapInterface %q not found", name)
	}
	return f(ctx, args...)
}

// For Pusher
var pushRegistry = make(map[string]pushFunc)

type pushFunc func(context.Context, args ...interface{}) (PushInterface, error)

func RegisterEventPusher(name string, pFunc pushFunc){
	if _, ok := pushRegistry[name]; ok {
		panic(fmt.Sprintf("%s is already registered", name))
	}
	pushRegistry[name] = wFunc
}

func GetPushEvent(name string, ctx context.Context, args ...interface{})(PushInterface, error){
	f, ok := pushRegistry[name]
	if !ok {
		return nil, fmt.Errorf("PushInterface %q not found", name)
	}
	return f(ctx, args...)
}

//Developer friendly push API :)
//Usage: event.Push("error", ctx, "some issue")
func Push(name string, ctx context.Context, args ...interface{}){
	GetPushEvent(name).Push(ctx, args...)
}

//Developer friendly event API :)
//Usage: event.Start("function", ctx, "method name", "arg1")
func Start(name string, ctx context.Context, args ...interface{}) (WrapInterface, error){
	return GetWrapEvent(name, ctx, args...)
}