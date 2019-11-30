package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	Model     gorm.Model `gorm:"embedded"`
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

	for _, f := range db.NewScope(&User{}).Fields() {
		println(f.Name)
	}

	println("done")
}
