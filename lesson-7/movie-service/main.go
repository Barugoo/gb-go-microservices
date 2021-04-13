package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"runtime"

	"google.golang.org/grpc"

	pb "github.com/barugoo/gb-go-microservices/lesson-3/movie-service/api"
)

func main() {
	srv := grpc.NewServer()

	pb.RegisterMovieServer(srv, &Service{})

	http.HandleFunc("/health", healthHandler)
	go func() {
		log.Printf("Starting http sever on port %d", 9097)
		http.ListenAndServe(":9097", nil)
	}()

	listener, err := net.Listen("tcp", "0.0.0.0:9098")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting server on localhost:9098")
	srv.Serve(listener)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)

	data, _ := json.Marshal(map[string]interface{}{
		"status":        "ok",
		"num_goroutine": runtime.NumGoroutine(),
		"mem":           m.HeapAlloc,
	})
	w.Write(data)
}
