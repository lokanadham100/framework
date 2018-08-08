package logger

import (
	"os"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/voonik/framework/pkg/config"
)

func init(){
	fmt.Println("test")
	Register("logrus", NewLogrus)
}

func NewLogrus() (loggerInterface ,error){
	logg := logrus.New()
	logg.Level = getLogLevel()	
	logg.Formatter = &logrus.TextFormatter{}
	logg.SetOutput(os.Stdout)
	return logg,nil	
}

func getLogLevel() logrus.Level{
	l, _ := logrus.ParseLevel(config.LogConfigLevel())	
	return l
}