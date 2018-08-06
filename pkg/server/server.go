package server

import (
	"net"
	"logger"
	"google.golang.org/grpc"
)

type ServiceToHandlerMap map[func(s *grpc.Server,srv interface{})]interface{}

func Init(){
	config.LoadConfig()
}

func RegisterHandlers(srvmap ServiceToHandlerMap){
	listener := createSocket()
	srvr := createGrpcServer()
	for service , handler := range ServiceToHandlerMap {
		service(srvr, handler)
	}
	startServer(srvr,listener)
}
	
func createSocket() net.Listener{	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	return lis
}

func createGrpcServer() *grpc.Server{
	return grpc.NewServer()
}

func startServer(srvr *grpc.Server, listener net.Listener){
	srvr.Serve(listener)
}