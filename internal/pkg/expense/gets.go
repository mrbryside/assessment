package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/util/httputil"
)

func (e *expense) GetExpensesHandler(c echo.Context) error {
	model := newModelExpense()

	results, err := e.store.Find(e.store.Script().GetExpenses(), model, model.Arguments()...)
	if err != nil {
		return httputil.InternalServerError(c)
	}
	return httputil.Success(c, results)
}
