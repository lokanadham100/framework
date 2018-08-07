package config

import (
	"fmt"
	"os"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source"
	"github.com/micro/go-config/source/file"
	"github.com/micro/go-config/source/env"
)

var Conf *Config

type Config struct{
	LogConfig logConfig 
	TraceConfig traceConfig
	MetricConfig metricConfig	
	DatabaseConfig databaseConfig
}

func LoadConfig(){
	readConfigFromFile()
	writeConfigToStruct()	
	fmt.Println(Conf)	
}

func readConfigFromFile(){
	cnfPath := make([]source.Source,0)
	cnfPath = addFileToConfig(cnfPath,"config/config.toml")	
	cnfPath = addFileToConfig(cnfPath,fmt.Sprintf("config/config_%s.toml",os.Getenv("ENV")))
	cnfPath = addFileToConfig(cnfPath,"config/database.toml")
	cnfPath = addFileToConfig(cnfPath,"config/redis.toml")
	cnfPath = addFileToConfig(cnfPath,"config/kafka.toml")
	cnfPath = addFileToConfig(cnfPath,"config/tracing.toml")
	cnfPath = append(cnfPath,env.NewSource())
	config.Load(
		cnfPath...,		
	)
}

func writeConfigToStruct(){
	config.Scan(&Conf)	
}

func addFileToConfig(cnfPath []source.Source, filePath string) ([]source.Source){
	if _, err := os.Stat(filePath); !os.IsNotExist(err){
   		cnfPath = append(
   			cnfPath,
   			file.NewSource(
   				file.WithPath(filePath),
   			),
   		)
	}
	return cnfPath	
}