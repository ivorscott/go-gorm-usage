package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	ID        uint
	Username  string
	FirstName string
	LastName  string
}

// There are two ways to change table names
// 1. override table name with receiver function, (e.g., func (u User) TableName() string)
// 2. remove automatic pluralization of table names (e.g., db.SingularTable(true))

// Override table name to something specific
// func (u User) TableName() string {
// 	return "test_users"
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

	// remove automatic pluralization of table names
	// db.SingularTable(true)

	db.CreateTable(&User{})

	println("done")
}
