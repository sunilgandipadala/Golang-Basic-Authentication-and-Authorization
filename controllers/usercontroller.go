package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"GO_Practice/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var Logged = false

// This method is to assign the values to the struct fields -Director
func NewUserRegister(builder models.UserBuilder) *models.RegisterUser {
	return &models.RegisterUser{Builder: builder}
}
func Register(c *gin.Context) {

	//will make the user register and we will store it into db
	var users map[string]interface{}
	var user models.User

	fmt.Println("You are in Registration Process")
	//var r *http.Request = c.Request
	fmt.Println("Request Method", c.Request.Method)
	if c.Request.Method == "POST" {

		if err := c.ShouldBind(&user); err != nil {
			fmt.Print("can't bind")
		}
		fmt.Println("Inside")

		inrec, _ := json.Marshal(user)
		json.Unmarshal(inrec, &users)

		dynamicBuilder := &models.DynamicUserBuilder{}
		new_user := NewUserRegister(dynamicBuilder)

		//it will become a new variable
		user := new_user.UserRegistration(users)
		name := user.Name
		email := user.Email
		age := user.Age
		phone := user.Phone
		password := user.Password
		cpassword := user.ConfirmPassword
		fmt.Print(age)
		if email == "" {
			c.HTML(200, "register.html", gin.H{
				"username":      name,
				"email":         email,
				"age":           user.Age,
				"phone":         phone,
				"password":      password,
				"cpassword":     cpassword,
				"error_message": "Invalid Data"})
			/**/
		} else {
			location := url.URL{Path: "/login"}
			c.Redirect(http.StatusFound, location.RequestURI())
		}
		//here we have to make redirection to login page if it is stored in db

	} else {
		c.HTML(200, "register.html", nil)
	}
}

func LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {
	fmt.Println("helo world")
	var r *http.Request = c.Request
	err := r.ParseForm()
	if err != nil {
		c.String(501, "No Form")
		return
	}
	current_user := r.FormValue("email")
	current_password := r.FormValue("password")
	var user models.User
	//registered users ..
	errr := db.Where(&models.User{Email: current_user, Password: current_password}).First(&user).Error
	if errr != nil {
		c.HTML(401, "login.html", gin.H{
			"error_message": "Invalid Login Credentials... Register to Login..", "email": current_user, "password": current_password,
		})
		return
	}
	if current_password == user.Password {
		fmt.Println("logged in")
		Logged = true
		//location := url.URL{Path: "/"}
		c.Redirect(http.StatusFound, "/")
		//c.Next()
	} else {
		c.HTML(401, "login.html", gin.H{
			"error_message": "Invalid Login Credentials... Register to Login..", "email": current_user, "password": current_password,
		})
	}

}
