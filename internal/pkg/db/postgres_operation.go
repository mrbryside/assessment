package db

import (
	"database/sql"
	"github.com/mrbryside/assessment/internal/pkg/util"
)

func (p *postgres) InitStore() error {
	pDb, err := sql.Open("postgres", p.url)
	if err != nil {
		return err
	}
	p.db = pDb
	realDB = pDb

	// initial table
	err = initTable(pDb)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgres) Script() script {
	return newScript()
}
func (p *postgres) Insert(script string, args ...interface{}) error {
	// initial argument from model without ID (args index 0)
	modelId := args[0]
	otherArgs := args[1:]

	// insert entity
	row := p.db.QueryRow(script, otherArgs...)
	err := row.Scan(modelId)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgres) FindOne(rowId int, script string, args ...interface{}) error {
	stmt, err := p.db.Prepare(script)
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

func (p *postgres) Find(script string, model interface{}, args ...interface{}) ([]interface{}, error) {
	// initial argument from model without ID (args index 0)
	results := make([]interface{}, 0)
	stmt, err := p.db.Prepare(script)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(args...)
		if err != nil {
			return nil, err
		}
		results = append(results, util.Value(model))
	}
	return results, nil

}

func (p *postgres) Update(script string, args ...interface{}) error {

	stmt, err := p.db.Prepare(script)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}
