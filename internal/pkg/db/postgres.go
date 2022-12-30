package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type postgres struct {
	url string
	db  *sql.DB
}

func NewPostgres(url string) *postgres {
	return &postgres{url: url}
}

func initTable(db *sql.DB) {
	createTb := newScript().CreateExpenseTable()
	_, err := db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table ", err)
	}
}
