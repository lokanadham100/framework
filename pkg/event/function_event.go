package event

type functionEvent struct{
	startTime time.Time
	functionName string
	packageName string
	extra map[string]interface{}
}

func init(){
	RegisterEventWrapper("function", newFunctionEvent)
}

func newFunctionEvent(ctx context.Context, args ...interface{})(*functionEvent, context.Context){
	return &functionEvent{extra: make(map[string]interface{})}
}

func (fe *functionEvent)Start(ctx context.Context, args ...interface{})(*functionEvent, context.Context){
	fe.parseArguments(arg...) // TODO : Use of goroutine for this call
	fe.startTime = time.Now()	
	return fe.startSpan(ctx)
}

func (fe *functionEvent)Push(ctx context.Context, args ...interface{}){
	FunctionEventHistogram(fe.packageName, fe.functionName, time.Since(fe.startTime).Seconds())
}

func (fe *functionEvent)Finish(ctx context.Context, args ...interface{}){
	fe.stopSpan()
	fe.Push(ctx, args...)
}

func (fe *functionEvent)startSpan(ctx context.Context)(*functionEvent, context.Context){
	span, ctx := opentracing.StartSpanFromContext(ctx, fe.functionName)
	span.LogFields(
		log.String("functionName", fe.functionName),
		log.String("packageName", fe.packageName)
	)
	fe.extra["span"] = span
	return fe, ctx
}

func (fe *functionEvent)stopSpan(){
	fe.extra["span"].Finish()
}

func (fe *functionEvent)parseArguments(args ...interface{})(){	
		if v := args[0]; len(v) > 0 {
			fe.packageName = args[0]
		}
		if v := args[1]; len(v) > 0 {
			fe.functionName = args[1]
		}
}