package DB

import (
	"database/sql"
	"fmt"

	"github.com/TheSamuelVitor/Go/tree/main/teste-bd/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	// string connection
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s DBname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	// connection
	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
