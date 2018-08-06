package config

type Config struct{
	logger string
	logLevel string

	tracer string
}

func Logger() string{
	return conf.logger
}

func LogLevel() string{
	return conf.logLevel
}

func Tracer() string{
	return conf.tracer
}