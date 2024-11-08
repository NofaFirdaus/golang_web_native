package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDb() {
	db, err := sql.Open("mysql", "root:nofafirdausananta@/golang")
	if err != nil {
		panic(err)
	}
	log.Println("databases connect")
	DB = db
}
