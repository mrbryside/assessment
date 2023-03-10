//go:build integration

package expense

import (
	"bytes"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/middleware"
	"github.com/mrbryside/assessment/internal/pkg/util/httputil"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

const (
	createdBody = `{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath",
		"tags": ["food", "beverage"]
	}`
	createdCode = http.StatusCreated
)

func TestIntegrationCreateExpense(t *testing.T) {
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
	})

	// Arrange
	payload := createdBody
	wantCode := createdCode
	wantTitle := "strawberry smoothie"
	wantAmount := 79
	wantNote := "night market promotion discount 10 bath"
	wantTags := []string{"food", "beverage"}

	// Act
	var e modelExpense
	body := bytes.NewBufferString(payload)
	res := httputil.Request(http.MethodPost, httputil.Uri("expenses"), body)
	err = res.Decode(&e)

	gotErr := err
	gotTitle := e.Title
	gotAmount := e.Amount
	gotNote := e.Note
	gotTags := e.Tags
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
