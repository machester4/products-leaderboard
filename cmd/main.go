package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// config := configs.Config{
	// 	Redis:  configs.NewRedisConfigFromEnv(),
	// 	PubSub: configs.NewPubSubConfigFromEnv(),
	// 	Server: configs.NewServerConfigFromEnv(),
	// }

	// // create redis client
	// client := redis.NewClient(&redis.Options{
	// 	Addr: config.Redis.Host,
	// })

	// // ping redis
	// _, err := client.Ping(context.Background()).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// defer client.Close()

	// // create pubsub client
	// psc, err := pubsub.NewClient(context.Background(), config.PubSub.ProjectID)
	// if err != nil {
	// 	panic(err)
	// }
	// defer psc.Close()

	// // Create queue
	// queue := queue.NewPubsubQueue(psc)

	// // create repository
	// repo := repositories.NewProductLeadBoardRedis(client)

	// // create service
	// service := services.NewProductLeadboard(repo)

	// // Initialize messaging
	// messaging.New(queue).ConsumeIncrementScore(service)

	// Create fiber server
	//fiber := server.NewFiberServer()

	// Initialize web
	//web.New(fiber).InitializeRoutes(service)

	// Start server
	// if err := fiber.Listen(config.Server.Port); err != nil {
	// 	panic(err)
	// }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	fmt.Println("Server starting...")
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
