package main

import (
	"fmt"

	"GO_Practice/handlers"
	"GO_Practice/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	//lets connect to DB
	models.ConnectDB()
	//routes

	r.GET("/", handlers.Greetings)
	r.GET("/employees", handlers.AuthMiddleware(), handlers.GetEmployees) //Here this middleware is working, but it is unable to redirect to this handler... have to check about this..
	r.GET("/addresses", handlers.AuthMiddleware(), handlers.GetAddresses)
	r.GET("/address/:id", handlers.GetAddress)
	r.POST("/addemployee", handlers.AuthMiddleware(), handlers.AddEmployee)
	r.PUT("/updateemployee/:id", handlers.AuthMiddleware(), handlers.UpdateEmoployee)
	r.DELETE("/deleteemployee/:id", handlers.AuthMiddleware(), handlers.DeleteEmployee)

	r.GET("/register", handlers.Register)
	r.POST("/register", handlers.Register)
	r.GET("/login", handlers.LoginPage)
	r.POST("/login", handlers.Login)

	r.Run() //here we can even mention the path to run the server like the port .
	fmt.Println("hello World")
}
