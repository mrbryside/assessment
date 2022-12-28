package expense

import "net/http"

func (e expenseCreator) CreateSuccess() expenseMock {
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
	return newExpenseMock(s, expenseJSON, expenseRespJSON, http.StatusCreated, called)
}

func (e expenseCreator) CreateBindFail() expenseMock {
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
	return newExpenseMock(s, expenseJSON, expenseRespJSON, http.StatusBadRequest, notCalled)

}

func (e expenseCreator) CreateValidateFail() expenseMock {
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
	return newExpenseMock(s, expenseJSON, expenseRespJSON, http.StatusBadRequest, notCalled)
}

func (e expenseCreator) CreateInternalFail() expenseMock {
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
	return newExpenseMock(s, expenseJSON, expenseRespJSON, http.StatusInternalServerError, called)
}
