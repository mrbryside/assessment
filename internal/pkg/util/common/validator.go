package common

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator"
)

type customValidator struct {
	validator *validator.Validate
}

func newCustomValidator(v *validator.Validate) *customValidator {
	return &customValidator{validator: v}
}

func Validator(v *validator.Validate) *customValidator {
	return newCustomValidator(v)
}

func (c customValidator) Validate(i interface{}) error {
	// Validate the input using the validation package
	err := c.validator.Struct(i)
	if err == nil {
		return nil
	}

	var message string
	for _, e := range err.(validator.ValidationErrors) {
		// Customize the error message for each validation error
		if e.Tag() == "required" {
			message = fmt.Sprintf("%s is a required field", e.Field())
			continue
		}
		message = fmt.Sprintf("%s is invalid", e.Field())
	}
	return errors.New(message)
}
