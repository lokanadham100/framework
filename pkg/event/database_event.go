package event

type databaseEvent struct{
	startTime time.Time
	query string
	qtype string
	extra map[string]interface{}
}

func (de *databaseEvent)Start(ctx context.Context, args ...interface{})