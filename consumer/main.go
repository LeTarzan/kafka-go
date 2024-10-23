package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func consumeMessages(reader *kafka.Reader, id int) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		time.Sleep(time.Duration(1+id) * time.Second)

		if err != nil {
			log.Printf("Consumer %d failed to read message: %v", id, err)
			continue
		}

		log.Printf("Consumer %d received: %s = %s", id, string(msg.Key), string(msg.Value))

		err = reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Printf("Consumer %d failed to commit message: %v", id, err)
		}
	}
}

func newReader(TOPIC string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   TOPIC,
		GroupID: "groupid",
	})
}

func main() {
	fmt.Printf("starting consumer...")
	const TOPIC = "messages"

	reader := newReader(TOPIC)
	defer reader.Close()

	for idx := range 5 {
		go consumeMessages(reader, idx)
	}

	select {}
}
