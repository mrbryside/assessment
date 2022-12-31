//go:build unit

package expense

import (
	"errors"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/util/common"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	createPayload = `{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`
)

var createTests = []struct {
	name     string
	code     int
	spy      db.StoreSpy
	payload  string
	response string
	called   bool
}{
	{
		name:    "should return expense response",
		code:    http.StatusCreated,
		spy:     newSpyCreateSuccess(),
		payload: createPayload,
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
		name: "should return bad request response",
		code: http.StatusBadRequest,
		spy:  newSpyCreateSuccess(),
		payload: `{
			"title": "strawberry smoothie",
    		"amount": "12345",
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
		name: "should return bad request required field response",
		code: http.StatusBadRequest,
		spy:  newSpyCreateSuccess(),
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
		name:    "should return response internal server error",
		code:    http.StatusInternalServerError,
		spy:     newSpyCreateFail(),
		payload: createPayload,
		response: `{
			"code": "5000",
			"message": "internal server error"
		}`,
		called: true,
	},
}

func setup(t *testing.T) {
	t.Parallel()
	log.SetOutput(io.Discard)
}

func TestCreateExpense(t *testing.T) {
	setup(t)
	for _, ctc := range createTests {
		ctc := ctc
		t.Run(ctc.name, func(t *testing.T) {
			// Arrange
			expenses := NewExpense(ctc.spy)
			e := echo.New()
			e.Validator = common.Validator(validator.New())
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(ctc.payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := expenses.CreateExpenseHandler(c)

			wantResp := ctc.response
			wantCalled := ctc.called
			wantCode := ctc.code

			// Act
			gotErr := err
			gotResp := rec.Body.String()
			gotCalled := ctc.spy.IsWasCalled()
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

// --- create fail spy
func newSpyCreateFail() db.StoreSpy {
	return db.NewStoreSpy(insertFail, nil, nil, nil)
}

func insertFail(args ...any) error {
	return errors.New("can't insert")
}

// --- create success spy
func newSpyCreateSuccess() db.StoreSpy {
	return db.NewStoreSpy(insertSuccess, nil, nil, nil)
}

func insertSuccess(args ...any) error {
	modelId := args[0]
	p, _ := modelId.(*int)
	*p = 5
	return nil
}
