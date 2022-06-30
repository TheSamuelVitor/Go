package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB
var err error

type sandbox struct {
	id int
	firstName string
	lastName string
	Age int
}

func init()  {
	connStr := "postgres://postgres:sysadmin@localhost/postgres?sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to postgres")

}

func main() {
	
}