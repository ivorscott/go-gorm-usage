package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// gorm.Model will add properties to our struct
// e.g.
// ID        uint `gorm:"primary_key"`
// CreatedAt time.Time
// UpdatedAt time.Time
// DeletedAt *time.Time

// Note: When DeletedAt at is present is will change the behavior of the delete operation

type User struct {
	gorm.Model
	// ID        uint
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

	// remove automatic pluralization of table names
	// db.SingularTable(true)

	db.CreateTable(&User{})

	println("done")
}
