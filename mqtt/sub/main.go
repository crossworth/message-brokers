package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func main() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://127.0.0.1:1883")
	opts.SetClientID("test-client-2")

	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Got message from topic %s, payload %q\n", msg.Topic(), msg.Payload())
	})

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	if token := client.Subscribe("my/fake/sensor", 2, nil); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	fmt.Println("waiting for messages")
	wait := make(chan struct{})

	<- wait
}
