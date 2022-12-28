package db

// InitPostgresScript postgres

type QueryScript struct{}

func Script() QueryScript {
	return QueryScript{}
}

func (q QueryScript) CreateExpenseTable() string {
	return `
			CREATE TABLE IF NOT EXISTS expenses (
				id SERIAL PRIMARY KEY,
				title TEXT,
				amount FLOAT,
				note TEXT,
				tags TEXT[]
			);`
}

func (q QueryScript) InsertExpense() string {
	return "INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id"
}
