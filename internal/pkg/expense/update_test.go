//go:build only

package expense

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	updateResponse = `{
		"id": 0,
		"title": "",
		"amount": 0,
		"note": "",
		"tags": []
	}`
)

var updateTests = []struct {
	name     string
	code     int
	spy      db.StoreSpy
	param    string
	payload  string
	response string
	called   bool
}{
	{
		name:     "should return response expense json",
		code:     http.StatusOK,
		spy:      nil,
		param:    "1",
		payload:  updateResponse,
		response: updateResponse,
		called:   true,
	},
}

func TestUpdateExpense(t *testing.T) {
	t.Parallel()
	for _, utc := range updateTests {
		utc := utc
		t.Run(utc.name, func(t *testing.T) {
			// Arrange
			expenses := NewExpense(utc.spy)
			e := echo.New()
			e.Validator = util.Validator(validator.New())
			req := httptest.NewRequest(
				http.MethodPut,
				fmt.Sprintf("/expenses/%s", utc.param),
				strings.NewReader(utc.payload),
			)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/expenses/:id")
			c.SetParamNames("id")
			c.SetParamValues(utc.param)
			err := expenses.UpdateExpenseHandler(c)

			_ = utc.response
			wantCode := utc.code

			// Act
			gotErr := err
			_ = rec.Body.String()
			gotCode := c.Response().Status

			// Assert
			assert.Nil(t, gotErr)
			assert.Equal(t, wantCode, gotCode)
		})
	}
}
