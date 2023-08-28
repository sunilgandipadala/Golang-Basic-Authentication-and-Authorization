package main

import (
	"GO_Practice/models"
	"GO_Practice/routers"
	"fmt"
)

func main() {
	//lets connect to DB
	models.ConnectDB()
	//routes
	routers.Route()
	fmt.Println("hello World")
}
