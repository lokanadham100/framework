package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/voonik/framework/pkg/config"
)

func init(){
	Register("logrus", NewLogrus)
}

func NewLogrus() (loggerInterface ,error){
	log := logrus.New()
	log.level = config.LogConfigLevel()
	return log,nil	
}