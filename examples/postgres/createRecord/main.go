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

func main() {
	db, err := gorm.Open("postgres", "postgres://ivorscott:@localhost:5432/gorm-usage?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	dbase := db.DB()
	defer dbase.Close()

	db.DropTable(&User{})
	db.CreateTable(&User{})

	user := User{
		Username:  "ivoscott",
		FirstName: "Ivor",
		LastName:  "Scott",
	}

	db.Create(&user)

	fmt.Println(user)
	println("done")
}
