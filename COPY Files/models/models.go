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
	gorm.Model
	Name    string `json:"name"`
	Id      string `gorm:"unique;autoIncrement" json:"id"`
	Gender  string `json:"gender"`
	Role    string `json:"role"`
	Address Adress `json:"address"`
}

// Here we are giving the data manually, we can retrieve data from database too
var Employ = []Employees{
	{Name: "Sunil", Id: "APT513", Gender: "Male", Role: "Associate Technology", Address: Adress{"Vizag", "AP", 535281}},
	{Name: "Sk", Id: "APT514", Gender: "Male", Role: "Associate Technology", Address: Adress{"Vizag", "AP", 535282}},
	{Name: "Kumar", Id: "APT515", Gender: "Male", Role: "Associate Technology", Address: Adress{"Vizag", "AP", 535283}},
	{Name: "Mauli", Id: "APT516", Gender: "Male", Role: "Associate Technology", Address: Adress{"Vizag", "AP", 535284}},
}

func ConnectDB() {
	mysql_conn := "root:Sunil@513@/practice_db"
	db, err := gorm.Open(mysql.Open(mysql_conn), &gorm.Config{})
	CheckError(err)
	fmt.Println(db)
	CheckError(err)
	fmt.Println("Data base Connected")
	db.Create(&Employ)
	fmt.Println("Table Created")

}
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
