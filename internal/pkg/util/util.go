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
