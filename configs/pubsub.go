package configs

import "os"

type PubSubConfig struct {
	ProjectID string
}

func NewPubSubConfigFromEnv() PubSubConfig {
	projectID := os.Getenv("PUBSUB_PROJECT")
	if projectID == "" {
		panic("PUBSUB_PROJECT is not set")
	}

	return PubSubConfig{
		ProjectID: projectID,
	}
}
