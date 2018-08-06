package config
	
import (
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"github.com/micro/go-config/source/env"	
)

var conf config

func LoadConfig(){
	config.Load(
		file.NewSource(
			file.WithPath("/config/config.json"),
		),
		file.NewSource(
			file.WithPath("/config/config.json"),
		),
		env.NewSource(),
	)
	config.Scan(&conf)
}

