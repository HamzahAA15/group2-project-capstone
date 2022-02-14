package database

import (
	"database/sql"
	"fmt"
	"log"
)

func MySQLConnection(database string) *sql.DB {
	fmt.Println("Connecting to MySQL DB")

	db, err := sql.Open("mysql", database)
	if err != nil {
		log.Fatalf("%s", err)
		panic(err)
	}
	return db
}
