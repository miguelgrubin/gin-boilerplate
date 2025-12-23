package sharedmodule_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule"
	"github.com/stretchr/testify/suite"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

const (
	database = "foo"
	user     = "root"
	pass     = "password"
)

type NewDbConnectionTestSuite struct {
	suite.Suite
	ctx               context.Context
	postgresContainer *postgres.PostgresContainer
	mysqlContainer    *mysql.MySQLContainer
}

func (suite *NewDbConnectionTestSuite) SetupTestSuite() {
	suite.ctx = context.Background()

	postgresContainer, _ := postgres.Run(suite.ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(database),
		postgres.WithUsername(user),
		postgres.WithPassword(pass),
		postgres.BasicWaitStrategies(),
	)
	suite.postgresContainer = postgresContainer

	mysqlContainer, _ := mysql.Run(suite.ctx,
		"mysql:8.0.36",
		mysql.WithDatabase(database),
		mysql.WithUsername(user),
		mysql.WithPassword(pass),
	)
	suite.mysqlContainer = mysqlContainer
}

func (suite *NewDbConnectionTestSuite) TearDownTestSuite() {
	testcontainers.CleanupContainer(suite.T(), suite.postgresContainer)
	testcontainers.CleanupContainer(suite.T(), suite.mysqlContainer)
}

func (suite *NewDbConnectionTestSuite) TestNewDbConnectionWithPostgres() {
	host, _ := suite.postgresContainer.Host(suite.ctx)
	port, _ := suite.postgresContainer.MappedPort(suite.ctx, "5432")
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		pass,
		database,
		port.Port(),
	)

	db := sharedmodule.NewDbConnection("postgres", connectionString)

	suite.NotNil(db, "Expected DB connection to be established")
}

func (suite *NewDbConnectionTestSuite) TestNewDbConnectionWithPostgresConnectionError() {
	suite.Panics(func() {
		sharedmodule.NewDbConnection("postgres", "invalid-connection-string")
	}, "Expected panic on invalid Postgres connection string")
}

func (suite *NewDbConnectionTestSuite) TestNewDbConnectionWithSqlite() {
	db := sharedmodule.NewDbConnection("sqlite3", ":memory:")

	suite.NotNil(db, "Expected DB connection to be established")
}

func (suite *NewDbConnectionTestSuite) TestNewDbConnectionWithSqliteConnectionError() {
	suite.Panics(func() {
		sharedmodule.NewDbConnection("sqlite3", "/invalid/path/to/db.sqlite")
	}, "Expected panic on invalid SQLite DB path")
}

func (suite *NewDbConnectionTestSuite) TestNewDbConnectionWithMysql() {
	host, _ := suite.mysqlContainer.Host(suite.ctx)
	port, _ := suite.mysqlContainer.MappedPort(suite.ctx, "3306")
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		user,
		pass,
		host,
		port.Port(),
		database,
	)

	db := sharedmodule.NewDbConnection("mysql", connectionString)

	suite.NotNil(db, "Expected DB connection to be established")
}

func (suite *NewDbConnectionTestSuite) TestNewDbConnectionWithMysqlConnectionError() {
	suite.Panics(func() {
		sharedmodule.NewDbConnection("mysql", "invalid-connection-string")
	}, "Expected panic on invalid MySQL connection string")
}

func TestNewDbConnection(t *testing.T) {
	dbConnSuite := new(NewDbConnectionTestSuite)
	dbConnSuite.SetupTestSuite()
	suite.Run(t, dbConnSuite)
	dbConnSuite.TearDownTestSuite()
}
