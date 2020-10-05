package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB


func Connect () *sql.DB {
	db, err := sql.Open("mysql", "JYalzdGUu9:uBRpLtlFL0@tcp(remotemysql.com:3306)/JYalzdGUu9")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	return db
}
