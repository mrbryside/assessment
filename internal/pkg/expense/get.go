package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/util"
)

func (e *expense) GetExpenseHandler(c echo.Context) error {
	model := newModelExpense()
	param := newParamDto()

	err := c.Bind(param)
	if err != nil {
		return util.BadRequest(c, "Request parameter is invalid.")
	}

	err = c.Validate(param)
	if err != nil {
		return util.BadRequest(c, err.Error())
	}

	err = e.store.FindOne(
		param.ID,
		e.store.Script().GetExpense(),
		model.Arguments()...,
	)
	if err != nil && util.Error().CompareError(err, util.Error().DBNotFound) {
		return util.NotFound(c, "expense not found")
	}
	if err != nil {
		return util.InternalServerError(c)
	}

	return util.Success(c, model)
}
