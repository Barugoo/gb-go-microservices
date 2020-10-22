package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"gopkg.in/alexcesaro/statsd.v2"

	pb "github.com/barugoo/gb-go-microservices/lesson-3/movie-service/api"
)

var Stat *statsd.Client

func main() {
	var err error
	Stat, err = statsd.New(
		statsd.Address("graphite:8125"),
		statsd.Prefix("movie"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer Stat.Close()

	go func() {
		for {
			Stat.Increment("cron")
			log.Println("incremented")
			time.Sleep(1 * time.Second)
		}
	}()

	srv := grpc.NewServer()

	pb.RegisterMovieServer(srv, &Service{})

	listener, err := net.Listen("tcp", ":9098")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting server on localhost:9098")
	srv.Serve(listener)
}
