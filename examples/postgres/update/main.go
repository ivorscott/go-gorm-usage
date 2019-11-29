package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	ID        uint
	Username  string
	FirstName string
	LastName  string
}

var users []User = []User{
	User{Username: "adent", FirstName: "Arthur", LastName: "Dent"},
	User{Username: "fperfect", FirstName: "Ford", LastName: "Perfect"},
	User{Username: "tmacmillan", FirstName: "Tricia", LastName: "MacMillan"},
	User{Username: "mrobot", FirstName: "Marvin", LastName: "Robot"},
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

	for _, user := range users {
		db.Create(&user)
	}

	u := User{Username: "tmacmillan"}

	db.Where(&u).First(&u)
	fmt.Println(u)

	u.LastName = "Beeblebrox"

	db.Save(&u)

	user := User{}
	db.Where(&u).First(&user)
	fmt.Println(user)

	println("done")
}
