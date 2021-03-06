package config

import (
	"fmt"
	"os"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source"
	"github.com/micro/go-config/source/file"
	"github.com/micro/go-config/source/env"
)

var conf *Config

func init(){
	loadConfig()
}

type Config struct{
	LogConfig logConfig
	TraceConfig traceConfig
	MetricConfig metricConfig
	DatabaseConfig databaseConfig	
}

func loadConfig(){
	readConfigFromFile()
	writeConfigToStruct()	
	fmt.Println(conf)
	// fmt.Println(config.Map())
}

func readConfigFromFile(){
	cnfPath := make([]source.Source,0)
	cnfPath = addFileToConfig(cnfPath,"config/config.toml")	
	cnfPath = addFileToConfig(cnfPath,fmt.Sprintf("config/config_%s.toml",os.Getenv("ENV")))
	cnfPath = addFileToConfig(cnfPath,"config/database.toml")
	cnfPath = addFileToConfig(cnfPath,"config/redis.toml")
	cnfPath = addFileToConfig(cnfPath,"config/kafka.toml")
	cnfPath = addFileToConfig(cnfPath,"config/tracing.toml")
	cnfPath = addFileToConfig(cnfPath,"config/metrics.toml")
	cnfPath = append(cnfPath,env.NewSource())
	config.Load(
		cnfPath...,		
	)
}

func writeConfigToStruct(){
	config.Scan(&conf)	
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