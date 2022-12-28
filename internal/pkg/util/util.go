package util

import (
	"github.com/go-playground/validator"
)

func JsonHandler() jsonHandler {
	return newJsonHandler()
}

func Validator(v *validator.Validate) *customValidator {
	return newCustomValidator(v)
}

func TestHelper() testHelper {
	return newTestHelper()
}

func Error() *customError {
	return newCustomError()
}

func CompareError(e1 error, e2 error) bool {
	return e1.Error() == e2.Error()
}
