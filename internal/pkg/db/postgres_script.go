package db

// InitPostgresScript postgres

type queryScript struct{}

func Script() queryScript {
	return queryScript{}
}

func (q queryScript) CreateExpenseTable() string {
	return `
			CREATE TABLE IF NOT EXISTS expenses (
				id SERIAL PRIMARY KEY,
				title TEXT,
				amount FLOAT,
				note TEXT,
				tags TEXT[]
			);`
}

func (q queryScript) InsertExpense() string {
	return "INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id"
}

func (q queryScript) GetExpense() string {
	return "SELECT id, title, amount, note, tags FROM expenses where id=$1"
}
