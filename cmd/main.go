package main

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/machester4/products-leaderboard/cmd/web"
	"github.com/machester4/products-leaderboard/configs"
	"github.com/machester4/products-leaderboard/infra/data/repositories"
	"github.com/machester4/products-leaderboard/infra/server"
	"github.com/machester4/products-leaderboard/internal/application/services"
)

func main() {
	config := configs.Config{
		Redis: configs.NewRedisConfigFromEnv(),
		//PubSub: configs.NewPubSubConfigFromEnv(),
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
	// psc, err := pubsub.NewClient(context.Background(), config.PubSub.ProjectID)
	// if err != nil {
	// 	panic(err)
	// }
	// defer psc.Close()

	// Create queue
	//queue := queue.NewPubsubQueue(psc)

	// create repository
	repo := repositories.NewProductLeadBoardRedis(client)

	// create service
	service := services.NewProductLeadboard(repo)

	// Initialize messaging
	//messaging.New(queue).ConsumeIncrementScore(service)

	// Create fiber server
	fiber := server.NewFiberServer()

	// Initialize web
	web.New(fiber).InitializeRoutes(service)

	// Start server
	if err := fiber.Listen(config.Server.Port); err != nil {
		panic(err)
	}
}
