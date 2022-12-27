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
func (p *postgres) FindOne(rowId int, model interface{}, queryLang string) error {
	// destructuring args
	//"SELECT id, name, age FROM users where id=$1"
	stmt, err := p.db.Prepare(queryLang)
	if err != nil {

	}

	row := stmt.QueryRow(rowId)
	err = row.Scan(model)
	if err != nil {
		return err
	}
	return nil
}
