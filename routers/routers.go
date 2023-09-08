package routers

import (
	"GO_Practice/controllers"
	"GO_Practice/middleware"

	"github.com/gin-gonic/gin"
)

func Route() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", controllers.Greetings)
	r.GET("/employees", middleware.AuthMiddleware(), controllers.GetEmployees) //Here this middleware is working, but it is unable to redirect to this handler... have to check about this..

	//lets apply grouping routes

	r.GET("/addresses", middleware.AuthMiddleware(), controllers.GetAddresses)
	r.GET("/address/:id", controllers.GetAddress)
	r.POST("/addemployee", middleware.AuthMiddleware(), controllers.AddEmployee)
	r.PUT("/updateemployee/:id", middleware.AuthMiddleware(), controllers.UpdateEmoployee)
	r.DELETE("/deleteemployee/:id", middleware.AuthMiddleware(), controllers.DeleteEmployee)

	r.GET("/register", controllers.Register)
	r.POST("/register", controllers.Register)
	r.GET("/login", controllers.LoginPage)
	r.POST("/login", controllers.Login)

	//For testing Purpose
	r.GET("/testget", controllers.Testget)
	r.POST("/testpost", controllers.TestPost)
	r.Run() //here we can even mention the path to run the server like the port .
}
