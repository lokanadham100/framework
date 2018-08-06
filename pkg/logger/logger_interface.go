package logger

type LoggerInterface interface{
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
}

type FieldInterface interface{
	WithFields(map[string]interface{}) LoggerInterface
	WithField(string, interface{}) LoggerInterface
}

var registry = make(map[string]LogInitFunc)

type LogInitFunc func(map[string]string) (LoggerInterface, error)

func Register(name string, logInitFunc LogInitFunc){
	if _, ok := registry[name]; ok {
		panic(fmt.Sprintf("%s is already registered", name))
	}
	registry[name] = logInitFunc
}

func Get(name string, config map[string]string)(LoggerInterface, error){
	f, ok := registry[name]
	if !ok {
		return nil, fmt.Errorf("logger %q not found", name)
	}
	return f(config)
}

var loggerInterfaceInstance LoggerInterface

// TODO: Make a init call from server init time and set the logger from config as this has to check every time.
func getOrCreate() LoggerInterface{ 
	if loggerInterfaceInstance == nil{		
		loggerInterfaceInstance = createLoggerInterfaceInstance()
	}
	return loggerInterfaceInstance
}

func createLoggerInterfaceInstance LoggerInterface{
	Get("logrus")
}