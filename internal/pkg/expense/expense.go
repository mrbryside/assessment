package expense

import (
	"github.com/lib/pq"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/util/common"
)

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

func (m *modelExpense) Arguments() []interface{} {
	return common.Arguments(
		&m.ID,
		&m.Title,
		&m.Amount,
		&m.Note,
		pq.Array(&m.Tags),
	)
}

type paramDto struct {
	ID int `json:"id" param:"id" validate:"required"`
}

func newParamDto() *paramDto {
	return &paramDto{}
}
