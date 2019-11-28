package main

import (
	"flag"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	dsn := flag.String("dsn", "postgres://ivorscott:@localhost:5432/gorm-usage?sslmode=disable", "Postgres data source name")
	flag.Parse()

	// Connect to postgres in two ways, via:

	// 1. url connection string
	// db, err := gorm.Open("postgres", "postgres://ivorscott:@localhost:5432/gorm-usage?sslmode=disable")

	// 2. or non-url connection string
	// db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=gorm-usage user=ivorscott password='' sslmode=disable")

	db, err := gorm.Open("postgres", *dsn)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	dbase := db.DB()
	defer dbase.Close()

	err = dbase.Ping()
	if err != nil {
		panic(err.Error())
	}

	println("Connection to database established")
}
