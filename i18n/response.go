package i18n

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/kordar/ginsys"
	response "github.com/kordar/ginsys/resp"
	"github.com/kordar/goi18n"
)

var (
	translocalemap = map[string]string{"zh-CN": "zh"}
)

func SetTransLocaleMapValue(key string, value string) {
	translocalemap[key] = value
}

func GetRealLocale(locale string) string {
	if translocalemap[locale] == "" {
		return locale
	}
	return translocalemap[locale]
}

func getlocale(c *gin.Context) string {
	locale := c.GetHeader("Locale")
	if locale == "" {
		return "en"
	}
	return locale
}

var i18n = func(message string, messagetype string, ctx *gin.Context) string {
	locale := getlocale(ctx)
	if messagetype == response.SuccessType {
		return goi18n.GetSectionValue(locale, "response.success", message, "ini").(string)
	} else if messagetype == response.ErrorType {
		return goi18n.GetSectionValue(locale, "response.errors", message, "ini").(string)
	} else {
		return goi18n.GetSectionValue(locale, "response.common", message, "ini").(string)
	}
}

func SetI18nFunc(f func(message string, messagetype string, ctx *gin.Context) string) {
	i18n = f
}

func gettrans(ctx *gin.Context) (trans ut.Translator, found bool) {
	locale := getlocale(ctx)
	return ginsys.GetTranslations().GetTrans(GetRealLocale(locale))
}

func initResponse() {
	response.RegResponseFunc(response.SuccessType, response.SuccessJsonI18n{I18nMessage: i18n})
	response.RegResponseFunc(response.ErrorType, response.ErrorJsonI18n2{I18nMessage: i18n, GetTrans: gettrans})
	response.RegResponseFunc(response.ValidErrorType, response.ErrorJsonI18n{I18nMessage: i18n, GetTrans: gettrans})
	response.RegResponseFunc(response.OutputType, response.OutputResponseI18n{I18nMessage: i18n})
	response.RegResponseFunc(response.UnauthorizedType, response.UnauthorizedJsonI18n{I18nMessage: i18n})
	response.RegResponseFunc(response.TenantType, response.MultiTenantResponse{})
}
