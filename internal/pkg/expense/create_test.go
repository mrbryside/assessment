//go:build unit

package expense

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/expense/mock"
	"github.com/mrbryside/assessment/internal/pkg/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var createTests = []struct {
	name string
	mock mock.CreateExpenseMock
}{
	{name: "should return response expense data", mock: mock.CreationMock().CreateSuccess()},
	{name: "should return response bad request invalid", mock: mock.CreationMock().CreateBindFail()},
	{name: "should return response bad request required field", mock: mock.CreationMock().CreateValidateFail()},
	{name: "should return response internal server error", mock: mock.CreationMock().CreateInternalFail()},
}

func TestCreateExpense(t *testing.T) {
	for _, tc := range createTests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			m := tc.mock
			expenses := NewExpense(m.SpyStore)
			e := echo.New()
			e.Validator = util.Validator(validator.New())
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(m.Payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := expenses.CreateExpenseHandler(c)

			wantResp := m.Response
			wantCalled := m.Called
			wantCode := m.Code

			// Act
			gotErr := err
			gotResp := rec.Body.String()
			gotCalled := m.SpyStore.IsWasCalled()
			gotCode := c.Response().Status

			// Assert
			assert.Nil(t, gotErr)
			assert.Equal(t, wantCode, gotCode)
			assert.JSONEq(t, wantResp, gotResp)
			assert.Equal(t, wantCode, gotCode)
			assert.Equal(t, wantCalled, gotCalled)

		})
	}
}
