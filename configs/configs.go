package configs

type Config struct {
	Redis  RedisConfig
	PubSub PubSubConfig
	Server ServerConfig
}
