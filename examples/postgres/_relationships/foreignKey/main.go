package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// One to One Relations

// Has-One Relationship
type User struct {
	gorm.Model
	Username string
	Calendar Calendar
}

type Calendar struct {
	gorm.Model
	Name   string
	UserID uint
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
	db.DropTableIfExists(&Calendar{})
	db.CreateTable(&User{})
	db.CreateTable(&Calendar{})

	// setup a foreign key with cascade update and cascade delete

	db.Debug().Model(&Calendar{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Debug().Save(&User{
		Username: "adent",
		Calendar: Calendar{
			Name: "Improbable Events",
		},
	})

	println("done")
}
