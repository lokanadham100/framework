package Event

type Function struct{
	startTime time.Time
	functionName string
	packageName string
	extra string
}

type Error struct{	
	functionName string
	packageName string
	error string
	message string
	extra string
}

