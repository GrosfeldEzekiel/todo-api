package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var con *sql.DB


func Connect () *sql.DB {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	return db
}
