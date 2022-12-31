//go:build integration

package expense

import (
	"bytes"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/util/httputil"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
	"time"
)

const (
	createBody = `{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath",
		"tags": ["update test","beverage"]
	}`
	updatedBody = `{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath",
		"tags": ["beverage"]
	}`
	updateCode = http.StatusOK
)

var (
	createdExpense, updatedExpense modelExpense
)

func TestIntegrationUpdateExpense(t *testing.T) {
	// Setup server
	database, err := db.InitDB(db.NewPostgres("postgresql://root:root@db/test-db?sslmode=disable"))
	if err != nil {
		t.Fatal("can't init db")
	}
	defer database.Close()
	expenses := NewExpense(db.DB)
	eh := echo.New()
	eh = httputil.InitItEcho(eh, func() {
		eh.POST("/expenses", expenses.CreateExpenseHandler)
		eh.PUT("/expenses/:id", expenses.UpdateExpenseHandler)
	})

	// Arrange
	wantCode := updateCode
	wantTitle := "strawberry smoothie"
	wantAmount := 79
	wantNote := "night market promotion discount 10 bath"
	wantTags := []string{"beverage"}

	// Act
	err = httputil.Seeder(&createdExpense, createBody, "expenses")
	if err != nil {
		t.Fatal("can't create expense:", err)
	}
	body := bytes.NewBufferString(updatedBody)
	res := httputil.Request(http.MethodPut, httputil.Uri("expenses", strconv.Itoa(createdExpense.ID)), body)
	err = res.Decode(&updatedExpense)

	gotErr := err
	gotTitle := updatedExpense.Title
	gotAmount := updatedExpense.Amount
	gotNote := updatedExpense.Note
	gotTags := updatedExpense.Tags
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
