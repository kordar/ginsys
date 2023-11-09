package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type PasswordValidator struct {
}

func (n PasswordValidator) Tag() string {
	//TODO implement me
	return "password"
}

func (n PasswordValidator) Valid(fl validator.FieldLevel) bool {
	mobileNum := fl.Field().String()
	fl.StructFieldName()
	regular := "^[a-zA-Z0-9_-]{8,16}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func (n PasswordValidator) I18n() (section string, key string) {
	//TODO implement me
	return "validate.default", "password"
}
