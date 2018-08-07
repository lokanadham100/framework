package logger

import (
	"context"
)


func WithCtx(ctx context.Context) loggerInterface{
	return &Logger{ctx:ctx, loggerInterface:getLoggerWithContext(ctx)}
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Print(args ...interface{}) {
	log.Print(args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warning(args ...interface{}) {
	log.Warning(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Debugln(args ...interface{}) {
	log.Debugln(args...)
}

func Println(args ...interface{}) {
	log.Println(args...)
}

func Infoln(args ...interface{}) {
	log.Infoln(args...)
}

func Warnln(args ...interface{}) {
	log.Warnln(args...)
}

func Warningln(args ...interface{}) {
	log.Warningln(args...)
}

func Errorln(args ...interface{}) {
	log.Errorln(args...)
}

func Panicln(args ...interface{}) {
	log.Panicln(args...)
}

func Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}