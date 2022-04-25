package main

import (
	"context"
	"os"
	"os/signal"

	"cloud.google.com/go/pubsub"
	"github.com/go-redis/redis/v8"
	"github.com/machester4/products-leaderboard/cmd/messaging"
	"github.com/machester4/products-leaderboard/cmd/web"
	"github.com/machester4/products-leaderboard/configs"
	"github.com/machester4/products-leaderboard/infra/data/repositories"
	"github.com/machester4/products-leaderboard/infra/queue"
	"github.com/machester4/products-leaderboard/infra/server"
	"github.com/machester4/products-leaderboard/internal/application/services"
)

func main() {
	config := configs.Config{
		Redis:  configs.NewRedisConfigFromEnv(),
		PubSub: configs.NewPubSubConfigFromEnv(),
		Server: configs.NewServerConfigFromEnv(),
	}

	// create redis client
	client := redis.NewClient(&redis.Options{
		Addr: config.Redis.Host,
	})

	// ping redis
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// create pubsub client
	psc, err := pubsub.NewClient(context.Background(), config.PubSub.ProjectID)
	if err != nil {
		panic(err)
	}
	defer psc.Close()

	// Create queue
	queue := queue.NewPubsubQueue(psc)

	// create repository
	repo := repositories.NewProductLeadBoardRedis(client)

	// create service
	service := services.NewProductLeadboard(repo)

	// Initialize messaging
	consumer := messaging.New(queue)

	// Consume messages
	go func(s services.ProductLeadboard) {
		consumer.ConsumeIncrementScore(service)
	}(service)

	// Create fiber server
	fiber := server.NewFiberServer()

	// Initialize web
	web.New(fiber).InitializeRoutes(service)

	// Start server
	go func(port string) {
		if err := fiber.Listen(port); err != nil {
			panic(err)
		}
	}(config.Server.Port)

	// wait for interrupt signal to gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
