package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	FirstName string
	LastName  string
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

	db.Model(&User{}).AddIndex("idx_first_name", "first_name")
	db.Model(&User{}).AddUniqueIndex("idx_last_name", "last_name")
	// db.Model(&User{}).RemoveIndex("idx_first_name")

	println("done")
}
