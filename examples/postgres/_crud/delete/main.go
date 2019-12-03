package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "postgres://postgres:@localhost:5432/gorm-usage?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})

	db.Create(&User{
		FirstName: "Tricia",
		LastName:  "MacMillan-Dent",
	})
	db.Create(&User{
		FirstName: "Arthur",
		LastName:  "MacMillan-Dent",
	})

	db.Debug().Where("last_name LIKE ?", "Mac%").Delete(&User{})
}

type User struct {
	ID        uint
	FirstName string
	LastName  string
}
