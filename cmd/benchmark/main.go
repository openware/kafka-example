package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/openware/kafka-benchmark/pkg/example"
)

var (
	amount  = flag.Uint("messages", 1000, "The number of messages to be written")
	brokers = flag.String("brokers", "localhost:19092,localhost:29092,localhost:39092", "Kafka brokers urls")
	client  = flag.String("client-id", "benchmark-kafka-client", "ID of kafka client")
	topic   = flag.String("topic", "benchmark", "Kafka topic to be used")
)

func LogError(text string) {
	fmt.Fprintf(os.Stderr, text)
}

func main() {
	flag.Parse()

	producer, err := example.Configure(strings.Split(*brokers, ","), *client, *topic)
	if err != nil {
		LogError(err.Error())
		return
	}
	defer producer.Close()

	start := time.Now()

	for i := 0; i < int(*amount); i++ {
		value := fmt.Sprintf("Message %d", i)
		log.Println("Message published!")
		if err = example.Publish(nil, nil, []byte(value)); err != nil {
			LogError(err.Error())
		}
	}

	fmt.Printf("Execution time: %s\n", time.Since(start))
}
