package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	ListenAddress string
	Proxies       []Proxy
}

type Proxy struct {
	Address string
	Headers map[string]string
}

func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config *Config
	err = json.NewDecoder(file).Decode(&config)
	return config, err
}
