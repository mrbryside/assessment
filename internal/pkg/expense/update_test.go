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
		name:  "should return response updated expense json",
		code:  http.StatusOK,
		spy:   newSpyUpdateSuccess(),
		param: "1",
		payload: `{
			"title": "strawberry smoothie",
			"amount": 79,
			"note": "night market promotion discount 10 bath", 
			"tags": ["food", "beverage"]
		}`,
		response: `{
			"id": 1,
			"title": "strawberry smoothie",
			"amount": 79,
			"note": "night market promotion discount 10 bath", 
			"tags": ["food", "beverage"]
		}`,
		called: true,
	},
	{
		name:  "should return bad request response",
		code:  http.StatusBadRequest,
		spy:   newSpyUpdateSuccess(),
		param: "1",
		payload: `{
			"title": "strawberry smoothie",
			"amount": "1234",
			"note": "night market promotion discount 10 bath", 
			"tags": ["food", "beverage"]
		}`,
		response: `{
			"code": "4000",
			"message": "Request parameters are invalid."
		}`,
		called: false,
	},
	{
		name:  "should return bad request validation",
		code:  http.StatusBadRequest,
		spy:   newSpyUpdateSuccess(),
		param: "1",
		payload: `{
			"title": "strawberry smoothie",
			"note": "night market promotion discount 10 bath", 
			"tags": ["food", "beverage"]
		}`,
		response: `{
			"code": "4000",
			"message": "Amount is a required field"
		}`,
		called: false,
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

			wantResp := utc.response
			wantCode := utc.code
			wantCalled := utc.called

			// Act
			gotErr := err
			gotResp := rec.Body.String()
			gotCode := c.Response().Status
			gotCalled := utc.spy.IsWasCalled()

			// Assert
			assert.Nil(t, gotErr)
			assert.JSONEq(t, wantResp, gotResp)
			assert.Equal(t, wantCode, gotCode)
			assert.Equal(t, wantCalled, gotCalled)
		})
	}
}

// --- update success spy
func newSpyUpdateSuccess() db.StoreSpy {
	return db.NewStoreSpy(nil, nil, updateSuccess)
}

func updateSuccess(args ...any) error {
	return nil
}
