package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/util"
)

func (e *expense) GetExpensesHandler(c echo.Context) error {
	model := newModelExpense()
	// get expenses
	results, _ := e.store.Find(e.store.Script().GetExpenses(), model, model.Arguments()...)
	return util.Success(c, results)
}
