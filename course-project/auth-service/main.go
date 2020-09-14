package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/barugoo/gb-go-microservices/course-project/auth-service/api"
)

var secret = []byte("secret")

func main() {
	addr := os.Getenv("GRPC_PORT")
	if addr == "" {
		log.Fatal("unable to retreive env GRPC_PORT")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", addr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, NewAuthService())
	log.Println("Serving GRPC on " + addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
