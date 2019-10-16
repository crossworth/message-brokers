package main

import (
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://127.0.0.1:1883")
	opts.SetClientID("test-client-1")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// qos
	// 0 - at most once
	// 1 - at least once
	// 2- exactly once

	// retained - if the broker should store the last message with qos
	// and send to new subscribers on connect

	for i := 0; i < 100; i++ {
		fmt.Println("---- Publishing ----")
		token := client.Publish("my/fake/sensor", 2, false, fmt.Sprintf("My-Fake-Data %d", i))
		token.Wait()
	}

	// 250 is a wait time
	client.Disconnect(250)
}