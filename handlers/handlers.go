package handlers

import (
	"fmt"
	"net/http"
	"net/url"

	"GO_Practice/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
)

var logged = false

var emp []models.Employees
var db = models.ConnectDB()
var object = func(c *gin.Context) {
	if err := db.Find(&emp).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch address"})
		return
	}
}

func Greetings(c *gin.Context) {
	c.HTML(http.StatusOK, "hello.html", nil)
}

// To read all the employees data into a structred JSON format
func GetEmployees(c *gin.Context) {
	object(c)
	c.HTML(http.StatusOK, "employees.html", gin.H{"employees": emp})
	//c.IndentedJSON(http.StatusOK, models.Employ)
}

// To get all the employees addresses
// this method is something like duplicative/repitative
func GetAddresses(c *gin.Context) {

	object(c)
	c.HTML(http.StatusOK, "employes_addresses.html", gin.H{"employees": emp})
}

// To get address of a particular Employee
func GetAddress(c *gin.Context) {

	emp_id := c.Param("id")
	object(c)
	for _, employ := range emp {
		if employ.Id == emp_id {
			c.HTML(200, "employ_address.html", gin.H{"emp_name": employ.Name,
				"emp_id":      employ.Id,
				"emp_address": employ.Address})
			return
		}
	}
	c.HTML(http.StatusNotFound, "notfound.html", nil)
}

func AddEmployee(c *gin.Context) {

	var newEmployee models.Employees
	if err := c.MustBindWith(&newEmployee, binding.JSON); err != nil {
		return
	}
	//models.Employ = append(models.Employ, newEmployee)
	db.Create(&newEmployee)

	fmt.Println("Employee added sucssessfully")
	c.String(301, "Employee Added Successfully")
	//Here lets define redirect to employees page..
	c.Redirect(301, "/employees")
}

func UpdateEmoployee(c *gin.Context) {
	id := c.Param("id")
	object(c)
	var flag = 0
	var employ models.Employees
	for _, employ := range emp {
		if id == employ.Id {
			employ.Id = id
			flag = 1
			break
		}
	}
	if flag == 0 {
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}
	var updatedEmployee models.Employees
	if err := c.BindJSON(&updatedEmployee); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	employ = updatedEmployee
	db.Save(&employ)
	c.IndentedJSON(200, updatedEmployee)

}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	object(c)
	status := db.Where("id = ?", id).Delete(&emp)
	fmt.Print(status)
	//here we have to modify the condition
	if status == nil {
		c.JSON(500, gin.H{"error": "No records Found"})
		return
	}
	c.String(200, "Deleted Succesfully")

}

func Register(c *gin.Context) {

	//will make the user register and we will store it into db
	var user models.User

	fmt.Println("You are in Registration Process")
	//var r *http.Request = c.Request
	//fmt.Println("Outside method....")
	fmt.Println("Request Method", c.Request.Method)
	if c.Request.Method == "POST" {
		//now here needs to save the data recived from the form to db
		if err := c.ShouldBind(&user); err != nil {
			fmt.Print("can't bind")
		}
		/*
			//In this way... we can parse the form and store it into db.. There is another way, directly we get
			Key,Value Pair using r.Form()

				err := r.ParseForm()
				if err != nil {
					c.String(501, "No Form")
					return
				}
				fmt.Println("Inside")
				//fmt.Println(r.Form)
				name := r.FormValue("username")
				email := r.FormValue("email")
				age := r.FormValue("age")
				phone := r.FormValue("phone")
				password := r.FormValue("password")
				cpassword := r.FormValue("cpassword")
				user.Name = name
				user.Email = email
				user.Age, err = strconv.Atoi(age)
				if err != nil {
					user.Age = 0
				}
				user.Phone = phone
				user.Password = password
				user.ConfirmPassword = cpassword*/
		//fmt.Println("\nuser..", user)

		err := db.Create(&user).Error
		//Here I have to pass user as key value pairs...
		name := user.Name
		email := user.Email
		age := user.Age
		phone := user.Phone
		password := user.Password
		cpassword := user.ConfirmPassword

		if err != nil {
			c.HTML(200, "register.html", gin.H{
				"username":      name,
				"email":         email,
				"age":           age,
				"phone":         phone,
				"password":      password,
				"cpassword":     cpassword,
				"error_message": err})
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
		logged = true
		//location := url.URL{Path: "/"}
		c.Redirect(http.StatusFound, "/")
	} else {
		c.HTML(401, "login.html", gin.H{
			"error_message": "Invalid Login Credentials... Register to Login..", "email": current_user, "password": current_password,
		})
	}

}

// This has to be checked,.....
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !logged {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			fmt.Println("finished...")
			return
		}
		c.Next()
	}
}
