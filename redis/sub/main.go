package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("connected")

	pubsub := client.Subscribe("testChannel")
	_, err = pubsub.Receive()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("subscribed")

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}