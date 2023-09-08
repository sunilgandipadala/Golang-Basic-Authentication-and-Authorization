package models

import (
	"encoding/json"
	"fmt"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database1 struct {
	ID   int    `json:"id" gorm:"Not Null"`
	Name string `json:"name"`
}
type Database2 struct {
	gorm.Model
	Name   string `json:"name"`
	Emp_id string `json:"emp_id"`
}
type Database3 struct {
	UserName string      `json:"user"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Blogs    []Database4 `json:"blogs" gorm:"-"`
	BlogsDb  string      `json:"-" gorm:"column:blogs"`
}

// here for Database3 to store in DB , requrie Patch In PatchOut...

type Database4 struct {
	//here need to modify gorm id into blog id cloumn
	gorm.Model
	Title string `json:"title" gorm:"Unique"`
	Body  string `json:"body"`
}

// auto migration using atlas
func ConnectingDB() *gorm.DB {
	mysql_conn := "root:Sunil@513@/practice_db"
	db, err := gorm.Open(mysql.Open(mysql_conn), &gorm.Config{})
	CheckError(err)
	fmt.Println(db)
	fmt.Println("Data base Connected")
	//Here we will test...

	//lets load each struct individaully...

	//these 2 which were already migrated and having data in it in the db.
	stmts, err := gormschema.New("mysql").Load(&Employees{})
	//stmts, err := gormschema.New("mysql").Load(&User{})

	//These are fresh DB's
	//we can even use Automigration also
	//stmts, err := gormschema.New("mysql").Load(&Database1{})
	//stmts, err := gormschema.New("mysql").Load(&Database2{})
	//stmts, err := gormschema.New("mysql").Load(&Database3{})
	//stmts, err := gormschema.New("mysql").Load(&Database4{})
	/*if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(stmts)*/
	fmt.Println(stmts)
	//While using Load -- no need of Automigration - as both Load() and Automigration works similar..
	//But everytime when we changed the Struct - we should generate migrations by using
	/*
		// 	atlas migrate diff [--env gorm] here we can pass parameters or a file which holds all the parameters - atlas.hcl
		//then atlas migrate apply  - to apply the migrations

	*/
	fmt.Println("Table Created")
	return db

}

// -----------patchIn \&Out methods to update db
var j []byte

func (data *Database3) PatchIn() {
	var err error
	if data.Blogs != nil {
		j, err = json.Marshal(data.Blogs)
		//fmt.Println(j)
		if err != nil {
			fmt.Println("unable to parse response model")
		}
		data.BlogsDb = string(j)
	}
}

func (data *Database3) PatchOut() {
	//the issue is here... for the first time it is unmarshing safely.. but later its not working as j.. is losing the byte array
	if []byte(data.BlogsDb) != nil {
		err := json.Unmarshal([]byte(data.BlogsDb), &data.Blogs)
		if err != nil {
			fmt.Println("Unable to Unmarshal the data")
		}
	}
}

//---------------upto here
