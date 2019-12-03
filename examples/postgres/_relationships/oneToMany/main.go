package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// One to Many Relations

type User struct {
	gorm.Model
	Username string
	Calendar Calendar
}

type Calendar struct {
	gorm.Model
	Name        string
	UserID      uint
	Appointment []Appointment
}

type Appointment struct {
	gorm.Model
	Subject     string
	Description string
	StartTime   time.Time
	Length      uint
	CalendarID  uint
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
	db.DropTableIfExists(&Calendar{})
	db.DropTableIfExists(&Appointment{})

	db.CreateTable(&User{})
	db.CreateTable(&Calendar{})
	db.CreateTable(&Appointment{})

	// setup a foreign key with cascade update and cascade delete

	db.Debug().Model(&Calendar{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Debug().Save(&User{
		Username: "adam",
		Calendar: Calendar{
			Name: "Coding Events",
			Appointment: []Appointment{
				{Subject: "Swarm vs Kubernetes Workshop"},
				{Subject: "React Days workshop"},
			},
		},
	})

	println("done")
}
