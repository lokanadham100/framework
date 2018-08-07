package config

//logger
func LogConfigLogger() string{
	return Conf.LogConfig.Logger
}

func LogConfigLevel() string{
	return Conf.LogConfig.Level	
}

func LogConfigOutput() string{
	return Conf.LogConfig.Output
}


//tracer
func TraceConfigHost() string{
	return Conf.TraceConfig.Host
}

func TraceConfigPort() string{
	return Conf.TraceConfig.Port
}

func TraceConfigServiceName() string{
	return Conf.TraceConfig.ServiceName
}

func TraceConfigType() string{
	return Conf.TraceConfig.Type
}

func TraceConfigParam() float64{
	return Conf.TraceConfig.Param
}


//database
func DatabaseConfig() *databaseConfig{
	return &Conf.DatabaseConfig
}


//metrics
func MetricConfigPort() string{
	return Conf.MetricConfig.Port
}

func MetricConfigServiceName() string{
	return Conf.MetricConfig.ServiceName
}