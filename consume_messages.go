package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/segmentio/kafka-go"
)

var wg sync.WaitGroup

func main() {
	topic := flag.String("topic", "testtopic", "the topic to consume")
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   *topic,
		GroupID: "myreader",
	})
	go func() {
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
	}()

	<-sigs
	go func() {
		r.Close()
		wg.Done()
	}()

	wg.Wait()
}
