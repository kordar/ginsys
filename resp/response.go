package resp

import (
	"fmt"
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
	headerKey      = "Locale"
	defaultLocale  = "en"
)

func SetHeaderKey(key string) {
	headerKey = key
}

func SetDefaultLocale(name string) {
	defaultLocale = name
}

func SetTransLocaleMapValue(key string, value string) {
	translocalemap[key] = value
}

// GetRealLocale 翻译器的国际化name和本地上传的locale名称可能不符，
// 例如实际locale=zh-CN，翻译器注册中文名为zh，
// 此时需要将zh-CN映射为zh以便正确获取翻译器。
func GetRealLocale(locale string) string {
	if translocalemap[locale] == "" {
		return defaultLocale
	}
	return translocalemap[locale]
}

func getlocale(c *gin.Context) string {
	locale := c.GetHeader(headerKey)
	if locale == "" {
		return defaultLocale
	}
	return locale
}

var i18nFunc = func(message string, messagetype string, c interface{}) string {
	ctx := c.(*gin.Context)
	locale := getlocale(ctx)
	if messagetype == response.SuccessType {
		return gocfg.GetSectionValue(locale, fmt.Sprintf("response.success.%s", message), "language")
	} else if messagetype == response.ErrorType {
		return gocfg.GetSectionValue(locale, fmt.Sprintf("response.errors.%s", message), "language")
	} else {
		return gocfg.GetSectionValue(locale, fmt.Sprintf("response.common.%s", message), "language")
	}
}

func SetI18nFunc(f func(message string, messagetype string, c interface{}) string) {
	i18nFunc = f
}

func gettrans(c interface{}) (trans ut.Translator, found bool) {
	ctx := c.(*gin.Context)
	locale := getlocale(ctx)
	return gotrans.Get().GetTranslator(GetRealLocale(locale))
}

func InitI18nResponse() {
	response.RegRespFunc(response.SuccessType, responseI18n.SuccessResultI18n{I18nMessage: i18nFunc})
	response.RegRespFunc(response.ErrorType, responseI18n.ErrorResultI18n{I18nMessage: i18nFunc, GetTrans: gettrans})
	response.RegRespFunc(response.ValidErrorType, responseI18n.ErrorResultI18n{I18nMessage: i18nFunc, GetTrans: gettrans})
	response.RegRespFunc(response.OutputType, responseI18n.OutputResponseI18n{I18nMessage: i18nFunc})
	response.RegRespFunc(response.UnauthorizedType, responseI18n.UnauthorizedJsonI18n{I18nMessage: i18nFunc})
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
	if err, ok := message.(error); ok {
		response.GetResultCallFunc()(c, e.HttpStatus(), response.Code("error"), err.Error(), data, count)
		return
	}
	response.GetResultCallFunc()(c, e.HttpStatus(), response.Code("error"), message.(string), data, count)
}

func InitI18nResponse002() {
	response.RegRespFunc(response.ErrorType, ErrWithValidate{})
	response.RegRespFunc(response.ValidErrorType, ErrWithValidate{})
}
