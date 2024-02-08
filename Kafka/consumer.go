package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
)

func main() {
	// Kafka broker address
	brokers := []string{"localhost:9092"}

	// Configuration for the consumer group
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0 // Change to your desired Kafka version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Create a new consumer group
	group, err := sarama.NewConsumerGroup(brokers, "test", config)
	if err != nil {
		fmt.Printf("Error creating consumer group: %v\n", err)
		os.Exit(1)
	}

	// Trap SIGINT and SIGTERM signals to gracefully shutdown the consumer
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Context for managing the consumer group
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start consuming messages
	go func() {
		for {
			// Consume messages from the topic "Users"
			err := group.Consume(ctx, []string{"Users"}, &consumerHandler{})
			if err != nil {
				fmt.Printf("Error consuming messages: %v\n", err)
				return
			}
		}
	}()

	fmt.Println("Consumer running. Press Ctrl+C to exit.")
	<-signals
	fmt.Println("Shutting down consumer...")
}

// consumerHandler implements the sarama.ConsumerGroupHandler interface
type consumerHandler struct{}

func (h *consumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	fmt.Println("Consumer ready")
	return nil
}

func (h *consumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	// Mark the consumer as completed
	fmt.Println("Consumer completed")
	return nil
}

func (h *consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// Process messages claimed by the consumer
	for msg := range claim.Messages() {
		fmt.Printf("Received message %s on partition %d\n", msg.Value, msg.Partition)
		session.MarkMessage(msg, "")
	}
	return nil
}
