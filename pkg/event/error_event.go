package event

type errorEvent struct{	
	functionName string
	packageName string
	error string
	message string
	extra map[string]interface{}
}