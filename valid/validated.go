package valid

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var _validate *validator.Validate

type IValidation interface {
	Tag() string
	Valid(fl validator.FieldLevel) bool
	I18n() (section string, key string)
}

func InitValidator() {
	_validate, _ = binding.Validator.Engine().(*validator.Validate)
}

func AddValidation(validation IValidation) {
	_ = _validate.RegisterValidation(validation.Tag(), validation.Valid)
}

func GetValidator() *validator.Validate {
	return _validate
}
