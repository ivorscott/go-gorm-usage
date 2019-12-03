package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Many to Many Relations

type User struct {
	gorm.Model
	Username string
	Calendar Calendar
}

type Calendar struct {
	gorm.Model
	Name        string
	UserID      uint
	Appointment []Appointment `gorm:"polymorphic:owner"`
}

type Appointment struct {
	gorm.Model
	Subject     string
	Description string
	StartTime   time.Time
	Length      uint
	OwnerID     uint
	OwnerType   string
	Attendees   []User `gorm:"many2many:appointment_user"`
}

type TaskList struct {
	gorm.Model
	Appointments []Appointment `gorm:"polymorphic:owner"`
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

	users := []User{
		{Username: "danwells"},
		{Username: "amygoodman"},
		{Username: "sammy"},
	}

	for i := range users {
		db.Save(&users[i])
	}

	db.Debug().Save(&User{
		Username: "adam",
		Calendar: Calendar{
			Name: "Coding Events",
			Appointment: []Appointment{
				{Subject: "Swarm vs Kubernetes Workshop", Attendees: users},
				{Subject: "React Days workshop", Attendees: users},
			},
		},
	})

	println("done")
}
