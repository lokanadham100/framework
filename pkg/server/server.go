package server

import (
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/voonik/framework/pkg/logger"	
	"github.com/voonik/framework/pkg/config"
	"github.com/voonik/framework/pkg/event"		
)

// type ServiceToHandlerMap map[func(s *grpc.Server,srv interface{})]interface{}

var serverEvent = event.GetWrapEvent("process", context.Background())
var listener net.Listener

func Init(){
	checkAndSetEnv()
	config.LoadConfig()
	createGrpcServer()
}

func setEnv(env string){
	os.Setenv("ENV", env)
	os.Setenv("ENVIRONMENT", env)
}

var grpcServer *grpc.Server

type protoDef func(s *grpc.Server,srv interface{})

func RegisterHandlers(pdef protoDef, handler interface{}){	
	pdef(grpcServer, handler)
}

func Start(){
	listener = createSocket()
	grpcServer.Serve(listener)
}

func Finish(){
	listener.Close()
	serverEvent.Finish(context.Background())
}

func checkAndSetEnv(){
	if env := os.Getenv("ENV"); env == ""{
		if env := os.Getenv("ENVIRONMENT"); env == ""{
			setEnv("development")
		}else{
			setEnv(env)
		}
	}else{
		setEnv(env)
	}
}

func createSocket() net.Listener{	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	return lis
}

func createGrpcServer() {
	grpcServer = grpc.NewServer()
}