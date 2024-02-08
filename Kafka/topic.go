package main

import (
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

func main() {
	// Kafka broker address
	brokers := []string{"localhost:9092"}

	// Configuration for the admin client
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0 // Change to your desired Kafka version

	// Create the Kafka admin client
	admin, err := sarama.NewClusterAdmin(brokers, config)
	if err != nil {
		fmt.Printf("Error creating admin client: %v", err)
		os.Exit(1)
	}
	defer func() {
		// Close the admin client when done
		if err := admin.Close(); err != nil {
			fmt.Printf("Error closing admin client: %v", err)
		}
	}()

	// Define the details of the topic to be created
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     2, // Number of partitions for the topic
		ReplicationFactor: 1, // Replication factor for the topic
	}
	topic := "Users" // Name of the topic to be created

	// Create the topic with the specified details
	err = admin.CreateTopic(topic, topicDetail, false)
	if err != nil {
		fmt.Printf("Error creating topic: %v", err)
		os.Exit(1)
	}
	fmt.Println("Topic Created Successfully!")
}
