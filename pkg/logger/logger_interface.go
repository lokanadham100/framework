package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type loggerInterface interface{	
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})

	//TODO : Need to change this one. Dont know how to use this. So placing here
	WithField(string,interface{}) *logrus.Entry 	
}

var registry = make(map[string]logInitFunc)

type logInitFunc func() (loggerInterface, error)

func Register(name string, lIFunc logInitFunc){
	if _, ok := registry[name]; ok {
		panic(fmt.Sprintf("%s is already registered", name))
	}
	registry[name] = lIFunc
}

func Get(name string)(loggerInterface, error){
	f, ok := registry[name]
	if !ok {
		return nil, fmt.Errorf("logger %q not found", name)
	}
	return f()
}

//TODO : Need to remove
func GetLoggerWithName(name string) (*logrus.Entry){
	return log.WithField("pkg", name)
}