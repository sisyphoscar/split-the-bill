package configs

import (
	"os"
)

type EndpointConfig struct {
	Broker string
}

var Endpoints EndpointConfig

// LoadBrokerConfig load broker endpoint
func LoadEndPointConfig() {
	Endpoints = EndpointConfig{
		Broker: os.Getenv("BROKER_ENDPOINT"),
	}
}
