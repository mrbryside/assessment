package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/mrbryside/assessment/internal/pkg/util"
)

func (e *expense) CreateExpenseHandler(c echo.Context) error {
	model := newModelExpense()

	err := c.Bind(model)
	if err != nil {
		return util.JsonHandler().BadRequest(c, "Request parameters are invalid.")
	}
	err = c.Validate(model)
	if err != nil {
		return util.JsonHandler().BadRequest(c, err.Error())
	}

	err = e.store.Insert(
		&model.ID,
		e.store.Script().InsertExpense(),
		model.Title,
		model.Amount,
		model.Note,
		pq.Array(model.Tags),
	)

	if err != nil {
		return util.JsonHandler().InternalServerError(c)
	}

	return util.JsonHandler().SuccessCreated(c, model)
}
