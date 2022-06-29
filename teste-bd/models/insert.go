package models

import (
"https://github.com/TheSamuelVitor/Go/tree/main/teste-bd/db"
)

func insert(todo ToDo) (id int64, err error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	sql := `insert into todos (title, description, done) values($1,$2,$3) returning id`

	err = conn.QueryRow(sql, todo.title, todo.description, todo.done).Scan(&id)

	return
}