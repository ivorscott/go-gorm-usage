package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// We can specify the field type or size by using go's struct tags
// https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go

// The default size of a varying character (e.i., VARCHAR) is 255
// Setting a string field's size to 100 is equivalent to setting VARCHAR(100)

type User struct {
	gorm.Model
	Username  string `sql:"type:VARCHAR(15)"`
	FirstName string `sql:"size:100"`
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

	println("done")
}
