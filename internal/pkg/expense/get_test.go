//go:build unit

package expense

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/util/common"
	"github.com/mrbryside/assessment/internal/pkg/util/errs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	getResponse = `{
		"id": 5,
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath",
		"tags": ["food", "beverage"]
	}`
)

var getTests = []struct {
	name     string
	code     int
	spy      db.StoreSpy
	payload  string
	response string
	called   bool
}{
	{
		name:     "should return response expense json",
		code:     http.StatusOK,
		spy:      newSpyGetSuccess(),
		payload:  "5",
		response: getResponse,
		called:   true,
	},
	{
		name:    "should return response required path params",
		code:    http.StatusBadRequest,
		spy:     newSpyGetSuccess(),
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
		spy:     newSpyGetSuccess(),
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
		spy:     newSpyGetFail(),
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
		spy:     newSpyGetNotFound(),
		payload: "5",
		response: `{
			"code": "4004",
			"message": "expense not found"
		}`,
		called: true,
	},
}

func TestGetExpense(t *testing.T) {
	setup(t)
	for _, gtc := range getTests {
		t.Run(gtc.name, func(t *testing.T) {
			// Arrange
			expenses := NewExpense(gtc.spy)
			e := echo.New()
			e.Validator = common.Validator(validator.New())
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

// --- get fail spy
func newSpyGetFail() db.StoreSpy {
	return db.NewStoreSpy(nil, findOneFail, nil, nil)
}

func findOneFail(args ...any) error {
	return errors.New("error")
}

// --- get not found spy
func newSpyGetNotFound() db.StoreSpy {
	return db.NewStoreSpy(nil, findOneNotFound, nil, nil)
}

func findOneNotFound(args ...any) error {
	return errs.Error().DBNotFound
}

// --- get success spy
func newSpyGetSuccess() db.StoreSpy {
	return db.NewStoreSpy(nil, findOneSuccess, nil, nil)
}

func findOneSuccess(args ...any) error {
	var model modelExpense
	_ = json.Unmarshal([]byte(getResponse), &model)
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
