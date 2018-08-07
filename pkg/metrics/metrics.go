package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/voonik/framework/pkg/config"
)

func init() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(append(":",config.MetricConfigPort()), nil))
}

type serverMetrics struct {
	databaseEventCounter          *prom.CounterVec
	functionEventCounter          *prom.CounterVec
	errorEventCounter             *prom.CounterVec	
	databaseEventHistogram        *prom.HistogramVec
	functionEventHistogram        *prom.HistogramVec
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
		databaseEventHistogram: prom.NewHistogramVec(
			prom.HistogramOpts{
				Name: formatMetricName("databaseEventHistogram"),
				Help:    "Histogram of database query (seconds)",
				Buckets: prom.DefBuckets,
			}[]string{"type", "query"}),
		functionEventHistogram: prom.NewHistogramVec(
			prom.HistogramOpts{
				Name: formatMetricName("functionEventHistogram"),
				Help:    "Histogram of function execution (seconds)",
				Buckets: prom.DefBuckets,
			}[]string{"package_name","function_name"}),
	}
}

func formatMetricName(s string) string{
	return fmt.Sprintf("%s-%s",config.MetricConfigServiceName(),s)
}