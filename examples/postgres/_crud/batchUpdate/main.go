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
		LastName:  "Dent",
		Salary:    50000,
	})
	db.Create(&User{
		FirstName: "Arthur",
		LastName:  "Dent",
		Salary:    30000,
	})

	db.Debug().Table("users").Where("salary > ?", 40000).
		Update("salary", gorm.Expr("salary + 5000"))
}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Salary    int
}
