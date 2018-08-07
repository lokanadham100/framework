package config
	
type logConfig struct{
	Logger string
	Level string
	Output string
}

type traceConfig struct{
	Host string
	Port string
	ServiceName string
	Type string
	Param float64
}

type metricConfig struct{

}

type databaseConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

