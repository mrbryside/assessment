package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/util"
)

func (e *expense) GetExpenseHandler(c echo.Context) error {
	modelDto := newModelDto()
	_ = c.Param("id")

	_ = e.store.FindOne(1, modelDto, "SELECT id, name, age FROM users where id=$1")

	return util.JsonHandler().Success(c, modelDto)
}
