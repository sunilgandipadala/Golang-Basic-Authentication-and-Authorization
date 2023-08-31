package main

import (
	"GO_Practice/infastructure"
	"GO_Practice/migrations"
	"GO_Practice/routers"
	"fmt"
)

func main() {
	//lets connect to DB
	infastructure.MysqlDatabase()
	migrations.Initialize()
	//routes
	routers.Route()
	fmt.Println("hello World")
}
