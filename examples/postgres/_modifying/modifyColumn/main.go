package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "postgres://postgres:@localhost:5432/gorm-usage?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	//db.DropTableIfExists(&User{}, &Calendar{}, &Appointment{}, "appointment_user")
	db.DropTableIfExists("attachments")
	db.AutoMigrate(&User{}, &Calendar{}, &Appointment{})
	//db.Debug().Model(&User{}).ModifyColumn("first_name", "VARCHAR(100)")
	db.Debug().Exec(`ALTER TABLE "users" ALTER "first_name" TYPE VARCHAR(100)`)

}

type User struct {
	gorm.Model
	Username  string
	FirstName string `sql:"type:VARCHAR(100)"`
	LastName  string
	Calendar  Calendar
}

type Calendar struct {
	gorm.Model
	Name         string
	UserID       uint `sql:"index:idx_calendar_user_id"`
	Appointments []*Appointment
}

type Appointment struct {
	gorm.Model
	Subject           string
	Description       string
	StartTime         time.Time
	Length            uint
	CalendarID        uint `sql:"index:idx_appointment_calendar_id"`
	Recurring         bool
	RecurrencePattern string
	Attendees         []*User `gorm:"many2many:appointment_user"`
}
