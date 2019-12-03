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
	seedDB(db)
	rows, _ := db.Debug().Model(&Appointment{}).Select("calendar_id, sum(length) as total_length").
		Group("calendar_id").Having("calendar_id = ?", 1).Rows()

	for rows.Next() {
		var id, length int
		err = rows.Scan(&id, &length)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, length)
	}

}

type UserViewModel struct {
	FirstName    string
	LastName     string
	CalendarName string
}

func seedDB(db *gorm.DB) {
	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&Calendar{})
	db.DropTableIfExists(&Appointment{})
	db.CreateTable(&User{})
	db.CreateTable(&Calendar{})
	db.CreateTable(&Appointment{})

	users := map[string]*User{
		"adent":       &User{Username: "adent", FirstName: "Arthur", LastName: "Dent"},
		"fprefect":    &User{Username: "fprefect", FirstName: "Ford", LastName: "Prefect"},
		"tmacmillan":  &User{Username: "tmacmillan", FirstName: "Tricia", LastName: "MacMillan"},
		"zbeeblebrox": &User{Username: "zbeeblebox", FirstName: "Zaphod", LastName: "Beeblebrox"},
		"mrobot":      &User{Username: "mrobot", FirstName: "Marvin", LastName: "Robot"},
	}
	for _, user := range users {
		user.Calendar = Calendar{Name: "Calendar"}
	}

	users["adent"].AddAppointment(&Appointment{
		Subject:   "Save House",
		StartTime: parseTime("1979-07-02 08:00"),
		Length:    60,
	})
	users["fprefect"].AddAppointment(&Appointment{
		Subject:   "Get a Drink at Local Pub",
		StartTime: parseTime("1979-07-02 10:00"),
		Length:    11,
		Attendees: []*User{users["adent"]},
	})
	users["fprefect"].AddAppointment(&Appointment{
		Subject:   "Hitch a Ride",
		StartTime: parseTime("1979-07-02 10:12"),
		Length:    60,
		Attendees: []*User{users["adent"]},
	})
	users["fprefect"].AddAppointment(&Appointment{
		Subject:   "Attend Poetry Reading",
		StartTime: parseTime("1979-07-02 11:00"),
		Length:    30,
		Attendees: []*User{users["adent"]},
	})
	users["fprefect"].AddAppointment(&Appointment{
		Subject:   "Get Thrown into Space",
		StartTime: parseTime("1979-07-02 11:40"),
		Length:    5,
		Attendees: []*User{users["adent"]},
	})
	users["fprefect"].AddAppointment(&Appointment{
		Subject:   "Get Saved from Space",
		StartTime: parseTime("1979-07-02 11:45"),
		Length:    1,
		Attendees: []*User{users["adent"]},
	})
	users["zbeeblebrox"].AddAppointment(&Appointment{
		Subject:   "Explore Planet Builder's Homeworld",
		StartTime: parseTime("1979-07-03 11:00"),
		Length:    240,
		Attendees: []*User{users["adent"], users["fprefect"], users["tmacmillan"],
			users["mrobot"]},
	})

	for _, user := range users {
		db.Save(&user)
	}

}

func parseTime(timeRaw string) time.Time {
	const timeLayout = "2006-01-02 15:04"
	t, _ := time.Parse(timeLayout, timeRaw)
	return t
}

type User struct {
	gorm.Model
	Username  string
	FirstName string
	LastName  string
	Calendar  Calendar
}

func (u *User) AddAppointment(appt *Appointment) {
	u.Calendar.Appointments = append(u.Calendar.Appointments, appt)
}

type Calendar struct {
	gorm.Model
	Name         string
	UserID       uint
	Appointments []*Appointment
}

type Appointment struct {
	gorm.Model
	Subject     string
	Description string
	StartTime   time.Time
	Length      uint
	CalendarID  uint
	Attendees   []*User `gorm:"many2many:appointment_user"`
}
