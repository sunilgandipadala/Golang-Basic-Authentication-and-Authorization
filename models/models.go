package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Adress struct {
	District string `json:"district"`
	State    string `gorm:"-" json:"state"`
	Pincode  int    `json:"pincode"`
}

type Employees struct {
	Name    string `json:"name" gorm:"not null"`
	Id      string `gorm:"primarykey;" json:"id"`
	Gender  string `json:"gender"`
	Role    string `json:"role"`
	Address Adress `json:"address" gorm:"embedded"`
}

type User struct {
	Name            string `json:"username" form:"username"  gorm:"type:text"`
	Email           string `json:"email" form:"email" gorm:"unique"`
	Age             int    `json:"age" form:"age" `
	Phone           string `json:"phone" form:"phone"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"cpassword" form:"cpassword"`
}

var user = User{Name: "Sunil Kumar", Email: "sgandipadala@zetaglobal.com", Age: 21, Phone: "+918184811287", Password: "sunil@90", ConfirmPassword: "sunil@90"}

// Here we are giving the data manually, we can retrieve data from database too
var Employ = []Employees{
	{Name: "Sunil", Id: "APT513", Gender: "Male", Role: "Associate Technology", Address: Adress{District: "Vizag", State: "AP", Pincode: 535281}},
	{Name: "Sk", Id: "APT514", Gender: "Male", Role: "Associate Technology", Address: Adress{District: "Vizag", State: "Andhra", Pincode: 535282}},
	{Name: "Kumar", Id: "APT515", Gender: "Male", Role: "Associate Technology", Address: Adress{District: "Vizag", State: "AP", Pincode: 535283}},
	{Name: "Mauli", Id: "APT516", Gender: "Male", Role: "Associate Technology", Address: Adress{District: "VZM", State: "AP", Pincode: 535284}},
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
	//db.Create(&user
	//db.Create(&Employ) //Here we stored the data into the database initially
	fmt.Println("Table Created")
	return db

}
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
