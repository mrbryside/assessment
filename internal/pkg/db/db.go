package db

import (
	"database/sql"
)

type Store interface {
	InitStore() error
	Script() script
	Insert(string, ...interface{}) error
	FindOne(int, string, ...interface{}) error
	Find(script string, model interface{}, args ...interface{}) ([]interface{}, error)
	Update(string, ...interface{}) error
}

var DB Store
var realDB *sql.DB

func InitDB(db Store) (*sql.DB, error) {
	err := db.InitStore()

	if err != nil {
		return nil, err
	}
	DB = db
	return realDB, nil
}
