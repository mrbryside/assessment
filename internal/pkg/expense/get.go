package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mrbryside/assessment/internal/pkg/util/errs"
	"github.com/mrbryside/assessment/internal/pkg/util/httputil"
)

func (e *expense) GetExpenseHandler(c echo.Context) error {
	model := newModelExpense()
	param := newParamDto()

	err := c.Bind(param)
	if err != nil {
		log.Errorf("Getting expense error with invalid path parameter")
		return httputil.BadRequest(c, "Request parameter is invalid.")
	}

	err = c.Validate(param)
	if err != nil {
		log.Errorf("Getting expense error with missing path parameter")
		return httputil.BadRequest(c, err.Error())
	}

	log.Info("Getting expense with ID: ", param.ID)
	err = e.store.FindOne(
		param.ID,
		e.store.Script().GetExpense(),
		model.Arguments()...,
	)
	if err != nil && errs.CompareError(err, errs.Error().DBNotFound) {
		log.Errorf("Getting expense not found, ", err.Error())
		return httputil.NotFound(c, "expense not found")
	}
	if err != nil {
		log.Errorf("Getting expense internal error, ", err.Error())
		return httputil.InternalServerError(c)
	}
	log.Info("Get expense success!!")
	return httputil.Success(c, model)
}
