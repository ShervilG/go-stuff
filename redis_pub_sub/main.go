package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	subscriber := NewSubscriber("localhost:6379")
	producer := NewProducer("localhost:6379")

	wg := &sync.WaitGroup{}
	wg.Add(1)

	fmt.Println("Pub-Sub Runner has started")

	// Example usage
	go func() {
		for msg := range subscriber.PSubscribe("example-channel*").Channel() {
			fmt.Printf("Received message: %s on channel: %s\n", msg.Payload, msg.Channel)
		}

		wg.Done()
	}()

	for i := 0; i <= 5; i += 1 {
		time.Sleep(1 * time.Second)
		err := producer.Publish("example-channel", fmt.Sprintf("Hello, Redis! %d", i))
		if err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}

		err = producer.Publish("example-channel-2", fmt.Sprintf("Hello, Redis 2! %d", i))
		if err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}
	}

	wg.Wait()
}
