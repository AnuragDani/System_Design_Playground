package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

func main() {
	// Kafka broker address
	brokers := []string{"localhost:9092"}

	// Get message from command line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a message as command line argument.")
		os.Exit(1)
	}
	msg := os.Args[1]

	// Configuration for the producer
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0 // Change to your desired Kafka version
	config.Producer.Return.Successes = true

	// Create the Kafka producer
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Printf("Error creating producer: %v", err)
		os.Exit(1)
	}
	defer func() {
		// Close the producer when done
		if err := producer.Close(); err != nil {
			fmt.Printf("Error closing producer: %v", err)
		}
	}()

	// Determine partition based on message
	partition := int32(0)
	if strings.ToUpper(msg)[0] >= 'N' {
		partition = int32(1)
	}

	// Prepare the message to be sent
	message := &sarama.ProducerMessage{
		Topic:     "Users",
		Partition: partition,
		Value:     sarama.StringEncoder(msg),
	}

	// Send the message
	_, _, err = producer.SendMessage(message)
	if err != nil {
		fmt.Printf("Error sending message: %v", err)
		os.Exit(1)
	}
	fmt.Println("Producer: Message Sent Successfully!")
	fmt.Printf("Topic: %s\n", message.Topic)
	fmt.Printf("Partition: %d\n", partition)
}
