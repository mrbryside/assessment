package util

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// custom status code
const badRequestCode = "4000"
const internalErrorCode = "5000"

// custom message
const internalErrorMessage = "internal server error"

type jsonHandler struct{}

func newJsonHandler() jsonHandler {
	return jsonHandler{}
}

func (j jsonHandler) BadRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, newResponse().ApiError(badRequestCode, message))
}

func (j jsonHandler) InternalServerError(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, newResponse().ApiError(internalErrorCode, internalErrorMessage))
}

func (j jsonHandler) SuccessCreated(c echo.Context, response interface{}) error {
	return c.JSON(http.StatusCreated, response)
}

func (j jsonHandler) Success(c echo.Context, response interface{}) error {
	return c.JSON(http.StatusOK, response)
}
