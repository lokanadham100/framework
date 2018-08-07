package event

type errorEvent struct{	
	functionName string
	packageName string
	error string
	message string
	extra map[string]interface{}
}

func init(){
	RegisterPushWrapper("error", newErrorEvent)
}

func newErrorEvent(ctx context.Context, args ...interface{})(*errorEvent){
	return &errorEvent{}
}

func (ee *errorEvent)Push(ctx context.Context, args ...interface{}){
	
}

func (ee *errorEvent)parseArguments(args ...interface{})(){	
		if v := args[0]; len(v) > 0 {
			ee.packageName = args[0]
		}
		if v := args[1]; len(v) > 0 {
			ee.functionName = args[1]
		}
		if v := args[2]; len(v) > 0 {
			ee.error = args[2]
		}
		if v := args[3]; len(v) > 0 {
			ee.message = args[3]
		}
}