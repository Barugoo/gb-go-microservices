package main

import (
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
)

const ServicePrefix = "service/gateway-user"

type Config struct {
	Port         string
	UserGRPCAddr string
	UserAddr     string
	MovieAddr    string
}

var cfg Config

func loadConfig(addr string) error {
	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = addr
	consul, err := consulapi.NewClient(consulConfig)
	if err != nil {
		return err
	}

	port, _, err := consul.KV().Get(ServicePrefix+"/port", nil)
	if err != nil || port == nil {
		return fmt.Errorf("Can't get port value")
	}
	userGRPCAddr, _, err := consul.KV().Get(ServicePrefix+"/user_grpc_addr", nil)
	if err != nil || port == nil {
		return fmt.Errorf("Can't get port value")
	}
	userAddr, _, err := consul.KV().Get(ServicePrefix+"/user_addr", nil)
	if err != nil || port == nil {
		return fmt.Errorf("Can't get port value")
	}
	movieAddr, _, err := consul.KV().Get(ServicePrefix+"/movie_addr", nil)
	if err != nil || port == nil {
		return fmt.Errorf("Can't get port value")
	}

	cfg.Port = string(port.Value)
	cfg.UserAddr = string(userAddr.Value)
	cfg.UserGRPCAddr = string(userGRPCAddr.Value)
	cfg.MovieAddr = string(movieAddr.Value)

	return nil
}
