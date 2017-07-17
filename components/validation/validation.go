package validation

import (
	"../../modules/myresume/model"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func ValidateMe(st interface{}) error {
	validate = validator.New()
	object, isOk := st.(model.User)
	if !isOk {
		return nil
	}
	err := validate.Struct(object)
	return err
}
