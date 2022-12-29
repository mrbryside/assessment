package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/util"
)

func (e *expense) UpdateExpenseHandler(c echo.Context) error {
	return util.JsonHandler().Success(c, nil)
}
