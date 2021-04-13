package main

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "movie-service/logger"
	metrics "movie-service/metrics"
	reqdata "movie-service/reqdata"
)

var logger = log.NewLogger()

func main() {
	rand.Seed(time.Now().Unix())

	ctx := context.Background()

	r := mux.NewRouter()
	r.Use(reqdata.RequestIDMiddleware, metrics.NewLoggingMiddleware(logger))

	r.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		// pretend some work
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")
	r.HandleFunc("/failure", func(w http.ResponseWriter, r *http.Request) {
		// pretend some work
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		w.WriteHeader(http.StatusInternalServerError)
	}).Methods("GET")

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		logger.Infof(ctx, "Starting http server on port 9099")
		http.ListenAndServe(":9099", nil)
	}()

	logger.Fatal(ctx, http.ListenAndServe(":9098", r))
}
