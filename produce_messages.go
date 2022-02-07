package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := flag.String("topic", "testtopic", "the topic to produce to")
	key := flag.String("key", "k", "the key of the message")
	value := flag.String("value", "v", "the value of the message")
	count := flag.Int("count", 1, "how many times to repeat the message")
	flag.Parse()

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   *topic,
	})
	defer w.Close()

	msgs := make([]kafka.Message, 0, *count)

	for i := 0; i < *count; i++ {
		msgs = append(msgs, kafka.Message{
			Key:   []byte(*key),
			Value: []byte(*value),
			Headers: []kafka.Header{
				{Key: "event_type", Value: []byte("availability_status")},
			},
		},
		)
	}

	err := w.WriteMessages(context.Background(), msgs...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "AAA failed %v\n", err)
		os.Exit(1)
	}
}
