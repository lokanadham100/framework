package middleware

import (
	"google.golang.org/grpc"

	"github.com/voonik/framework/pkg/logger"	

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
)

func StreamServerInterceptor() (grpc.ServerOption){
	return grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
    	    grpc_ctxtags.StreamServerInterceptor(),
        	grpc_opentracing.StreamServerInterceptor(),
        	grpc_prometheus.StreamServerInterceptor,
        	grpc_logrus.StreamServerInterceptor(logger.GetLoggerWithName("middleware")),
        	// grpc_auth.StreamServerInterceptor(myAuthFunction),
        	grpc_recovery.StreamServerInterceptor(),
 	))
}
    
func UnaryServerInterceptor() (grpc.ServerOption){    
    	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
        	grpc_ctxtags.UnaryServerInterceptor(),
        	grpc_opentracing.UnaryServerInterceptor(),
        	grpc_prometheus.UnaryServerInterceptor,
        	grpc_logrus.UnaryServerInterceptor(logger.GetLoggerWithName("middleware")),
        	// grpc_auth.UnaryServerInterceptor(myAuthFunction),
        	grpc_recovery.UnaryServerInterceptor(),
    	))
    }