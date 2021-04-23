package metrics

import (
	"fmt"
	log "movie-service/logger"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var httpResponseLatencyMetric = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_latency_seconds",
	},
	[]string{"query"})

var httpResponseStatusMetric = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_response_status_code",
	},
	[]string{"status_code", "query"})

func NewLoggingMiddleware(logger log.Logger) mux.MiddlewareFunc {
	return mux.MiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			path := r.URL.Path
			wrappedWriter := wrapResponseWriter(w)

			start := time.Now()
			next.ServeHTTP(wrappedWriter, r)
			latency := time.Since(start)

			statusCode := wrappedWriter.status

			httpResponseLatencyMetric.WithLabelValues(path).Observe(latency.Seconds())

			switch {
			case statusCode > 499:
				logger.Errorf(ctx, "%s %s %d (%dms)", r.Method, path, statusCode, latency.Milliseconds())
			case statusCode > 399:
				logger.Warnf(ctx, "%s %s %d (%dms)", r.Method, path, statusCode, latency.Milliseconds())
			default:
				logger.Infof(ctx, "%s %s %d (%dms)", r.Method, path, statusCode, latency.Milliseconds())
			}
			httpResponseStatusMetric.WithLabelValues(fmt.Sprintf("%dxx", statusCode/100), path).Inc()
		})
	})
}

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}
