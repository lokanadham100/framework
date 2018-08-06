package config
	
import (
	"fmt"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"github.com/micro/go-config/source/env"
)

type Host struct {
        Address string `json:"address"`
        Port int `json:"port"`
}

type Config struct{
	Hosts map[string]Host `json:"hosts"`
}

var conf Config

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
	fmt.Println(conf)
}

