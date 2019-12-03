package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// One to One Relations

// Imagine a relationship between a User and a Calendar
// with cascade update and cascade delete requirements

// There are two ways to do this

// 1. Has-One - adding a user reference in the calendar
// a. the user owns the calendar reference
// b. deleting the calendar would not delete the user
// c. but deleting the user would delete the calendar

// 2. Belongs-To - adding a calendar reference in the user
// a. the user belongs to the calendar
// b. deleting the user record would not delete the calendar
// c. but deleting the calendar would delete the user

// note: foreignKey example shows cascade update and delete config

// Has-One Relationship example
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

// Belongs-To Relationship example

// (note: shouldn't be used in practice with cascade delete, just for demonstration
// Belongs-To is just the reverse of Has-One, so you're looking at the relationship
// in the opposite direction.)

// type User struct {
// 	gorm.Model
// 	Username   string
// 	Calendar   Calendar
// 	CalendarID uint
// }

// type Calendar struct {
// 	gorm.Model
// 	Name string
// }

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

	db.Debug().Save(&User{
		Username: "adent",
		Calendar: Calendar{
			Name: "Improbable Events",
		},
	})

	u := User{}
	c := Calendar{}
	db.First(&u).Related(&c, "calendar")
	fmt.Println(u)
	fmt.Println(c)

	println("done")
}
