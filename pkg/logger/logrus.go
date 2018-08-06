package logger

import (
	"github.com/sirupsen/logrus"
)

func init(){
	Register("logrus", NewLogrus)
}

func NewLogrus(config map[string]string){
	log := logrus.New()
	return log,nil	
}