package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	PORT     = "5432"
	USER     = "postgres"
	PASSWORD = "qwerty"
	DB_NAME  = "postgres"
	DRIVER   = "postgres"
)

func getConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DB_NAME)
}

func NewConnection() *sql.DB {
	connection := getConnectionString()
	db, err := sql.Open(DRIVER, connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
