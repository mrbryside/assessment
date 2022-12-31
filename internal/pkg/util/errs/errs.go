package errs

type customError struct {
	DBNotFound DbNotFoundError
}

func newCustomError() customError {
	return customError{DBNotFound: newDBNotFoundError()}
}

func Error() customError {
	return newCustomError()
}

type DbNotFoundError struct{}

func newDBNotFoundError() DbNotFoundError {
	return DbNotFoundError{}
}

func (d DbNotFoundError) Error() string {
	return "db query not found"
}

// CompareError -- another error util
func CompareError(e1 error, e2 error) bool {
	return e1.Error() == e2.Error()
}
