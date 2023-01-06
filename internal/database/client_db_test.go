package database

import (
	"database/sql"
	"testing"

	"github.com/brunomann/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSave() {
	client, _ := entity.NewClient("Bruno Mann", "bruno@bruno.com")
	err := s.clientDB.Save(client)

	s.Nil(err)
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("Bruno Mann", "bruno@bruno.com")
	s.clientDB.Save(client)

	clientDb, err := s.clientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(clientDb.ID, client.ID)
	s.Equal(clientDb.Name, client.Name)
	s.Equal(clientDb.Email, client.Email)
}
