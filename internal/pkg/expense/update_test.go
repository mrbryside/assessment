//go:build only

package expense

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
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
		"id": 1,
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath",
		"tags": ["food", "beverage"]
	}`
	updatePayload = `{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
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
		name:     "should return response updated expense json",
		code:     http.StatusOK,
		spy:      newSpyUpdateSuccess(),
		param:    "1",
		payload:  updatePayload,
		response: updateResponse,
		called:   true,
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
	{
		name:    "should return response expense not found",
		code:    http.StatusNotFound,
		spy:     newSpyCheckExistNotFound(),
		payload: updatePayload,
		response: `{
			"code": "4004",
			"message": "expense not found"
		}`,
		called: true,
	},
	{
		name:    "should return internal server error from check expense",
		code:    http.StatusInternalServerError,
		spy:     newSpyCheckExistFail(),
		param:   "1",
		payload: updatePayload,
		response: `{
			"code": "5000",
			"message": "internal server error"
		}`,
		called: true,
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
	return db.NewStoreSpy(nil, findOneCheckSuccess, updateSuccess)
}

func updateSuccess(args ...any) error {
	return nil
}

func findOneCheckSuccess(args ...any) error {
	var model modelExpense
	err := json.Unmarshal([]byte(updateResponse), &model)
	if err != nil {
		return err
	}
	id, _ := args[0].(*int)
	*id = model.ID

	title, _ := args[1].(*string)
	*title = model.Title

	amount, _ := args[2].(*int)
	*amount = model.Amount

	note, _ := args[3].(*string)
	*note = model.Note

	tags, _ := args[4].(*pq.StringArray)
	*tags = model.Tags

	return nil
}

// --- get fail spy
func newSpyCheckExistFail() db.StoreSpy {
	return db.NewStoreSpy(nil, findOneCheckFail, nil)
}

func findOneCheckFail(args ...any) error {
	return errors.New("error")
}

// --- get not found spy
func newSpyCheckExistNotFound() db.StoreSpy {
	return db.NewStoreSpy(nil, findOneCheckNotFound, nil)
}

func findOneCheckNotFound(args ...any) error {
	return util.Error().DBNotFound
}
