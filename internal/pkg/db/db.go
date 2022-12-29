package db

import (
	"log"
)

type Store interface {
	InitStore() error
	Script() script
	Insert(script string, args ...interface{}) error
	FindOne(int, string, ...interface{}) error
}

var DB Store

func InitDB(db Store) {
	err := db.InitStore()

	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	DB = db
}
