package main

import (
	"github.com/yuuulf/go-hello/pkg/db"
	"github.com/yuuulf/go-hello/pkg/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := db.Init("root:root@tcp(127.0.0.1:3307)/sample")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	s := server.New()
	s.Start()
}
