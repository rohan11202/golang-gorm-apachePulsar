package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Total number of requests received, labeled by HTTP method and route
	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests received",
		},
		[]string{"method", "route"},
	)

	// Total number of responses by HTTP status code, labeled by route and status
	responseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_response_status_total",
			Help: "Total number of HTTP responses by status code",
		},
		[]string{"route", "status"},
	)

	// Histogram for tracking request duration (latency) in seconds, labeled by route and status
	responseDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_duration_seconds",
			Help:    "Duration of HTTP responses in seconds",
			Buckets: prometheus.DefBuckets, // Default latency buckets
		},
		[]string{"route", "status"},
	)
)

func Init() {
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(responseStatus)
	prometheus.MustRegister(responseDuration)
}
