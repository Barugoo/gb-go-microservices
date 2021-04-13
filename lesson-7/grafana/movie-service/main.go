package main

import (
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	pb "github.com/barugoo/gb-go-microservices/lesson-3/movie-service/api"
)

func NewMovieService(cc grpc.Conn) MovieClient {
	return pb.NewMovieClientWithPrometheus(cc)
}

func main() {
	srv := grpc.NewServer()

	pb.RegisterMovieServer(srv, &Service{})

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Println("Starting http server on port 9099")
		http.ListenAndServe(":9099", nil)
	}()

	listener, err := net.Listen("tcp", ":9098")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting server on localhost:9098")
	srv.Serve(listener)
}
