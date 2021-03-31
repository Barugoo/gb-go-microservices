package main

import (
	"context"
	"fmt"
	"log"
	pb "movie-service/client/api"

	"google.golang.org/grpc"
)

func NewMovieClient(addr string) (pb.MovieClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return pb.NewMovieClient(conn), nil
}

func main() {
	movieClient, err := NewMovieClient("localhost:9098")
	if err != nil {
		log.Fatal(err)
	}

	req := &pb.GetMovieRequest{
		MovieId: 1,
	}
	resp, err := movieClient.GetMovie(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*resp)
}
