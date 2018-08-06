package logger

import (
	"github.com/sirupsen/logrus"
)

func init(){
	Register("logrus", NewLogrus)
}

func NewLogrus() (loggerInterface ,error){
	log := logrus.New()	
	return log,nil	
}