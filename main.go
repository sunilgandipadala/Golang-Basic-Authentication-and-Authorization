package main

import (
	"GO_Practice/infastructure"
	"GO_Practice/migrations"
	"GO_Practice/routers"
	"fmt"
)

func main() {
	//lets connect to DB
<<<<<<< HEAD:main.go
	models.ConnectingDB()
=======
	infastructure.MysqlDatabase()
	migrations.Initialize()
>>>>>>> 956df004664b79d57c06f1f9f7bf8ac659b19ebb:gin_api.go
	//routes
	routers.Route()
	fmt.Println("hello World")
}
