package config

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = "5431"
	user     = "clean-archer"
	password = "arrow"
	dbname   = "cleanarch-db"
)

var psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func GetPgDbConnection() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Pg database connection created")

	return db
}
