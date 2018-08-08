package event

type processEvent struct{
	name string
	extra map[string]interface{}
}

func init(){
	RegisterEventWrapper("process", newProcessEvent)
}

func newProcessEvent(ctx context.Context, args ...interface{})(*processEvent, context.Context){
	return &processEvent{extra: make(map[string]interface{})}
}

func (pe *processEvent)Start(ctx context.Context, args ...interface{})(*processEvent, context.Context){
	pe.parseArguments(arg...) // TODO : Use of goroutine for this call
	pe.startTime = time.Now()
	return pe.startSpan(ctx)
}

func (fe *processEvent)Push(ctx context.Context, args ...interface{}){
	FunctionEventHistogram(pe.packageName, pe.functionName, time.Since(pe.startTime).Seconds())
}

func (fe *processEvent)Finish(ctx context.Context, args ...interface{}){
	pe.stopSpan()
	pe.Push(ctx, args...)
}

func (fe *processEvent)parseArguments(args ...interface{})(){	
		
}