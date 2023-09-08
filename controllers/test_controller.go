package controllers

import (
	"GO_Practice/models"

	"github.com/gin-gonic/gin"
)

func Testget(c *gin.Context) {
	var databasedata models.Database1
	test_db := models.ConnectingDB()

	if err := test_db.Find(&databasedata).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch address"})
		return
	}
	//databasedata.PatchOut() -- this is used when Database3 is used..
	c.IndentedJSON(200, &databasedata)

}

func TestPost(c *gin.Context) {
	var datadatabase models.Database1
	if err := c.ShouldBindJSON(&datadatabase); err != nil {
		c.String(500, "Invalid JSON data, check again..")
		return
	}
	//datadatabase.PatchIn() -- it is only used when Database3 is used
	if err := db.Create(&datadatabase).Error; err != nil {
		c.String(501, err.Error())
		return
	}
	c.String(201, "Object Added")
}
