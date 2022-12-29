package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/mrbryside/assessment/internal/pkg/util"
)

func (e *expense) GetExpenseHandler(c echo.Context) error {
	model := newModelExpense()
	param := newParamDto()

	err := c.Bind(param)
	if err != nil {
		return util.JsonHandler().BadRequest(c, "Request parameter is invalid.")
	}

	err = c.Validate(param)
	if err != nil {
		return util.JsonHandler().BadRequest(c, err.Error())
	}

	err = e.store.FindOne(
		param.ID,
		e.store.Script().GetExpense(),
		&model.ID,
		&model.Title,
		&model.Amount,
		&model.Note,
		pq.Array(&model.Tags),
	)
	if err != nil && util.Error().CompareError(err, util.Error().DBNotFound) {
		return util.JsonHandler().NotFound(c, "expense not found")
	}
	if err != nil {
		return util.JsonHandler().InternalServerError(c)
	}

	return util.JsonHandler().Success(c, model)
}
