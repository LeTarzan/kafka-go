package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

func produceMessage(writer *kafka.Writer, id int) {
	for range rand.Intn(100) {
		msg := fmt.Sprintf("message from producer %d: %d", id, rand.Intn(100))

		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("key-%d", id)),
				Value: []byte(msg),
			},
		)

		if err != nil {
			fmt.Printf("Failed to write message from producer %d: %v", id, err)
		} else {
			fmt.Printf("Producer %d sent message: %s", id, msg)
		}
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	}
}

func newWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:  kafka.TCP("kafka:9092"),
		Topic: topic,
	}
}

func main() {
	fmt.Printf("starting producer...")

	const TOPIC = "messages"

	writer := newWriter(TOPIC)
	defer writer.Close()

	for idx := range 5 {
		go produceMessage(writer, idx)
	}

	select {}
}
