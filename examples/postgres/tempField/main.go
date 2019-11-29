package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Some apps use transient fields that don't need to be save
// Temporary fields can be specified with a "-" (e.g., `gorm:"-"`)

type User struct {
	gorm.Model
	Username  string `sql:"type:VARCHAR(15)"`
	FirstName string `sql:"size:100"`
	LastName  string
	Count     int `gorm:"AUTO_INCREMENT"`
	TempField int `gorm:"-"`
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
