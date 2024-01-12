package main

import (
	"context"
	"encoding/json"
	"fmt"
	j "github.com/cjmalloy/jasper-sdk/jasper-sdk-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"log"
)

var ctx = context.Background()

func logRedis() error {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	subscribeAndLogRef := func(channel string) {
		fmt.Println("Subscribing to channel:", channel)
		pubsub := rdb.PSubscribe(ctx, channel)
		defer pubsub.Close()
		_, err := pubsub.Receive(ctx)
		if err != nil {
			fmt.Printf("Error subscribing to channel %s: %v\n", channel, err)
			return
		}

		for msg := range pubsub.Channel() {
			var ref j.Ref
			if err := json.Unmarshal([]byte(msg.Payload), &ref); err != nil {
				log.Fatalf("Error unmarshalling JSON: %v", err)
			}
			fmt.Println("Channel:", msg.Channel, "Ref:", ref.Title)
		}
		fmt.Println("Exited the loop for channel:", channel)
	}

	subscribeAndLog := func(channel string) {
		fmt.Println("Subscribing to channel:", channel)
		pubsub := rdb.PSubscribe(ctx, channel)
		defer pubsub.Close()
		_, err := pubsub.Receive(ctx)
		if err != nil {
			fmt.Printf("Error subscribing to channel %s: %v\n", channel, err)
			return
		}

		for msg := range pubsub.Channel() {
			fmt.Println("Channel:", msg.Channel, "Message:", msg.Payload)
		}
		fmt.Println("Exited the loop for channel:", channel)
	}

	go subscribeAndLogRef("ref/*")
	go subscribeAndLog("tag/*")
	go subscribeAndLog("response/*")
	return nil
}

func logMq() error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:", r)
		}
	}()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err)
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel:", err)
		return err
	}

	q, err := ch.QueueDeclare(
		"hello-queue", // name
		false,         // durable
		false,         // delete when usused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Println("Failed to declare a queue:", err)
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Println("Failed to register a consumer:", err)
		return err
	}

	go func() {
		defer conn.Close()
		defer ch.Close()
		for d := range msgs {
			log.Println("Received a message:", d.Body)
		}
	}()

	return nil
}

func main() {
	redisErr := logRedis()
	if redisErr == nil {
		log.Println("Logging Redis")
	} else {
		log.Println("Ignoring Redis")
	}
	mqErr := logMq()
	if mqErr == nil {
		log.Println("Logging Redis")
	} else {
		log.Println("Ignoring MQ")
	}
	select {}
}
