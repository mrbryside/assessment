package db

import (
	"database/sql"
	"github.com/mrbryside/assessment/internal/pkg/util"
)

type Store interface {
	InitStore() error
	Insert(interface{}, ...any) error
	FindOne(int, string, ...any) error
}

func (p *postgres) InitStore() error {
	pDb, err := sql.Open("postgres", p.url)
	if err != nil {
		return err
	}
	p.db = pDb

	// initial table
	initTable(pDb)

	return nil
}

func (p *postgres) Insert(modelId interface{}, args ...any) error {
	// destructuring args
	queryLang := args[0].(string)
	otherArgs := args[1:]

	// insert entity
	row := p.db.QueryRow(queryLang, otherArgs...)

	err := row.Scan(modelId)
	if err != nil {
		return err
	}
	return nil
}
func (p *postgres) FindOne(rowId int, queryLang string, args ...any) error {
	stmt, err := p.db.Prepare(queryLang)
	if err != nil {
		return err
	}

	row := stmt.QueryRow(rowId)
	err = row.Scan(args...)
	if err != nil && err == sql.ErrNoRows {
		return util.Error().DBNotFound
	}
	if err != nil {
		return err
	}

	return nil
}
