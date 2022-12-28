package expense

import "github.com/mrbryside/assessment/internal/pkg/db"

type modelDto struct {
	ID     int      `json:"id"`
	Title  string   `json:"title" validate:"required"`
	Amount int      `json:"amount" validate:"required"`
	Note   string   `json:"note" validate:"required"`
	Tags   []string `json:"tags" param:"tags" validate:"required,dive,required"`
}

func newModelDto() *modelDto {
	return &modelDto{}
}

type paramDto struct {
	ID int `json:"id" param:"id" validate:"required"`
}

func newParamDto() *paramDto {
	return &paramDto{}
}

// expense
type expense struct {
	store db.Store
}

func NewExpense(store db.Store) *expense {
	return &expense{store: store}
}
