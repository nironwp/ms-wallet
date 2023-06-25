package main

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/nironwp/ms-wallet/pkg/kafka"
)

func main() {
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	consumer := kafka.NewConsumer(&configMap, []string{"transactions", "balances"})

	msgChan := make(chan *ckafka.Message)

	go consumer.Consume(msgChan)

	for msg := range msgChan {
		fmt.Sprintf("Received message: %s\n", string(msg.Value))
	}
}
