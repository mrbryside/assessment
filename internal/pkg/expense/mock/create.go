package mock

import "net/http"

const (
	called    = true
	notCalled = false
)

type CreateExpenseMock struct {
	SpyStore spyStore
	Payload  string
	Response string
	Code     int
	Called   bool
}

func newCreateExpenseMock(s spyStore, p string, r string, c int, ca bool) CreateExpenseMock {
	return CreateExpenseMock{
		SpyStore: s,
		Payload:  p,
		Response: r,
		Code:     c,
		Called:   ca,
	}
}

type expenseCreator struct{}

func (e expenseCreator) CreateSuccess() CreateExpenseMock {
	var (
		expenseJSON = `{
			"title": "strawberry smoothie",
    		"amount": 79,
    		"note": "night market promotion discount 10 bath", 
    		"tags": ["food", "beverage"]
		}`
		expenseRespJSON = `{
			"id": 5,
			"title": "strawberry smoothie",
    		"amount": 79,
    		"note": "night market promotion discount 10 bath", 
    		"tags": ["food", "beverage"]
		}`
	)

	s := newSpyStoreWithCreateSuccess()
	return newCreateExpenseMock(s, expenseJSON, expenseRespJSON, http.StatusCreated, called)
}

func (e expenseCreator) CreateBindFail() CreateExpenseMock {
	var (
		expenseJSON = `{
			"title": "strawberry smoothie",
			"amount": "12345",
			"note": "night market promotion discount 10 bath",
			"tags": ["food", "beverage"]
		}`
		expenseRespJSON = `{
			"code": "4000",
			"message": "Request parameters are invalid."
		}`
	)

	s := newSpyStoreWithCreateSuccess()
	return newCreateExpenseMock(s, expenseJSON, expenseRespJSON, http.StatusBadRequest, notCalled)

}

func (e expenseCreator) CreateValidateFail() CreateExpenseMock {
	var (
		expenseJSON = `{
			"title": "strawberry smoothie",
			"note": "night market promotion discount 10 bath",
			"tags": ["food", "beverage"]
		}`
		expenseRespJSON = `{
			"code": "4000",
			"message": "Amount is a required field"
		}`
	)

	s := newSpyStoreWithCreateSuccess()
	return newCreateExpenseMock(s, expenseJSON, expenseRespJSON, http.StatusBadRequest, notCalled)
}

func (e expenseCreator) CreateInternalFail() CreateExpenseMock {
	var (
		expenseJSON = `{
			"title": "strawberry smoothie",
			"amount": 79,
			"note": "night market promotion discount 10 bath",
			"tags": ["food", "beverage"]
		}`
		expenseRespJSON = `{
			"code": "5000",
			"message": "internal server error"
		}`
	)

	s := newSpyStoreWithCreateFail()
	return newCreateExpenseMock(s, expenseJSON, expenseRespJSON, http.StatusInternalServerError, called)
}
