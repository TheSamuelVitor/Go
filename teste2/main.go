package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "sysadmin"
	dbname   = "postgres"
)

func main() {
	psqlconn := fmt.Sprintf("host= %s port = %d user = %s password = %s dbname = %s sslmode = disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertStmt := `insert into "employee"("name", "empid") values('Samuel', 21)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	insertDynStmt := `insert into "employee"("name", "empid") values($1, $2)`
	_, e = db.Exec(insertDynStmt, "Miguel", 18)
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}