package expense

import "github.com/mrbryside/assessment/internal/pkg/db"

type expense struct {
	store db.Store
}

func NewExpense(store db.Store) *expense {
	return &expense{store: store}
}

// model
type modelExpense struct {
	ID     int      `json:"id"`
	Title  string   `json:"title" validate:"required"`
	Amount int      `json:"amount" validate:"required"`
	Note   string   `json:"note" validate:"required"`
	Tags   []string `json:"tags" param:"tags" validate:"required,dive,required"`
}

func newModelExpense() *modelExpense {
	return &modelExpense{}
}

type paramDto struct {
	ID int `json:"id" param:"id" validate:"required"`
}

func newParamDto() *paramDto {
	return &paramDto{}
}
