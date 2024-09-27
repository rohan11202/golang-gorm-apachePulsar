package middleware

import (
	"log"
	"net/http"
	"time"
	// "time"
)

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Println("Metrics middleware working")

		recorder := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(recorder, r)

		totalRequests.WithLabelValues(r.Method, r.URL.Path).Inc()

		responseStatus.WithLabelValues(r.URL.Path, http.StatusText(recorder.statusCode)).Inc()

		responseDuration.WithLabelValues(r.URL.Path, http.StatusText(recorder.statusCode)).Observe(time.Since(start).Seconds())
	})
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}
