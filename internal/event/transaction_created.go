package event

import "time"

type TransactionCreated struct {
	Name    string
	Payload interface{}
}

func NewTransactionCreated() *TransactionCreated {
	return &TransactionCreated{
		Name: "TransactionCreated",
	}
}

func (event *TransactionCreated) GetName() string {
	return event.Name
}

func (event *TransactionCreated) GetDateTime() time.Time {
	return time.Now()
}

func (event *TransactionCreated) GetPayload() interface{} {
	return event.Payload
}

func (event *TransactionCreated) SetPayload(payload interface{}) {
	event.Payload = payload
}
