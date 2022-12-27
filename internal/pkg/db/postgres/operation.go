package postgres

import (
	"database/sql"
)

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
