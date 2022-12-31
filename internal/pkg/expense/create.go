package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mrbryside/assessment/internal/pkg/util/httputil"
)

func (e *expense) CreateExpenseHandler(c echo.Context) error {
	model := newModelExpense()

	err := c.Bind(model)
	if err != nil {
		log.Errorf("Creating expense error with invalid request parameter")
		return httputil.BadRequest(c, "Request parameters are invalid.")
	}

	err = c.Validate(model)
	if err != nil {
		log.Errorf("Creating expense error with missing request parameter")
		return httputil.BadRequest(c, err.Error())
	}

	err = e.store.Insert(
		e.store.Script().InsertExpense(),
		model.Arguments()...,
	)
	if err != nil {
		return httputil.InternalServerError(c)
	}

	return httputil.SuccessCreated(c, model)
}
