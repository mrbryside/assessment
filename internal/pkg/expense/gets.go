package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mrbryside/assessment/internal/pkg/util/httputil"
)

func (e *expense) GetExpensesHandler(c echo.Context) error {
	model := newModelExpense()

	log.Info("Getting expenses list")
	results, err := e.store.Find(e.store.Script().GetExpenses(), model, model.Arguments()...)
	if err != nil {
		log.Error("Get expenses list internal error, ", err)
		return httputil.InternalServerError(c)
	}
	log.Info("Get expenses list success!")
	return httputil.Success(c, results)
}
