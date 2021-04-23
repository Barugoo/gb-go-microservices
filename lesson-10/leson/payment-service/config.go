package main

import (
	"encoding/json"
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
)

const ServicePrefix = "service/payment-service"

type Config struct {
	Port string `json:"port"`
}

var cfg Config

func loadConfig(addr string) error {
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

	return json.Unmarshal(cfgRaw.Value, &cfg)
}
