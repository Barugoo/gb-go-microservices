package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	pb "payment-service/api"
	log "payment-service/logger"
)

var logger = log.NewLogger()

func main() {
	rand.Seed(time.Now().Unix())

	ctx := context.Background()

	f, err := os.Create(
		fmt.Sprintf("/var/log/super-cinema/payment-service.log"),
	)
	if err != nil {
		logger.Fatalf(ctx, "error opening file: %v", err)
	}
	defer f.Close()
	logger.SetOutput(f)

	consulAddr := flag.String("consul_addr", "consul:8500", "Consul address")
	flag.Parse()

	if err := loadConfig(*consulAddr); err != nil {
		logger.Fatal(ctx, err)
	}

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

	srv := grpc.NewServer()

	pb.RegisterPaymentServiceServer(srv, &Service{})

	listener, err := net.Listen("tcp", "0.0.0.0:"+cfg.Port)
	if err != nil {
		logger.Fatalf(ctx, "failed to listen: %v", err)
	}

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		logger.Infof(ctx, "Starting metrics http server on port 9099")
		http.ListenAndServe(":9099", nil)
	}()

	logger.Infof(ctx, "Starting grpc server on port "+cfg.Port)
	srv.Serve(listener)
}
