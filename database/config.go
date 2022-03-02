package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func MySQLConnection() *sql.DB {
	database := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=%s&loc=%s",
		os.Getenv("mysqlUser"),
		os.Getenv("mysqlPassword"),
		os.Getenv("mysqlHost"),
		os.Getenv("mysqlPort"),
		os.Getenv("mysqlName"),
		os.Getenv("mysqlParsetime"),
		os.Getenv("mysqlTimeLocation"),
	)

	db, err := sql.Open("mysql", database)
	if err != nil {
		log.Fatalf("%s", err)
		panic(err)
	}

	return db
}
