//go:build integration

package expense

import (
	"bytes"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/middleware"
	"github.com/mrbryside/assessment/internal/pkg/util/httputil"
	"github.com/stretchr/testify/assert"
	"net/http"
	"regexp"
	"testing"
	"time"
)

const (
	seedPayload = `{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath",
		"tags": ["food", "beverage"]
	}`
	getsCode = http.StatusOK
)

var (
	createdSeed  modelExpense
	getsResponse = `[{
		"id": 1,
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath",
		"tags": ["food", "beverage"]
	}]`
)

func TestIntegrationGetExpenses(t *testing.T) {
	// Setup server
	database, err := db.InitDB(db.NewPostgres("postgresql://root:root@db/test-db?sslmode=disable"))
	if err != nil {
		t.Fatal("can't init db")
	}
	defer database.Close()
	expenses := NewExpense(db.DB)
	_, err = database.Exec("TRUNCATE expenses;")
	eh := echo.New()
	eh.Use(middleware.VerifyAuthorization)
	eh = httputil.InitItEcho(eh, func() {
		eh.POST("/expenses", expenses.CreateExpenseHandler)
		eh.GET("/expenses", expenses.GetExpensesHandler)
	})

	// Arrange
	err = httputil.Seeder(&createdSeed, seedPayload, "expenses")
	if err != nil {
		t.Fatal("can't create expense:", err)
	}
	re := regexp.MustCompile(`"id": \d+`)

	wantResp := re.ReplaceAllString(getsResponse, fmt.Sprintf(`"id": %d`, createdSeed.ID))
	wantCode := getsCode

	// Act
	res := httputil.Request(http.MethodGet, httputil.Uri("expenses"), nil)
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		t.Fatal("can't read response body:", err)
	}
	bodyStr := buf.String()
	gotErr := err
	gotCode := res.StatusCode

	// Assert
	assert.Nil(t, gotErr)
	assert.JSONEq(t, bodyStr, wantResp)
	assert.Equal(t, wantCode, gotCode)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
}
