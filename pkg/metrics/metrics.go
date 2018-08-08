package metrics

import (
	"fmt"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/voonik/framework/pkg/config"
//	"github.com/voonik/framework/pkg/logger"
)

func init() {	
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(fmt.Sprintf(":%s",config.MetricConfigPort()), nil)
	registerWithRegistry()	
}

type serverMetrics struct {
	databaseEventCounter          *prometheus.CounterVec
	functionEventCounter          *prometheus.CounterVec
	errorEventCounter             *prometheus.CounterVec	
	databaseEventHistogram        *prometheus.HistogramVec
	functionEventHistogram        *prometheus.HistogramVec
}

var defaultServerMetrics = NewServerMetrics()

func NewServerMetrics() *serverMetrics{
	return &serverMetrics{
		databaseEventCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: formatMetricName("databaseEventCounter"),
				Help: "Total number of database queries executed",
			}, []string{"type","query"}),
		functionEventCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: formatMetricName("functionEventCounter"),
				Help: "Total number of functions called",
			}, []string{"package_name","function_name"}),
		errorEventCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: formatMetricName("errorEventCounter"),
				Help: "Total number of errors raised",
			}, []string{"package_name","function_name","error","message"}),
		databaseEventHistogram: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: formatMetricName("databaseEventHistogram"),
				Help:    "Histogram of database query (seconds)",
				Buckets: prometheus.DefBuckets,
			}, []string{"type", "query"}),
		functionEventHistogram: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: formatMetricName("functionEventHistogram"),
				Help:    "Histogram of function execution (seconds)",
				Buckets: prometheus.DefBuckets,
			}, []string{"package_name","function_name"}),
	}
}

//TODO : Need to register all grpc apis also, to get 0 count metrics also.
func registerWithRegistry(){
	prometheus.MustRegister(defaultServerMetrics.databaseEventCounter)
	prometheus.MustRegister(defaultServerMetrics.functionEventCounter)
	prometheus.MustRegister(defaultServerMetrics.errorEventCounter)
	prometheus.MustRegister(defaultServerMetrics.databaseEventHistogram)
	prometheus.MustRegister(defaultServerMetrics.functionEventHistogram)
}

func formatMetricName(s string) string{
	if len(config.MetricConfigServiceName()) > 0 {
		return fmt.Sprintf("%s-%s",config.MetricConfigServiceName(),s)
	}else{
		return s
	}
}