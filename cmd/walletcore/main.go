package main

import (
	"context"
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nironwp/ms-wallet/internal/database"
	"github.com/nironwp/ms-wallet/internal/event"
	"github.com/nironwp/ms-wallet/internal/event/handler"
	createaccount "github.com/nironwp/ms-wallet/internal/usecase/create_account"
	createclient "github.com/nironwp/ms-wallet/internal/usecase/create_client"
	createtransaction "github.com/nironwp/ms-wallet/internal/usecase/create_transaction"
	"github.com/nironwp/ms-wallet/internal/web"
	"github.com/nironwp/ms-wallet/internal/web/webserver"
	"github.com/nironwp/ms-wallet/pkg/events"
	"github.com/nironwp/ms-wallet/pkg/kafka"
	"github.com/nironwp/ms-wallet/pkg/uow"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	defer db.Close()
	eventDispacher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()
	eventDispacher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispacher.Register("BalanceUpdated", handler.NewBalanceUpdatedKafkaHandler(kafkaProducer))
	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	context := context.Background()
	uow := uow.NewUow(context, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	transactionAccountUseCase := createtransaction.NewCreateTransactionUseCase(uow, eventDispacher, transactionCreatedEvent, balanceUpdatedEvent)

	webserver := webserver.NewWebServer(":6000")
	clienthandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewCreateAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*transactionAccountUseCase)

	webserver.AddHandler("/clients", clienthandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webserver.Start()
}
