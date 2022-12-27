package db

import (
	"log"
)

var DB Store

type Store interface {
	InitStore() error
	Insert(interface{}, ...any) error
}

func InitDB(db Store) {
	err := db.InitStore()

	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	DB = db
}
