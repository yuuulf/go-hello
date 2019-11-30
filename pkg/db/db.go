package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var connection *sql.DB

func Init(dataSourceName string) error {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	connection = db
	return nil
}

func Conn() *sql.DB {
	return connection
}

func Close() error {
	return connection.Close()
}
