package example

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

var writer *kafka.Writer

var timeout = 10 * time.Second

// Publish publishes message into kafka.
func Publish(parent context.Context, key, value []byte) error {
	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}

	return writer.WriteMessages(context.Background(), message)
}

// Configure returns new Kafka writer for given params.
func Configure(brokers []string, client string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  timeout,
		ClientID: client,
	}

	config := kafka.WriterConfig{
		Brokers:          brokers,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     timeout,
		ReadTimeout:      timeout,
		CompressionCodec: snappy.NewCompressionCodec(),
	}

	w = kafka.NewWriter(config)
	writer = w
	return w, nil
}
