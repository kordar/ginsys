package ginsys

import (
	"github.com/kordar/ginsys/trans"
	"github.com/kordar/ginsys/valid"
)

var translations *trans.Translations

func GetTranslations() *trans.Translations {
	return translations
}

func InitValidateAndTranslations(tr ...trans.ITranslation) {
	validator := valid.GetValidator()
	if validator == nil {
		valid.InitValidator()
		validator = valid.GetValidator()
	}
	translations = trans.NewTranslations(validator).RegisterTranslators(tr...)
}

func RegI18n(tag string, section string, key string) {
	translations.RegisterTranslationWithGI18n(tag, section, key)
}

func RegValidation(va valid.IValidation) {
	valid.AddValidation(va)
	if section, key := va.I18n(); section != "" && key != "" {
		translations.RegisterTranslationWithGI18n(va.Tag(), section, key)
	}
}
