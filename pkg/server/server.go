package server

import (
	"net"
	"fmt"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"	
	

	"github.com/voonik/framework/pkg/logger"
	"github.com/voonik/framework/pkg/middleware"
	"github.com/voonik/framework/pkg/event"		
)

// type ServiceToHandlerMap map[func(s *grpc.Server,srv interface{})]interface{}

var serverEvent event.WrapInterface
var listener net.Listener

func Init(){	
	logger.Init()
	serverEvent, _ = event.GetWrapEvent("process", context.Background())
	createGrpcServer()
}

var GrpcServer *grpc.Server

func Start(){
	serverEvent.Start(context.Background())
	listener = createSocket()
	GrpcServer.Serve(listener)
}

func Finish(){
	listener.Close()
	serverEvent.Finish(context.Background())
}

func createSocket() net.Listener{	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 10000))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	return lis
}

func createGrpcServer() {
	grpclog.SetLogger(logger.GetLoggerWithName("grpc"))
	GrpcServer = grpc.NewServer(
		middleware.StreamServerInterceptor(),
		middleware.UnaryServerInterceptor(),		
    )	
}