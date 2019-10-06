package main

import (
	"fmt"
	"time"

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

	channel := "testChannel"

	i := 0
	for {
		message := i
		fmt.Printf("Publish message %q on %q\n", message, channel)

		err = client.Publish(channel, message).Err()
		if err != nil {
			fmt.Println(err)
			continue
		}

		i++
		time.Sleep(time.Second * 3)
	}
}
