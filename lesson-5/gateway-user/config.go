package main

import (
	"encoding/json"
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/mitchellh/consulstructure"
)

const ServicePrefix = "service/gateway-user"

type Config struct {
	Port         string `consul:"port"`
	UserGRPCAddr string `consul:"user_grpc_addr"`
	UserAddr     string `consul:"user_addr"`
	MovieAddr    string `consul:"movie_addr"`
}

var cfg Config

func loadConfig(addr string) error {
	go func() {
		updateCh := make(chan interface{})
		errCh := make(chan error)
		decoder := &consulstructure.Decoder{
			Target:   &Config{},
			Prefix:   "service/gateway-user/config",
			UpdateCh: updateCh,
			ErrCh:    errCh,
		}
		go decoder.Run()
		for {
			select {
			case v := <-updateCh:
				fmt.Printf("updated cfg %v", v.(*Config))
			case err := <-errCh:
				fmt.Printf("Error: %s\n", err)
			}
		}
	}()

	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = addr
	consul, err := consulapi.NewClient(consulConfig)
	if err != nil {
		return err
	}

	cfgRaw, _, err := consul.KV().Get(ServicePrefix+"/config", nil)
	if err != nil || cfgRaw == nil {
		return fmt.Errorf("Can't get cfg value: %w", err)
	}

	return json.Unmarshal([]byte(cfgRaw.Value), &cfg)
}
