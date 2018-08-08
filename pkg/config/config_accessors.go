package config

//logger
func LogConfigLogger() string{
	return conf.LogConfig.Logger
}

func LogConfigLevel() string{
	return conf.LogConfig.Level	
}

func LogConfigOutput() string{
	return conf.LogConfig.Output
}


//tracer
func TraceConfigHost() string{
	return conf.TraceConfig.Host
}

func TraceConfigPort() string{
	return conf.TraceConfig.Port
}

func TraceConfigServiceName() string{
	return conf.TraceConfig.ServiceName
}

func TraceConfigType() string{
	return conf.TraceConfig.Type
}

func TraceConfigParam() float64{
	return conf.TraceConfig.Param
}


//database
func DatabaseConfig() *databaseConfig{
	return &conf.DatabaseConfig
}


//metrics
func MetricConfigPort() string{
	return conf.MetricConfig.Port
}

func MetricConfigServiceName() string{
	return conf.MetricConfig.ServiceName
}