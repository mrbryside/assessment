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

	_ = c.Bind(&model)

	model.ID = id

	return util.JsonHandler().Success(c, model)
}
