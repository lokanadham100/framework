package config

func LogConfigLogger() string{
	return Conf.LogConfig.Logger
}

func LogConfigLevel() string{
	return Conf.LogConfig.Level	
}

func LogConfigOutput() string{
	return Conf.LogConfig.Output
}
