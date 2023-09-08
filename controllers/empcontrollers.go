package controllers

import (
	"fmt"
	"net/http"

	"GO_Practice/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
)

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
	db.Find(&emp)
	c.HTML(http.StatusOK, "employees.html", gin.H{"employees": emp})
	//c.IndentedJSON(http.StatusOK, models.Employ)
}

// To get all the employees addresses
// this method is something like duplicative/repitative
func GetAddresses(c *gin.Context) {
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
