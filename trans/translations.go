package trans

import (
	"fmt"
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/kordar/goi18n"
)

type ITranslation interface {
	GetTranslator() locales.Translator
	RegisterValidate(trans ut.Translator, validate *validator.Validate) error
}

type Translations struct {
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    map[string]ut.Translator
}

func NewTranslations(validate *validator.Validate) *Translations {
	translation := NewEnTranslation()
	translator := translation.GetTranslator()
	uni := ut.New(translator)
	trans, _ := uni.GetTranslator(translator.Locale())
	_ = translation.RegisterValidate(trans, validate)
	return &Translations{uni: uni, validate: validate, trans: map[string]ut.Translator{trans.Locale(): trans}}
}

func (t *Translations) RegisterTranslators(translators ...ITranslation) *Translations {
	for i := range translators {
		translator := translators[i].GetTranslator()
		err := t.uni.AddTranslator(translator, true)
		if err == nil {
			locale := translator.Locale()
			if trans, found := t.GetTrans(locale); found {
				t.trans[locale] = trans
				_ = translators[i].RegisterValidate(trans, t.validate)
			}
		}
	}
	return t
}

func (t *Translations) GetTrans(locale string) (trans ut.Translator, found bool) {
	return t.uni.GetTranslator(locale)
}

func (t *Translations) GetValidate() *validator.Validate {
	return t.validate
}

func (t *Translations) RegisterTranslationWithGI18n(tag string, section string, key string) *Translations {
	for locale, trans := range t.trans {
		err := t.validate.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
			text := goi18n.GetSectionValue(locale, section, key, "ini").(string)
			return ut.Add(tag, text, true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			text := goi18n.GetSectionValue(locale, "dictionary", fe.Field(), "ini").(string)
			if text == "" {
				text = fe.Field()
			}
			t2, _ := ut.T(tag, text)
			return t2
		})
		if err != nil {
			fmt.Printf("%s注册翻译器失败！\n", locale)
		}
	}
	return t
}
