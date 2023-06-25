package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nironwp/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type ClientDbTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDb *ClientDB
}

func (s *ClientDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")

	s.Nil(err)

	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDb = NewClientDB(db)
}

func (s *ClientDbTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP table clients")
}

func TestClientDbTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDbTestSuite))
}

func (s *ClientDbTestSuite) TestGet() {
	client, _ := entity.NewClient("John", "john@gmail.com")
	s.clientDb.Save(client)

	clientDb, _ := s.clientDb.Get(client.ID)
	s.Equal(clientDb.Name, client.Name)
	s.Equal(client.ID, clientDb.ID)
	s.Equal(client.Email, clientDb.Email)
	s.Equal(client.Name, clientDb.Name)
}

func (s *ClientDbTestSuite) TestSave() {
	client, _ := entity.NewClient("John", "john@gmail.com")

	err := s.clientDb.Save(client)

	s.Nil(err)
}
