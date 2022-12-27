//go:build unit

package expense

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/assessment/internal/pkg/expense/mock"
	"github.com/mrbryside/assessment/internal/pkg/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var tests = []struct {
	name string
	mock mock.CreateExpenseMock
}{
	{name: "should return response expense data", mock: mock.CreationMock().CreateSuccess()},
	{name: "should return response bad request invalid", mock: mock.CreationMock().CreateBindFail()},
	{name: "should return response bad request required field", mock: mock.CreationMock().CreateValidateFail()},
	{name: "should return response internal server error", mock: mock.CreationMock().CreateInternalFail()},
}

func TestCreateExpense(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing Suite")
}

var _ = Describe("create expenses", func() {
	for _, tc := range tests {
		tc := tc
		Context("Create expense when body is "+tc.mock.Payload, func() {
			It(tc.name, func() {
				// Arrange
				mock := tc.mock
				expenses := NewExpense(mock.SpyStore)
				e := echo.New()
				e.Validator = util.Validator(validator.New())
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(mock.Payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				err := expenses.CreateExpenseHandler(c)

				wantResp := mock.Response
				wantCalled := mock.Called
				wantCode := mock.Code

				// Act
				gotErr := err
				gotResp := rec.Body.String()
				gotCalled := mock.SpyStore.CreateWasCalled()
				gotCode := c.Response().Status

				// Assert
				Expect(gotErr).ShouldNot(HaveOccurred())
				Expect(gotCalled).To(BeEquivalentTo(wantCalled))
				Expect(gotResp).To(MatchJSON(wantResp))
				Expect(gotCode).To(BeEquivalentTo(wantCode))
			})
		})
	}
})
