//go:build unit

package expense

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var getTests = []struct {
	name     string
	code     int
	spy      spyStore
	payload  string
	response string
	called   bool
}{
	{
		name:    "should return response expense json",
		code:    http.StatusOK,
		spy:     newSpyStoreWithGetExpenseSuccess(),
		payload: "5",
		response: `{
			"id": 5,
			"title": "strawberry smoothie",
    		"amount": 79,
    		"note": "night market promotion discount 10 bath",
    		"tags": ["food", "beverage"]
		}`,
		called: true,
	},
	{
		name:    "should return response required path params",
		code:    http.StatusBadRequest,
		spy:     newSpyStoreWithGetExpenseSuccess(),
		payload: "",
		response: `{
			"code": "4000",
			"message": "ID is a required field"
		}`,
		called: false,
	},
	{
		name:    "should return response invalid path params",
		code:    http.StatusBadRequest,
		spy:     newSpyStoreWithGetExpenseSuccess(),
		payload: "error",
		response: `{
			"code": "4000",
			"message": "Request parameter is invalid."
		}`,
		called: false,
	},
	{
		name:    "should return internal server error",
		code:    http.StatusInternalServerError,
		spy:     newSpyStoreWithGetExpenseFail(),
		payload: "5",
		response: `{
			"code": "5000",
			"message": "internal server error"
		}`,
		called: true,
	},
	{
		name:    "should return response expense not found",
		code:    http.StatusNotFound,
		spy:     newSpyStoreWithGetExpenseNotFound(),
		payload: "5",
		response: `{
			"code": "4004",
			"message": "expense not found"
		}`,
		called: true,
	},
}

func TestGetExpense(t *testing.T) {
	t.Parallel()
	for _, gtc := range getTests {
		t.Run(gtc.name, func(t *testing.T) {
			// Arrange
			expenses := NewExpense(gtc.spy)
			e := echo.New()
			e.Validator = util.Validator(validator.New())
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/expenses/%s", gtc.payload), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/expenses/:id")
			c.SetParamNames("id")
			c.SetParamValues(gtc.payload)
			err := expenses.GetExpenseHandler(c)

			wantResp := gtc.response
			wantCode := gtc.code
			wantCalled := gtc.called

			// Act
			gotResp := rec.Body.String()
			gotStatus := c.Response().Status
			gotError := err
			gotCalled := gtc.spy.IsWasCalled()

			// Assert
			assert.Nil(t, gotError)
			assert.JSONEq(t, wantResp, gotResp)
			assert.Equal(t, wantCalled, gotCalled)
			assert.Equal(t, wantCode, gotStatus)
		})
	}
}
