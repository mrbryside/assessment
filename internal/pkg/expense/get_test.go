//go:build unit

package expense

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/expense/mock"
	"github.com/mrbryside/assessment/internal/pkg/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var getTests = []struct {
	name string
	mock mock.GetExpenseMock
}{
	{name: "should return response expense data", mock: mock.GetterMock().GetExpenseSuccess()},
	{name: "should return required path params", mock: mock.GetterMock().GetExpenseParamsFailed()},
}

func TestGetExpense(t *testing.T) {
	for _, gtc := range getTests {
		gtc := gtc
		t.Run(gtc.name, func(t *testing.T) {
			// Arrange
			m := gtc.mock
			expenses := NewExpense(m.SpyStore)
			e := echo.New()
			e.Validator = util.Validator(validator.New())
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/expenses/%s", m.Payload), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/expenses/:id")
			c.SetParamNames("id")
			c.SetParamValues(m.Payload)
			err := expenses.GetExpenseHandler(c)

			wantResp := m.Response
			wantCode := m.Code

			// Act
			gotResp := rec.Body.String()
			gotStatus := c.Response().Status
			gotError := err

			// Assert
			assert.Nil(t, gotError)
			assert.JSONEq(t, gotResp, wantResp)
			assert.Equal(t, gotStatus, wantCode)
		})
	}
}
