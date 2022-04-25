package configs

type ServerConfig struct {
	Port string
}

func NewServerConfigFromEnv() ServerConfig {
	port := ":9090"
	// if portEnv := os.Getenv("PORT"); portEnv != "" {
	// 	port = portEnv
	// }

	return ServerConfig{
		Port: port,
	}
}
