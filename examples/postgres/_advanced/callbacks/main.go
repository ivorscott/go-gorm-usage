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

	db.DropTableIfExists(&User{}, &Calendar{}, &Appointment{}, "appointment_user")
	db.AutoMigrate(&User{}, &Calendar{}, &Appointment{})

	user := &User{
		Username:  "adent",
		FirstName: "Arthur",
		LastName:  "Dent",
		Calendar:  Calendar{Name: "Arthur's Calendar"},
	}

	fmt.Println("Creating")
	db.Create(&user)

	user.Calendar.Name = "Arthur's Itinerary"

	fmt.Println("Updating")
	db.Save(&user)

	fmt.Println("Deleting")
	db.Delete(&user)

}

type User struct {
	gorm.Model
	Username  string
	FirstName string `sql:"type:VARCHAR(100)"`
	LastName  string
	Calendar  Calendar
}

func (u *User) BeforeSave() error {
	fmt.Println("Before Save")
	return nil
}
func (u *User) BeforeCreate() error {
	fmt.Println("Before Create")
	return nil
}
func (u *User) AfterSave() error {
	fmt.Println("After Save")
	return nil
}
func (u *User) AfterCreate() error {
	fmt.Println("After Create")
	return nil
}

type Calendar struct {
	gorm.Model
	Name         string
	UserID       uint `sql:"index:idx_calendar_user_id"`
	Appointments []*Appointment
}

type Appointment struct {
	gorm.Model
	Subject     string
	Description string
	StartTime   time.Time
	Length      uint
	CalendarID  uint    `sql:"index:idx_appointment_calendar_id"`
	Attendees   []*User `gorm:"many2many:appointment_user"`
}
