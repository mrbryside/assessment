package util

type customError struct {
	DBNotFound *DbNotFoundError
}

func newCustomError() *customError {
	return &customError{DBNotFound: newDBNotFoundError()}
}

type DbNotFoundError struct{}

func newDBNotFoundError() *DbNotFoundError {
	return &DbNotFoundError{}
}

func (c *DbNotFoundError) Error() string {
	return "db query not found"
}
