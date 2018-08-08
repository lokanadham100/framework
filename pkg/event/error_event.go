package event

import (
	"context"

	"github.com/voonik/framework/pkg/metrics"	
)


type errorEvent struct{	
	packageName string
	functionName string	
	error string
	message string
	extra map[string]interface{}
}

func init(){
	RegisterPushWrapper("error", newErrorEvent)
}

func newErrorEvent()(PushInterface){
	return &errorEvent{}
}

func (ee *errorEvent)Push(ctx context.Context, args ...interface{}){
	metrics.ErrorEventCounter(ee.packageName, ee.functionName, ee.error, ee.message)	
}

func (ee *errorEvent)parseArguments(args ...interface{})(){	
		if v := args[0].(string); len(v) > 0 {
			ee.packageName = v
		}
		if v := args[1].(string); len(v) > 0 {
			ee.functionName = v
		}
		if v := args[2].(string); len(v) > 0 {
			ee.error = v
		}
		if v := args[3].(string); len(v) > 0 {
			ee.message = v
		}
}