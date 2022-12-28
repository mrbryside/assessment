package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
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

	_ = e.store.FindOne(
		param.ID,
		"SELECT id, title, amount, note, tags FROM expenses where id=$1",
		&model.ID,
		&model.Title,
		&model.Amount,
		&model.Note,
		pq.Array(&model.Tags),
	)

	return util.JsonHandler().Success(c, model)
}
