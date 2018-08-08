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
	log.Level = getLogLevel()
	return log,nil	
}

func getLogLevel() logrus.Level{
	l, _ := logrus.ParseLevel(config.LogConfigLogger())
	return l
}