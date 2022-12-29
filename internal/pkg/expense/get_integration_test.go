//go:build integration

package expense

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
	"time"
)

const (
	getBody = `{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath",
		"tags": ["food", "beverage"]
	}`
	getCode = http.StatusOK
)

var (
	created, latest modelExpense
)

func TestIntegrationGetExpense(t *testing.T) {
	// Setup server
	db.InitDB(db.NewPostgres("postgresql://root:root@db/test-db?sslmode=disable"))
	expenses := NewExpense(db.DB)
	eh := echo.New()
	eh = util.TestHelper().InitItEcho(eh, func() {
		eh.POST("/expenses", expenses.CreateExpenseHandler)
		eh.GET("/expenses/:id", expenses.GetExpenseHandler)
	})

	// Arrange
	wantCode := getCode
	wantTitle := "strawberry smoothie"
	wantAmount := 79
	wantNote := "night market promotion discount 10 bath"
	wantTags := []string{"food", "beverage"}

	// Act
	th := util.TestHelper()
	err := th.Seeder(&created, getBody, "expenses")
	if err != nil {
		t.Fatal("can't create expense:", err)
	}

	res := th.Request(http.MethodGet, th.Uri("expenses", strconv.Itoa(created.ID)), nil)
	err = res.Decode(&latest)

	gotErr := err
	gotTitle := latest.Title
	gotAmount := latest.Amount
	gotNote := latest.Note
	gotTags := latest.Tags
	gotCode := res.StatusCode

	// Assert
	assert.Nil(t, gotErr)
	assert.Equal(t, wantCode, gotCode)
	assert.Equal(t, wantTitle, gotTitle)
	assert.Equal(t, wantAmount, gotAmount)
	assert.Equal(t, wantNote, gotNote)
	assert.Equal(t, wantTags, gotTags)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
}
