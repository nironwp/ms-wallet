package handler

import (
	"sync"

	"github.com/nironwp/ms-wallet/pkg/events"
	"github.com/nironwp/ms-wallet/pkg/kafka"
	log "github.com/sirupsen/logrus"
)

type BalanceUpdatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewBalanceUpdatedKafkaHandler(kafka *kafka.Producer) *BalanceUpdatedKafkaHandler {
	return &BalanceUpdatedKafkaHandler{
		Kafka: kafka,
	}
}

func (h *BalanceUpdatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	err := h.Kafka.Publish(message, nil, "balances")
	if err != nil {
		log.Errorf("Failed to publish message: %s", err)
		// You can choose to return the error here or handle it in another way
	}

	log.Infof("BalanceUpdatedKafkaHandler: %s", message.GetPayload())
}
