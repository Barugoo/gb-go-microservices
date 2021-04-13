package main

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using tools/gowrap/prometheus_grpc_client template

//go:generate gowrap gen -p movie-service/api -i MovieClient -t tools/gowrap/prometheus_grpc_client -o movie_service_prometheus.go

import (
	"context"
	"movie-service/api"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
)

// MovieClientWithPrometheus implements api.MovieClient interface with all methods wrapped
// with Prometheus metrics
type MovieClientWithPrometheus struct {
	base         api.MovieClient
	instanceName string
}

var movieclientDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "movieclient_duration_seconds",
		Help:       "movieclient runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewMovieClientWithPrometheus returns an instance of the api.MovieClient decorated with prometheus summary metric
func NewMovieClientWithPrometheus(base api.MovieClient, instanceName string) MovieClientWithPrometheus {
	return MovieClientWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// GetMovie implements api.MovieClient
func (_d MovieClientWithPrometheus) GetMovie(ctx context.Context, in *api.GetMovieRequest, opts ...grpc.CallOption) (gp1 *api.GetMovieResponse, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		movieclientDurationSummaryVec.WithLabelValues(_d.instanceName, "GetMovie", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetMovie(ctx, in, opts...)
}
