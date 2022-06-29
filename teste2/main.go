package main

import "fmt"

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "sysadmin"
	dbname = "postgres"
)

func main() {
	psqlconn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)
	fmt.Printf(psqlconn)
}