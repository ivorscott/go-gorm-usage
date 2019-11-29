package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	gorm.Model
	Username  string `sql:"unique"`
	FirstName string `sql:"size:100"`
	LastName  string `sql:"type:VARCHAR(15)"`
	Count     int    `gorm:"AUTO_INCREMENT"`
	TempField int    `gorm:"-"`
}

func main() {
	db, err := gorm.Open("postgres", "postgres://postgres:@localhost:5432/gorm-usage?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	dbase := db.DB()
	defer dbase.Close()

	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})

	println("done")
}
