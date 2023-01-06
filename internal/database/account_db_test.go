package database

import (
	"database/sql"
	"testing"

	"github.com/brunomann/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("Bruno", "bruno@bruno.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindById() {
	s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES(?, ?, ?, ?)", s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)

	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)

	accountDb, err := s.accountDB.FindById(account.ID)

	s.Nil(err)
	s.Equal(accountDb.ID, account.ID)
	s.Equal(accountDb.Client.ID, account.Client.ID)
	s.Equal(accountDb.Client.Name, account.Client.Name)
	s.Equal(accountDb.Client.Email, account.Client.Email)
	s.Equal(accountDb.Balance, account.Balance)
}
