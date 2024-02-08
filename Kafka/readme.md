# Kafka Go Application

This repository contains a Kafka-based Go application consisting of a producer (`producer.go`) and a consumer (`consumer.go`). The producer sends messages to a Kafka topic, and the consumer consumes messages from the same topic.

## Requirements

To run this application, you need to have Go and Docker installed on your system.

## Setup

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/yourusername/kafka-go-application.git
    ```

2. Navigate to the project directory:

    ```bash
    cd kafka-go-application
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Start Kafka and Zookeeper services using Docker Compose:

    ```bash
    docker-compose up -d
    ```

## Running the Application

### Producer

To run the producer, execute the following command:

```bash
go run producer.go "your_message_here"
```

### Consumer
To run the consumer, execute the following command:

```bash
go run consumer.go
```


## FAQ

### 1. Why am I getting a "kafka: invalid configuration" error when running the producer?

If you encounter a "kafka: invalid configuration" error when running the producer, it's likely due to missing or incorrect configuration settings. One common cause of this error is forgetting to set the `Producer.Return.Successes` configuration to `true`. Make sure to include this configuration option and set it to `true` in your producer configuration.

### 2. Why am I getting a "client has run out of available brokers to talk to" error when running the consumer?

The "client has run out of available brokers to talk to" error usually indicates that the Kafka consumer client cannot establish a connection to the Kafka brokers. This could be due to various reasons such as incorrect broker address, networking issues, or Kafka cluster availability problems. Refer to the troubleshooting section in the README for potential solutions to this issue.

### 3. How can I scale the Kafka cluster and handle higher message throughput?

To scale the Kafka cluster and handle higher message throughput, you can:
- Add more Kafka brokers to distribute partitions across multiple nodes.
- Increase the number of consumer instances within the same consumer group to parallelize message processing.
- Adjust Kafka configuration settings such as the number of partitions per topic and replication factor to optimize cluster performance.

