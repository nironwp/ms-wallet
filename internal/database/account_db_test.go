package database

import (
	"database/sql"
	"testing"

	"github.com/nironwp/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db           *sql.DB
	accountDb    AccountDB
	first_client *entity.Client
}

func (a *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")

	a.Nil(err)

	a.db = db
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance float(8), created_at date)")
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	first_client, _ := entity.NewClient("John", "john@gmail.com")
	clientDb := NewClientDB(db)
	a.first_client = first_client
	err = clientDb.Save(first_client)
	a.Nil(err)

	a.accountDb = *NewAccountDB(db)
}

func (a *AccountDBTestSuite) TearDownSuite() {
	defer a.db.Close()
	a.db.Exec("DROP table accounts")
	a.db.Exec("DROP table clients")
}

func TestAccountDBSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (suite *AccountDBTestSuite) TestSave() {
	account, err := entity.NewAccount(suite.first_client)

	suite.Nil(err)
	suite.NotNil(account)

	err = suite.accountDb.Save(account)
	suite.Nil(err)
}

func (suite *AccountDBTestSuite) TestFindById() {
	suite.db.Exec("Insert into clients (id, name, email, created_at) values (?, ?, ?, ?)",
		suite.first_client.ID, suite.first_client.Name, suite.first_client.Email, suite.first_client.CreatedAt,
	)

	account, err := entity.NewAccount(suite.first_client)
	suite.Nil(err)
	suite.NotNil(account)

	err = suite.accountDb.Save(account)
	suite.Nil(err)

	accountDb, err := suite.accountDb.FindById(account.ID)

	suite.Nil(err)

	suite.Equal(account.ID, accountDb.ID)
	suite.Equal(account.Balance, accountDb.Balance)

	suite.Equal(account.Client.ID, accountDb.Client.ID)
	suite.Equal(account.Client.Name, accountDb.Client.Name)
	suite.Equal(account.Client.Email, accountDb.Client.Email)
}
