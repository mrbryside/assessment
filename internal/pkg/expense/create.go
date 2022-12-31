package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mrbryside/assessment/internal/pkg/util/common"
	"github.com/mrbryside/assessment/internal/pkg/util/httputil"
)

func (e *expense) CreateExpenseHandler(c echo.Context) error {
	model := newModelExpense()

	err := c.Bind(model)
	if err != nil {
		log.Error("Create expense error with invalid request parameter, ", err)
		return httputil.BadRequest(c, "Request parameters are invalid.")
	}

	err = c.Validate(model)
	if err != nil {
		log.Error("Create expense error with missing request parameter, ", err)
		return httputil.BadRequest(c, err.Error())
	}

	log.Info("Creating expense with payload: ", common.JsonFormat(model))
	err = e.store.Insert(
		e.store.Script().InsertExpense(),
		model.Arguments()...,
	)
	if err != nil {
		log.Error("Create expense internal error, ", err)
		return httputil.InternalServerError(c)
	}

	log.Info("Create expense success!")
	return httputil.SuccessCreated(c, model)
}
