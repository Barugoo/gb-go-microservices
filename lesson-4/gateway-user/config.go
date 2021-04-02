package main

type Config struct {
	Addr         string
	UserGRPCAddr string
	UserAddr     string
	MovieAddr    string
}

var cfg = Config{
	Addr:         ":8182",
	UserGRPCAddr: ":8010",
	UserAddr:     "http://localhost:8081",
	MovieAddr:    "http://localhost:8080",
}
