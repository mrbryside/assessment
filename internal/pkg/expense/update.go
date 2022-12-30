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

	// check expense exist
	var found modelExpense
	err = e.store.FindOne(
		id,
		e.store.Script().GetExpense(),
		found.Arguments()...,
	)
	if err != nil && util.Error().CompareError(err, util.Error().DBNotFound) {
		return util.NotFound(c, "expense not found")
	}
	if err != nil {
		return util.InternalServerError(c)
	}

	// update operation
	model.ID = id
	err = e.store.Update(e.store.Script().UpdateExpense(), model.Arguments()...)
	if err != nil {
		return util.InternalServerError(c)
	}

	return util.Success(c, model)
}
