package event

type databaseEvent struct{
	startTime time.Time
	query string
	qtype string
	extra map[string]interface{}
}

func init(){
	RegisterEventWrapper("database", newDatabaseEvent)
}

func newDatabaseEvent(ctx context.Context, args ...interface{})(*databaseEvent, context.Context){
	return &databaseEvent{}
}

func (de *databaseEvent)Start(ctx context.Context, args ...interface{})(*databaseEvent, context.Context){
	// commenting for now as wwe are not using
	// de.parseArguments(arg...) // TODO : Use of goroutine for this call
	de.startTime = time.Now()	
	return de.startSpan(ctx)
}

func (de *databaseEvent)Push(ctx context.Context, args ...interface{}){

}

func (de *databaseEvent)Stop(ctx context.Context, args ...interface{}){
	de.parseArguments(arg...)
	de.stopSpan()
	de.Push(ctx, args...)
}

func (de *databaseEvent)startSpan(ctx context.Context)(*databaseEvent, context.Context){
	return de, ctx
}

func (de *databaseEvent)stopSpan(){

}

func (de *databaseEvent)parseArguments(args ...interface{})(){
	if len(args) > 0 {
		if v := args[0].qtype; len(v) > 0 {
			de.qtype = args[0].qtype
		}
		if v := args[0].query; len(v) > 0 {
			de.query = args[0].query
		}
	}
}