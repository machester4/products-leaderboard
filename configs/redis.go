package configs

import "os"

type RedisConfig struct {
	Host string
}

func NewRedisConfigFromEnv() RedisConfig {
	host := "localhost:6379"
	if hostEnv := os.Getenv("REDISHOST"); hostEnv != "" {
		host = hostEnv
	}

	return RedisConfig{
		Host: host,
	}
}
