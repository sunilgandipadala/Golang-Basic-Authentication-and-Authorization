package migrations

import (
	"GO_Practice/infastructure"
	"GO_Practice/models"
)

func Initialize() (err error) {
	return infastructure.DB.AutoMigrate(&models.Employees{}, &models.User{})
}
