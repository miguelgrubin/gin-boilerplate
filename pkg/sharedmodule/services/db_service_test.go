package services_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
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

type DBServiceTestSuite struct {
	suite.Suite
	ctx               context.Context
	postgresContainer *postgres.PostgresContainer
	mysqlContainer    *mysql.MySQLContainer
}

func (suite *DBServiceTestSuite) SetupTestSuite() {
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

func (suite *DBServiceTestSuite) TearDownTestSuite() {
	testcontainers.CleanupContainer(suite.T(), suite.postgresContainer)
	testcontainers.CleanupContainer(suite.T(), suite.mysqlContainer)
}

func (suite *DBServiceTestSuite) TestNewDbConnectionWithPostgres() {
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

	dbService := services.NewDBServiceGorm(services.DatabaseConfig{
		Driver:  "postgres",
		Address: connectionString,
	})
	defer dbService.Close()
	err := dbService.Connect()

	suite.Nil(err, "Expected no error on Postgres connection")
}

func (suite *DBServiceTestSuite) TestNewDbConnectionWithPostgresConnectionError() {
	dbService := services.NewDBServiceGorm(services.DatabaseConfig{
		Driver:  "postgres",
		Address: "invalid-connection-string",
	})

	suite.Panics(func() {
		dbService.Connect()
	}, "Expected panic on invalid Postgres connection string")
}

func (suite *DBServiceTestSuite) TestNewDbConnectionWithSqlite() {
	dbService := services.NewDBServiceGorm(services.DatabaseConfig{
		Driver:  "sqlite3",
		Address: ":memory:",
	})

	err := dbService.Connect()

	suite.Nil(err, "Expected DB connection to be established")
}

func (suite *DBServiceTestSuite) TestNewDbConnectionWithSqliteConnectionError() {
	dbService := services.NewDBServiceGorm(services.DatabaseConfig{
		Driver:  "sqlite3",
		Address: "/invalid/path/to/db.sqlite",
	})
	suite.Panics(func() {
		dbService.Connect()
	}, "Expected panic on invalid SQLite DB path")
}

func (suite *DBServiceTestSuite) TestNewDbConnectionWithMysql() {
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

	dbService := services.NewDBServiceGorm(services.DatabaseConfig{
		Driver:  "mysql",
		Address: connectionString,
	})
	defer dbService.Close()
	err := dbService.Connect()

	suite.Nil(err, "Expected DB connection to be established")
}

func TestNewDbConnection(t *testing.T) {
	dbConnSuite := new(DBServiceTestSuite)
	dbConnSuite.SetupTestSuite()
	suite.Run(t, dbConnSuite)
	dbConnSuite.TearDownTestSuite()
}
