package main

import (
	"database/sql"
	"fmt"

	"gin_api/handlers"
	"gin_api/models"

	"github.com/gin-gonic/gin"
)

var Database *sql.DB //why i delcared this here?

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	//lets connect to DB
	models.ConnectDB()
	//routes

	r.GET("/", handlers.Greetings)
	r.GET("/employees", handlers.GetEmployees)
	r.GET("/addresses", handlers.GetAddresses)
	r.GET("/address/:id", handlers.GetAddress)
	r.POST("/addemployee", handlers.AddEmployee)
	r.PUT("/updateemployee/:id", handlers.UpdateEmoployee)
	r.DELETE("/deleteemployee/:id", handlers.DeleteEmployee)

	r.Run() //here we can even mention the path to run the server like the port .
	fmt.Println("hello World")
}
