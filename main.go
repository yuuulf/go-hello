package main

import (
	"os"

	"github.com/yuuulf/go-hello/pkg/db"
	"github.com/yuuulf/go-hello/pkg/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataSourceName := os.Getenv("DATA_SOURCE_NAME")
	if dataSourceName == "" {
		dataSourceName = "root:root@tcp(127.0.0.1:3306)/sample"
	}
	err := db.Init(dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	s := server.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = ":1323"
	}
	s.Start(port)
}
