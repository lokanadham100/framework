package event

type functionEvent struct{
	startTime time.Time
	functionName string
	packageName string
	extra map[string]interface{}
}
