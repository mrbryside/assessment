package expense

import "github.com/mrbryside/assessment/internal/pkg/db"

type ModelDto struct {
	ID     int      `json:"id"`
	Title  string   `json:"title" validate:"required"`
	Amount int      `json:"amount" validate:"required"`
	Note   string   `json:"note" validate:"required"`
	Tags   []string `json:"tags" param:"tags" validate:"required,dive,required"`
}

func newModelDto() *ModelDto {
	return &ModelDto{}
}

func newExpenseDto(id int, title string, amount int, note string, tags []string) *ModelDto {
	return &ModelDto{
		ID:     id,
		Title:  title,
		Amount: amount,
		Note:   note,
		Tags:   tags,
	}
}

// paramDto
type paramDto struct {
	Id int `json:"id" param:"id" validate:"required"`
}

// expense
type expense struct {
	store db.Store
}

func NewExpense(store db.Store) *expense {
	return &expense{store: store}
}
