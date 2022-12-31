//go:build unit

package expense

import (
	"errors"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/util/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var getsTests = []struct {
	name     string
	code     int
	spy      db.StoreSpy
	response string
	called   bool
}{
	{
		name: "should return response expenses array",
		code: http.StatusOK,
		spy:  newSpyGetsSuccess(),
		response: `[{
			"id": 1,
			"title": "strawberry smoothie",
			"amount": 79,
			"note": "night market promotion discount 10 bath", 
			"tags": ["food", "beverage"]
		}]`,
		called: true,
	},
	{
		name: "should return response internal server error",
		code: http.StatusInternalServerError,
		spy:  newSpyGetsFail(),
		response: `{
			"code": "5000",
			"message": "internal server error"
		}`,
		called: true,
	},
}

func TestGetExpenses(t *testing.T) {
	t.Parallel()
	for _, gt := range getsTests {
		t.Run(gt.name, func(t *testing.T) {
			// Arrange
			expenses := NewExpense(gt.spy)
			e := echo.New()
			e.Validator = common.Validator(validator.New())
			req := httptest.NewRequest(http.MethodGet, "/expenses", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := expenses.GetExpensesHandler(c)

			wantResp := gt.response

			// Act
			gotErr := err
			gotResp := rec.Body.String()

			// Assert
			assert.Nil(t, gotErr)
			assert.JSONEq(t, wantResp, gotResp)
		})
	}
}

// --- get expenses success spy
func newSpyGetsSuccess() db.StoreSpy {
	return db.NewStoreSpy(nil, nil, findSuccess, nil)
}

func findSuccess(args ...any) ([]interface{}, error) {
	var results []interface{}
	model := newModelExpense()
	model.ID = 1
	model.Title = "strawberry smoothie"
	model.Amount = 79
	model.Note = "night market promotion discount 10 bath"
	model.Tags = []string{"food", "beverage"}

	results = append(results, model)
	return results, nil
}

// --- get expenses fail internal spy
func newSpyGetsFail() db.StoreSpy {
	return db.NewStoreSpy(nil, nil, findFail, nil)
}

func findFail(args ...any) ([]interface{}, error) {
	return nil, errors.New("error")
}
