package db

// InitPostgresScript postgres

type script struct{}

func newScript() script {
	return script{}
}

func (q script) CreateExpenseTable() string {
	return `
			CREATE TABLE IF NOT EXISTS expenses (
				id SERIAL PRIMARY KEY,
				title TEXT,
				amount FLOAT,
				note TEXT,
				tags TEXT[]
			);`
}

func (q script) InsertExpense() string {
	return "INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id"
}

func (q script) GetExpense() string {
	return "SELECT id, title, amount, note, tags FROM expenses where id=$1"
}

func (q script) GetExpenses() string {
	return "SELECT id, title, amount, note, tags FROM expenses"
}

func (q script) UpdateExpense() string {
	return "UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1"
}
