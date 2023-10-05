package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectToDb() *sql.DB {
	url := "postgres://deliviry_db_user:deliviry_db_pass@db_deli_compose:5432/deliviry_db?sslmode=disable"
	db, err := sql.Open("postgres", url)
	//db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=uzumdeliv sslmode= disable")
	if err != nil {
		log.Println(err, "db connection error")
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)

	}

	return db
}
