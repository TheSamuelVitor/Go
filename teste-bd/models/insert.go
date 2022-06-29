package models

import (
	"github.com/TheSamuelVitor/Go/tree/main/teste-bd/db"
)

func insert(todo ToDo) (id int64, err error) {

	conn, err := db.OpenConnection()

	return
}