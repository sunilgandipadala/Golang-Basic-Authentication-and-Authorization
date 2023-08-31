package infastructure

import (
	"GO_Practice/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MysqlDatabase(DbConfig config.Database) {
	dsn := DbConfig.DbUserName + ":" + DbConfig.DbPassword + "@(" + DbConfig.DbHost + ")/" + DbConfig.DbDatabase + "?charset=utf8&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(mysql_conn), &gorm.Config{})
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		c.String(200, "Database Connection Failed")
		return
	}
}
