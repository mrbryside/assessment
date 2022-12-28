package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/util"
)

func (e *expense) GetExpenseHandler(c echo.Context) error {
	model := newModelDto()
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
		db.Script().GetExpense(),
		&model.ID,
		&model.Title,
		&model.Amount,
		&model.Note,
		pq.Array(&model.Tags),
	)
	if err != nil {
		return util.JsonHandler().InternalServerError(c)
	}

	return util.JsonHandler().Success(c, model)
}
