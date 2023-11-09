package validate

import (
	"github.com/go-playground/validator/v10"
)

type NotEmptyValidator struct {
}

func (n NotEmptyValidator) Tag() string {
	//TODO implement me
	return "notempty"
}

func (n NotEmptyValidator) Valid(fl validator.FieldLevel) bool {
	return fl.Field().Len() > 0
}

func (n NotEmptyValidator) I18n() (section string, key string) {
	//TODO implement me
	return "validate.default", "notempty"
}
