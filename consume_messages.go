package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := flag.String("topic", "testtopic", "the topic to consume")
	flag.Parse()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   *topic,
	})
	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("ERROR: %v", err)
			os.Exit(1)
		}

		for _, h := range msg.Headers {
			fmt.Printf(`{"key": %q, "value": %q}`, h.Key, string(h.Value))
			fmt.Println()
		}

		fmt.Println(string(msg.Value))
		fmt.Println()
	}
}
