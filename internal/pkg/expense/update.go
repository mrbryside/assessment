package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/util"
	"strconv"
)

func (e *expense) UpdateExpenseHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	model := newModelExpense()

	err := c.Bind(&model)
	if err != nil {
		return util.BadRequest(c, "Request parameters are invalid.")
	}

	err = c.Validate(model)
	if err != nil {
		return util.BadRequest(c, err.Error())
	}

	err = e.store.FindOne(
		id,
		e.store.Script().GetExpense(),
		model.Arguments()...,
	)
	if err != nil && util.Error().CompareError(err, util.Error().DBNotFound) {
		return util.NotFound(c, "expense not found")
	}
	if err != nil {
		return util.InternalServerError(c)
	}

	model.ID = id

	_ = e.store.Update(e.store.Script().UpdateExpense(), model.Arguments()...)

	return util.Success(c, model)
}
