package resp

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/kordar/gocfg"
	response "github.com/kordar/goframework_resp"
	responseI18n "github.com/kordar/goframework_resp_i18n"
	"github.com/kordar/gotrans"
	"net/http"
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

var i18n = func(message string, messagetype string, c interface{}) string {
	ctx := c.(*gin.Context)
	locale := getlocale(ctx)
	if messagetype == response.SuccessType {
		return gocfg.GetGroupSectionValue(locale, "response.success", message)
	} else if messagetype == response.ErrorType {
		return gocfg.GetGroupSectionValue(locale, "response.errors", message)
	} else {
		return gocfg.GetGroupSectionValue(locale, "response.common", message)
	}
}

func SetI18nFunc(f func(message string, messagetype string, c interface{}) string) {
	i18n = f
}

func gettrans(c interface{}) (trans ut.Translator, found bool) {
	ctx := c.(*gin.Context)
	locale := getlocale(ctx)
	return gotrans.GetTranslations().GetTrans(GetRealLocale(locale))
}

type ErrWithValidate struct {
}

func (e ErrWithValidate) HttpStatus() int {
	return http.StatusOK
}

func (e ErrWithValidate) Result(c interface{}, message interface{}, data interface{}, count int64) {
	if err, ok := message.(validator.ValidationErrors); ok {
		for _, ee := range err {
			response.GetResultCallFunc()(c, e.HttpStatus(), response.Code("error"), ee.Error(), data, count)
			return
		}
	}
	response.GetResultCallFunc()(c, e.HttpStatus(), response.Code("error"), message.(string), data, count)
}

func InitI18nResponse() {
	response.RegRespFunc(response.SuccessType, responseI18n.SuccessResultI18n{I18nMessage: i18n})
	response.RegRespFunc(response.ErrorType, responseI18n.ErrorResultI18n{I18nMessage: i18n, GetTrans: gettrans})
	response.RegRespFunc(response.ValidErrorType, responseI18n.ErrorResultI18n{I18nMessage: i18n, GetTrans: gettrans})
	response.RegRespFunc(response.OutputType, responseI18n.OutputResponseI18n{I18nMessage: i18n})
	response.RegRespFunc(response.UnauthorizedType, responseI18n.UnauthorizedJsonI18n{I18nMessage: i18n})
}

func InitI18nResponse002() {
	response.RegRespFunc(response.ErrorType, ErrWithValidate{})
	response.RegRespFunc(response.ValidErrorType, ErrWithValidate{})
}
