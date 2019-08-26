package main

import (
	"context"
	"flag"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

var (
	amount = flag.Uint("messages", 1000, "The number of messages to be written")
	uri    = flag.String("uri", "localhost:9092", "The kafka uri")
)

// NewKafkaWriter creates new instanse of kafka producer.
func NewKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

func main() {
	flag.Parse()

	writer := NewKafkaWriter(*uri, "benchmark-topic")
	defer writer.Close()

	for i := 0; i < int(*amount); i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("Key-%d", i)),
			Value: []byte(fmt.Sprint("Test")),
		}

		if err := writer.WriteMessages(context.Background(), msg); err != nil {
			fmt.Println(err)
		}
	}
}
