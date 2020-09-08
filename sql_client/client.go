package sql_client

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	isMocked bool
	dbClient SqlClientInterface
)

const (
	goEnv      = "GO_ENVIRONMENT"
	production = "production"
)

type client struct {
	db *sql.DB
}

type SqlClientInterface interface {
	QueryRow(query string, args ...interface{}) row
}

func isProduction() bool {
	return os.Getenv(goEnv) == production
}

func StartMockServer() {
	isMocked = true
}

func StopMockServer() {
	isMocked = false
}

func Open(driverName, dataSourceName string) (SqlClientInterface, error) {
	if isMocked && !isProduction() {
		dbClient = &clientMock{}
		fmt.Println("starting mock server")
		return dbClient, nil
	}
	if driverName == "" {
		return nil, errors.New("invalid driver name")
	}

	database, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, errors.New("connection error")
	}

	dbClient = &client{
		db: database,
	}

	return dbClient, nil
}
func (c client) QueryRow(query string, args ...interface{}) row {
	returnedRow := c.db.QueryRow(query, args...)
	if returnedRow == nil {
		return nil
	}

	result := sqlRow{
		row: returnedRow,
	}

	return &result
}
