package handler

import (
	"sync"

	"github.com/nironwp/ms-wallet/pkg/events"
	"github.com/nironwp/ms-wallet/pkg/kafka"
	log "github.com/sirupsen/logrus"
)

type TransactionCreatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{
		Kafka: kafka,
	}
}

func (h *TransactionCreatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	err := h.Kafka.Publish(message, nil, "transactions")
	if err != nil {
		log.Errorf("Failed to publish message: %s", err)
		// You can choose to return the error here or handle it in another way
	}

	log.Infof("TransactionCreatedKafkaHandler: %s", message.GetPayload())
}
