package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/util"
)

func (e *expense) GetExpenseHandler(c echo.Context) error {
	modelDto := newModelDto()
	param := newParamDto()
	err := c.Bind(param)
	if err != nil {
		return util.JsonHandler().BadRequest(c, "Request parameter is invalid.")
	}

	err = c.Validate(param)
	if err != nil {
		return util.JsonHandler().BadRequest(c, err.Error())
	}

	_ = e.store.FindOne(1, modelDto, "SELECT id, name, age FROM users where id=$1")

	return util.JsonHandler().Success(c, modelDto)
}
