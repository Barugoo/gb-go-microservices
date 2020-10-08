package main

type Config struct {
	Addr         string
	UserGRPCAddr string
	UserAddr     string
	MovieAddr    string
}

var cfg = Config{
	Addr:         ":8082",
	UserGRPCAddr: ":9094",
	UserAddr:     "http://localhost:8081",
	MovieAddr:    "http://localhost:8080",
}
