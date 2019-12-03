package main

import (
	"fmt"

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
		FirstName: "Ford",
		LastName:  "Prefect",
	}

	db.Create(&u)

	fmt.Println(u)
	fmt.Println()

	db.Debug().Model(&u).UpdateColumns(
		map[string]interface{}{
			"first_name": "Zaphod",
			"last_name":  "Beeblebrox",
		})

	fmt.Println(u)
}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
}

func (u *User) BeforeUpdate() error {
	println("Before Update")

	return nil
}

func (u *User) AfterUpdate() error {
	println("After Update")

	return nil
}
