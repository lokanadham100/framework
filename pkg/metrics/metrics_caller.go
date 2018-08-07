package metrics

func DatabaseEventCounter(qtype, query string){
	defaultServerMetrics.databaseEventCounter.WithLabelValues(qtype,query).Inc()
}

func FunctionEventCounter(packageName, functionName string){
	defaultServerMetrics.functionEventCounter.WithLabelValues(packageName, functionName).Inc()
}

func ErrorEventCounter(packageName, functionName, err, message string){
	defaultServerMetrics.errorEventCounter.WithLabelValues(packageName, functionName, err, message).Inc()
}

func DatabaseEventHistogram(qtype string, query string, timer float64){
	defaultServerMetrics.databaseEventHistogram.WithLabelValues(qtype, query).Observe(timer)
}

func FunctionEventHistogram(packageName string, functionName string, timer float64) {
	defaultServerMetrics.functionEventHistogram.WithLabelValues(packageName, functionName).Observe(timer)
}