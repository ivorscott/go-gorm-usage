package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "postgres://postgres:@localhost:5432/gorm-usage?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&Appointment{})
	db.CreateTable(&User{})
	db.CreateTable(&Appointment{})

	u := User{
		FirstName: "Arthur",
		LastName:  "Dent",
	}

	appointments := []Appointment{
		Appointment{Subject: "First"},
		Appointment{Subject: "Second", Attendees: []*User{&u}},
		Appointment{Subject: "Third"},
	}

	u.Appointments = appointments

	db.Debug().Create(&u)

	fmt.Println(db.NewRecord(&u))
}

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Appointments []Appointment
}

type Appointment struct {
	gorm.Model
	UserID      uint
	StartTime   *time.Time
	Duration    uint
	Attendees   []*User
	Subject     string
	Description string
}
