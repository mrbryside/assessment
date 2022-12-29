package db

import (
	"log"
)

var DB Store

func InitDB(db Store) {
	err := db.InitStore()

	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	DB = db
}
