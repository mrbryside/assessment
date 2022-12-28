package expense

import (
	"net/http"
)

func (e expenseGetter) GetExpenseSuccess() expenseMock {
	var (
		queryId         = "5"
		expenseRespJSON = `{
			"id": 5,
			"title": "strawberry smoothie",
    		"amount": 79,
    		"note": "night market promotion discount 10 bath",
    		"tags": ["food", "beverage"]
			}`
	)

	s := newSpyStoreWithGetExpenseSuccess()
	return newExpenseMock(s, queryId, expenseRespJSON, http.StatusOK, called)
}

func (e expenseGetter) GetExpenseValidateFailed() expenseMock {
	var (
		queryId         = ""
		expenseRespJSON = `{
			"code": "4000",
			"message": "ID is a required field"
		}`
	)

	s := newSpyStoreWithGetExpenseSuccess()
	return newExpenseMock(s, queryId, expenseRespJSON, http.StatusBadRequest, notCalled)
}

func (e expenseGetter) GetExpenseBindFailed() expenseMock {
	var (
		queryId         = "error"
		expenseRespJSON = `{
			"code": "4000",
			"message": "Request parameter is invalid."
		}`
	)

	s := newSpyStoreWithGetExpenseSuccess()
	return newExpenseMock(s, queryId, expenseRespJSON, http.StatusBadRequest, notCalled)
}

func (e expenseGetter) GetExpenseInternalFailed() expenseMock {
	var (
		queryId         = "5"
		expenseRespJSON = `{
			"code": "5000",
			"message": "internal server error"
		}`
	)

	s := newSpyStoreWithGetExpenseFail()
	return newExpenseMock(s, queryId, expenseRespJSON, http.StatusInternalServerError, called)
}

func (e expenseGetter) GetExpenseNotFound() expenseMock {
	var (
		queryId         = "5"
		expenseRespJSON = `{
			"code": "4004",
			"message": "expense not found"
		}`
	)

	s := newSpyStoreWithGetExpenseNotFound()
	return newExpenseMock(s, queryId, expenseRespJSON, http.StatusNotFound, called)
}
