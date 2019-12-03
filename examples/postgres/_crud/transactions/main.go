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

	u := User{
		FirstName: "Marvin",
		LastName:  "Robot",
	}

	tx := db.Begin()

	if err = tx.Create(&u).Error; err != nil {
		tx.Rollback()
	}

	u.LastName = "The Happy Robot"
	if err = tx.Save(&u).Error; err == nil {
		tx.Rollback()
	}

	tx.Commit()
}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
}
