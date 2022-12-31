package expense

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mrbryside/assessment/internal/pkg/util/common"
	"github.com/mrbryside/assessment/internal/pkg/util/errs"
	"github.com/mrbryside/assessment/internal/pkg/util/httputil"
	"strconv"
)

func (e *expense) UpdateExpenseHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return httputil.BadRequest(c, "Request path parameter is invalid.")

	}
	model := newModelExpense()

	err = c.Bind(&model)
	if err != nil {
		log.Errorf("Update expense error with invalid request parameter")
		return httputil.BadRequest(c, "Request parameters are invalid.")
	}

	err = c.Validate(model)
	if err != nil {
		log.Errorf("Update expense error with missing request parameter")
		return httputil.BadRequest(c, err.Error())
	}

	// check expense exist
	log.Info("Checking expense with ID: ", id)
	var found modelExpense
	err = e.store.FindOne(
		id,
		e.store.Script().GetExpense(),
		found.Arguments()...,
	)
	if err != nil && errs.CompareError(err, errs.Error().DBNotFound) {
		log.Error("Check expense not found, ", err)
		return httputil.NotFound(c, "expense not found")
	}
	if err != nil {
		log.Error("Check expense internal error, ", err)
		return httputil.InternalServerError(c)
	}

	// update operation
	log.Info("Updating expense with payload: ", common.JsonFormat(model))
	model.ID = id
	err = e.store.Update(e.store.Script().UpdateExpense(), model.Arguments()...)
	if err != nil {
		log.Error("Update expense internal error, ", err)
		return httputil.InternalServerError(c)
	}

	log.Info("Update expense success!")
	return httputil.Success(c, model)
}
