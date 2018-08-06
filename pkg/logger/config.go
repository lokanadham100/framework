package logger

type config struct{
	logger string
	level string
	tracer string
	formatter string
	output string
}

var Config *config