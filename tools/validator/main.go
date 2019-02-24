package validator

import (
	validatorV9 "gopkg.in/go-playground/validator.v9"
	"time"
)

var validator = validatorV9.New()

func init() {
	validator.RegisterValidation("datetime", isDatetime)
}

func Process(i interface{}) error {
	return validator.Struct(i)
}

func isDatetime(f validatorV9.FieldLevel) bool {
	v := f.Field().String()
	_, err := time.Parse(time.RFC3339, v)
	return err == nil
}
