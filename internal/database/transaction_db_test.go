package database

import (
	"database/sql"
	"testing"

	"github.com/nironwp/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	transactionDB TransactionDB
	accountFrom   *entity.Account
	accountTo     *entity.Account
}

func (t *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")

	t.Nil(err)

	t.db = db
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance float(8), created_at date)")
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float(8), created_at date)")
	clientDb := NewClientDB(db)
	accountDb := NewAccountDB(db)
	clientFrom, err := entity.NewClient("John", "john@gmail.com")
	t.Nil(err)
	err = clientDb.Save(clientFrom)
	t.Nil(err)

	accountFrom, err := entity.NewAccount(clientFrom)
	t.Nil(err)
	accountFrom.Credit(100)
	err = accountDb.Save(accountFrom)
	t.Nil(err)
	t.accountFrom = accountFrom

	clientTo, err := entity.NewClient("pedro", "pedro@gmail.com")
	t.Nil(err)
	err = clientDb.Save(clientTo)
	t.Nil(err)

	accountTo, err := entity.NewAccount(clientTo)
	t.Nil(err)
	accountTo.Credit(100)
	err = accountDb.Save(accountTo)
	t.Nil(err)

	t.accountTo = accountTo

	t.transactionDB = *NewTransactionDB(db)
}

func (t *TransactionDBTestSuite) TearDownSuite() {
	defer t.db.Close()
	t.db.Exec("DROP table accounts")
	t.db.Exec("DROP table clients")
	t.db.Exec("DROP table transactions")
}

func TestTransactionDBSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (suite *TransactionDBTestSuite) TestSave() {
	transaction, err := entity.NewTransaction(suite.accountFrom, suite.accountTo, 100)
	suite.Nil(err)

	err = suite.transactionDB.Create(transaction)
	suite.Nil(err)
}
