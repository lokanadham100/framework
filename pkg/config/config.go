package config

import (
	"fmt"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"github.com/micro/go-config/source/env"
)

var Conf *Config

type Config struct{
	LogConfig logConfig 
	TraceConfig traceConfig
	MetricConfig metricConfig	
}

func LoadConfig(){
	config.Load(
		file.NewSource(
			file.WithPath("config/config.toml"),
		),		
		env.NewSource(),
	)
	config.Scan(&Conf)
	fmt.Println(Conf)	
}