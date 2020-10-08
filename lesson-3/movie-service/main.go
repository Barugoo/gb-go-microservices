package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/barugoo/gb-go-microservices/lesson-3/movie-service/api"
)

func main() {
	srv := grpc.NewServer()

	pb.RegisterMovieServer(srv, &Service{})

	listener, err := net.Listen("tcp", ":9094")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting server on localhost:9094")
	srv.Serve(listener)
}
