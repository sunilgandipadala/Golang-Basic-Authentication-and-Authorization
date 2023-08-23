package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"gin_api/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Greetings(c *gin.Context) {
	c.HTML(http.StatusOK, "hello.html", nil)
}

// To read all the employees data into a structred JSON format
func GetEmployees(c *gin.Context) {
	/*employees, err := db.Query("SELECT * FROM Employees")
	if err != nil {
		log.Fatal(err)
	}*/
	c.IndentedJSON(http.StatusOK, models.Employ)
}

// To get all the employees addresses
func GetAddresses(c *gin.Context) {
	for _, emp := range models.Employ {
		c.IndentedJSON(http.StatusOK, emp.Address)
	}
}

// To get address of a particular Employee
func GetAddress(c *gin.Context) {
	emp_id := c.Param("id")

	for _, emp := range models.Employ {
		if emp.Id == emp_id {
			c.IndentedJSON(http.StatusOK, emp.Name)
			c.IndentedJSON(http.StatusOK, emp.Address)
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
	models.Employ = append(models.Employ, newEmployee)

	fmt.Println("Employee added sucssessfully")
	c.IndentedJSON(http.StatusCreated, newEmployee)
}

func UpdateEmoployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employees
	var j int
	for i, emp := range models.Employ {
		if id == emp.Id {
			employee = models.Employ[i]
			j = i
			break
		}
	}
	if employee.Id == "" {
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}
	var updatedEmployee models.Employees
	if err := c.BindJSON(&updatedEmployee); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	models.Employ[j] = updatedEmployee

	c.IndentedJSON(200, updatedEmployee)

}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	var j = -1
	for i, emp := range models.Employ {
		if emp.Id == id {
			j = i
			break
		}
	}
	if j == -1 {
		c.HTML(http.StatusNotFound, "notfound.html", nil)
		return
	}
	models.Employ = append(models.Employ[:j], models.Employ[j+1:]...)
	c.IndentedJSON(200, models.Employ)

}
