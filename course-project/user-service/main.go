package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	pb "user-service/api"

	"github.com/go-xorm/xorm"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	dbHost, ok := os.LookupEnv("DATABASE_HOST")
	if !ok {
		log.Fatal("DATABASE_HOST env is not set")
	}
	dbName, ok := os.LookupEnv("DATABASE_NAME")
	if !ok {
		log.Fatal("DATABASE_NAME env is not set")
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("postgres://%s/%s?sslmode=disable&user=postgres&database=%s&password=postgres", dbHost, dbName, dbName),
	)
	if err != nil {
		log.Fatal(err)
	}
	for db.Ping() != nil {
		log.Println(db.Ping())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgtres", driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		log.Println(err)
	}

	engine, err := xorm.NewEngine("postgres", fmt.Sprintf("postgres://%s/%s?sslmode=disable&user=postgres&database=%s&password=postgres", dbHost, dbName, dbName))

	addr := os.Getenv("GRPC_PORT")
	if addr == "" {
		log.Fatal("unable to retreive env GRPC_PORT")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", addr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, NewService(engine))
	log.Println("Serving GRPC on " + addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
