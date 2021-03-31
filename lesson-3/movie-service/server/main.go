package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	pb "movie-service/server/api"
)

func main() {
	db, err := sql.Open(
		"postgres",
		"postgres://localhost:5432/postgres?sslmode=disable&user=postgres&database=postgres&password=postgres",
	)
	if err != nil {
		log.Fatal(err)
	}
	for db.Ping() != nil {
		log.Println(db.Ping())
	}
	s := &Service{db: db}

	srv := grpc.NewServer()

	pb.RegisterMovieServer(srv, s)

	listener, err := net.Listen("tcp", ":9098")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting server on localhost:9094")
	srv.Serve(listener)
}
