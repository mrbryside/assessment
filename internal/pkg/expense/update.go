package expense

import (
	"github.com/labstack/echo/v4"
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
		return httputil.BadRequest(c, "Request parameters are invalid.")
	}

	err = c.Validate(model)
	if err != nil {
		return httputil.BadRequest(c, err.Error())
	}

	// check expense exist
	var found modelExpense
	err = e.store.FindOne(
		id,
		e.store.Script().GetExpense(),
		found.Arguments()...,
	)
	if err != nil && errs.CompareError(err, errs.Error().DBNotFound) {
		return httputil.NotFound(c, "expense not found")
	}
	if err != nil {
		return httputil.InternalServerError(c)
	}

	// update operation
	model.ID = id
	err = e.store.Update(e.store.Script().UpdateExpense(), model.Arguments()...)
	if err != nil {
		return httputil.InternalServerError(c)
	}

	return httputil.Success(c, model)
}
