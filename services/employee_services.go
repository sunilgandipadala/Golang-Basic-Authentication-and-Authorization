package services

import (
	"GO_Practice/infastructure"
	"GO_Practice/models"
)

func GetEmployData() (models.Employees,err error) {
	var emp []models.Employees
	var db = infastructure.DB
	if err := db.Find(&emp).Error; err != nil {
		return emp, "No Employee Data Found"
	}
	return emp,nil
}
