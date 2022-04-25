package configs

import "os"

type ServerConfig struct {
	Port string
}

func NewServerConfigFromEnv() ServerConfig {
	port := ":8080"
	if portEnv := os.Getenv("PORT"); portEnv != "" {
		port = portEnv
	}

	return ServerConfig{
		Port: port,
	}
}
