package util

import "github.com/go-playground/validator"

// all register utility

func Error() *customError {
	return newCustomError()
}

func TestHelper() testHelper {
	return newTestHelper()
}

func Validator(v *validator.Validate) *customValidator {
	return newCustomValidator(v)
}
