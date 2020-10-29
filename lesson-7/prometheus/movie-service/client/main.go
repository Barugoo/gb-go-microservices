package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "client/api"
)

func NewMovieClient(addr string) (pb.MovieClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return pb.NewMovieClient(conn), nil
}

func main() {
	log.Println(1)
	movieClient, err := NewMovieClient("localhost:9098")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(2)

	req := &pb.GetMovieRequest{
		MovieId: 1,
	}
	for {
		_, err = movieClient.GetMovie(context.Background(), req)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}