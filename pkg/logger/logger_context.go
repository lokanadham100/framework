package logger

type LoggerContext interface{
	LoggerFromContext()
	ContextWithLogger()
}

func getTraceID(ctx context.Context) string{
	GetLogContext(ctx,"jaeger")
}

type LogInitFunc func(context.Context) (string, error)

var logContextRegistry = make(map[string]LogTraceIdFunc)

func RegisterLogContext(name string, logTraceIdFunc LogTraceIdFunc){
	if _, ok := logContextRegistry[name]; ok {
		panic(fmt.Sprintf("%s is already registered", name))
	}
	logContextRegistry[name] = logInitFunc
}

func GetLogContext(ctx context.Context, name string)(string, error){
	f, ok := logContextRegistry[name]
	if !ok {
		return nil, fmt.Errorf("logger %q not found", name)
	}
	return f(ctx)
}
