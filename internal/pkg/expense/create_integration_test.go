//go:build integration

package expense

import (
	"bytes"
	"context"
	"github.com/mrbryside/assessment/internal/pkg/db"
	"github.com/mrbryside/assessment/internal/pkg/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"testing"
	"time"
)

var (
	mockCreation = CreationMock().CreateSuccess()
	serverPort   = ":2565"
)

func TestItCreateExpense(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing Suite")
}

var _ = Describe("create expenses", func() {
	Context("Create expense when body is "+mockCreation.Payload, func() {
		It("integration test create expense", func() {
			// Setup server
			db.InitDB(db.NewPostgres("postgresql://root:root@db/test-db?sslmode=disable"))
			expenses := NewExpense(db.DB)
			eh := util.TestHelper().InitItEcho(expenses.CreateExpenseHandler, "/expenses")
			// Arrange
			payload := mockCreation.Payload
			wantCode := mockCreation.Code
			wantTitle := "strawberry smoothie"
			wantAmount := 79
			wantNote := "night market promotion discount 10 bath"
			wantTags := []string{"food", "beverage"}

			// Act
			var e ModelDto
			body := bytes.NewBufferString(payload)
			testHelper := util.TestHelper()
			res := testHelper.Request(http.MethodPost, testHelper.Uri("expenses"), body)
			err := res.Decode(&e)

			gotErr := err
			gotTitle := e.Title
			gotAmount := e.Amount
			gotNote := e.Note
			gotTags := e.Tags
			gotCode := res.StatusCode

			// Assert
			Expect(gotErr).ShouldNot(HaveOccurred())
			Expect(gotCode).To(BeEquivalentTo(wantCode))
			Expect(gotTitle).To(BeEquivalentTo(wantTitle))
			Expect(gotAmount).To(BeEquivalentTo(wantAmount))
			Expect(gotNote).To(BeEquivalentTo(wantNote))
			Expect(gotTags).To(BeEquivalentTo(wantTags))

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			err = eh.Shutdown(ctx)
		})
	})
})
