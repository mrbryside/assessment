//go:build integration

package expense

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/middleware"
	"github.com/mrbryside/assessment/internal/pkg/util/httputil"
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
	database, err := db.InitDB(db.NewPostgres("postgresql://root:root@db/test-db?sslmode=disable"))
	if err != nil {
		t.Fatal("can't init db")
	}
	defer database.Close()
	expenses := NewExpense(db.DB)
	eh := echo.New()
	eh.Use(middleware.VerifyAuthorization)
	eh = httputil.InitItEcho(eh, func() {
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
	err = httputil.Seeder(&created, getBody, "expenses")
	if err != nil {
		t.Fatal("can't create expense:", err)
	}

	res := httputil.Request(http.MethodGet, httputil.Uri("expenses", strconv.Itoa(created.ID)), nil)
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
