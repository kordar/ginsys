package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kordar/ginsys"
	"github.com/kordar/ginsys/middleware"
	"github.com/kordar/ginsys/resp"
	"github.com/kordar/ginsys/test/aaa/bbb"
	"github.com/kordar/ginsys/util"
	"github.com/kordar/gocfg"
	"github.com/kordar/goframework_resp"
	"github.com/kordar/gotrans"
	"github.com/kordar/govalidator"
	"testing"
)

func router(r *gin.Engine) *gin.Engine {

	r.GET("/hello", func(ctx *gin.Context) {
		goframework_resp.Success(ctx, "success", nil)
	})

	r.POST("/tt", func(ctx *gin.Context) {
		f := bbb.Demo001Form{}
		if err := util.DefaultGetValidParams(ctx, &f); err != nil {
			goframework_resp.Error(ctx, err, nil)
			return
		}
		goframework_resp.Success(ctx, "success", nil)
	})

	return r
}

type ValidTom struct {
}

func (v ValidTom) DefaultTpl() (tpl string, override bool) {
	return "", true
}

func (v ValidTom) Tpl() (section string, key string) {
	return "dictionary", "ttt"
}

func (v ValidTom) I18n(fe validator.FieldError, locale string) []string {
	return nil
}

func (v ValidTom) Tag() string {
	//TODO implement me
	return "tom"
}

func (v ValidTom) Valid(fl validator.FieldLevel) bool {
	//TODO implement me
	vv := fl.Field().String()
	fmt.Println("---------------------", vv)
	return false
}

func TestNameEqTomValid(t *testing.T) {

	gocfg.InitConfigWithParentDir("language", "ini")

	resp.InitJsonResp001()
	resp.InitI18nResponse()

	ginsys.NewGinServer().
		Router(router).
		OpenValidateAndTranslations(gotrans.NewEnTranslation(), gotrans.NewZhTranslation()).
		Middleware(middleware.RecoveryMiddleware(), resp.TransLocaleMiddleware("Locale", "zh")).
		AddValidate(ValidTom{}, govalidator.PhoneValidation{Rules: map[string]string{
			"r8": "^0\\w+4$",
		}}).
		Middleware(middleware.CorsMiddleware()).
		StartD("0.0.0.0", "9099")

}

func TestAA(t *testing.T) {
	//base64.StdEncoding.EncodeToString()
}
