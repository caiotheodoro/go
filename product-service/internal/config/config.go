package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort         uint16
	AuthGRPCAddress    string
	OrderServiceAddress string
}

func Load() Config {
	cfg := Config{
		ServerPort:         8082,
		AuthGRPCAddress:    "auth-grpc:8000",
		OrderServiceAddress: "http://order-redis:8081",
	}

	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			cfg.ServerPort = uint16(port)
		}
	}

	if authAddr, exists := os.LookupEnv("AUTH_GRPC_ADDRESS"); exists {
		cfg.AuthGRPCAddress = authAddr
	}

	if orderAddr, exists := os.LookupEnv("ORDER_SERVICE_ADDRESS"); exists {
		cfg.OrderServiceAddress = orderAddr
	}

	return cfg
}