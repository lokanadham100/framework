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
	serverEvent, _ = event.GetWrapEvent("process", context.Background())
	createGrpcServer()
}

var grpcServer *grpc.Server

type protoDef func(s *grpc.Server,srv interface{})

func RegisterHandlers(pdef protoDef, handler interface{}){	
	pdef(grpcServer, handler)
}

func Start(){
	serverEvent.Start(context.Background())
	listener = createSocket()
	grpcServer.Serve(listener)
}

func Finish(){
	listener.Close()
	serverEvent.Finish(context.Background())
}

func createSocket() net.Listener{	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	return lis
}

func createGrpcServer() {
	grpclog.SetLogger(logger.GetLoggerWithName("grpc"))
	grpcServer = grpc.NewServer(
		middleware.StreamServerInterceptor(),
		middleware.UnaryServerInterceptor(),		
    )	
}