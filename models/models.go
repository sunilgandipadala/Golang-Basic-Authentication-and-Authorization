package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employees struct {
	Name    string `json:"name" gorm:"not null"`
	Id      string `gorm:"primarykey;" json:"id"`
	Gender  string `json:"gender"`
	Role    string `json:"role"`
	Address Adress `json:"address" gorm:"embedded"`
}
type Adress struct {
	District string `json:"district"`
	State    string `gorm:"-" json:"state"`
	Pincode  int    `json:"pincode"`
}

// these register users can only see the employees data and can add employees
type User struct {
	Name            string `json:"username" form:"username"  gorm:"type:text"`
	Email           string `json:"email" form:"email" gorm:"unique"`
	Age             int    `json:"age" form:"age" `
	Phone           string `json:"phone" form:"phone"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"cpassword" form:"cpassword"`
}

func ConnectDB() *gorm.DB {
	mysql_conn := "root:Sunil@513@/practice_db"
	db, err := gorm.Open(mysql.Open(mysql_conn), &gorm.Config{})
	CheckError(err)
	fmt.Println(db)
	CheckError(err)
	fmt.Println("Data base Connected")
	db.AutoMigrate(&Employees{})
	db.AutoMigrate(&User{})
	//db.Create(&user)
	//db.Create(&Employ) //Here we stored the data into the database initially
	fmt.Println("Table Created")
	return db

}
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
