package ginsys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/kordar/gocfg"
	"github.com/kordar/gotrans"
	"github.com/kordar/govalidator"
)

type GinServer struct {
	r *gin.Engine
}

func (g *GinServer) R() *gin.Engine {
	return g.r
}

func NewGinEngineServer(engine *gin.Engine) *GinServer {
	mode := gocfg.GetSystemValue("gin_mode")
	if mode != "" {
		SetGinLevel(mode)
	}

	validate := binding.Validator.Engine().(*validator.Validate)
	govalidator.LoadValidate(validate)

	return &GinServer{engine}
}

func NewGinServer() *GinServer {
	return NewGinEngineServer(gin.Default())
}

func (g *GinServer) OpenValidateAndTranslations(tr ...gotrans.GoTranslation) *GinServer {
	gotrans.Initialize(tr...)
	return g
}

func (g *GinServer) AddValidate(validations ...govalidator.IValidation) *GinServer {
	for _, validation := range validations {
		govalidator.AddValidation(validation)
		if !gotrans.Exists() {
			continue
		}
		trans := gotrans.Get()
		trans.BindTranslatorToValidate(
			validation.Tag(),
			func(locale string) (string, bool) {
				defaultTpl, override := validation.DefaultTpl()
				section, key := validation.Tpl()
				if section == "" || key == "" {
					return defaultTpl, override
				}
				sk := fmt.Sprintf("%s.%s", section, key)
				value := gocfg.GetSectionValueM(sk, "language", locale)
				if value == "" {
					return defaultTpl, override
				} else {
					return value, override
				}
			},
			func(locale string, fe validator.FieldError) []string {
				n := validation.I18n(fe, locale)
				if n == nil || len(n) == 0 {
					text := gocfg.GetSectionValueM("dictionary."+fe.StructNamespace(), "language", locale)
					if text == "" {
						text = fe.Field()
					}
					return []string{text}
				}
				//logger.Infof("=========field======%+v", fe.Field())
				//logger.Infof("=========param======%+v", fe.Param())
				//logger.Infof("=========tag======%+v", fe.Tag())
				//logger.Infof("=========error======%+v", fe.Error())
				//logger.Infof("=========StructField======%+v", fe.StructField())
				//logger.Infof("=========Namespace======%+v", fe.Namespace())
				//logger.Infof("=========StructNamespace======%+v", fe.StructNamespace())
				//logger.Infof("=========ActualTag======%+v", fe.ActualTag())
				return n
			})
	}
	return g
}

func (g *GinServer) Start() {
	serverName := gocfg.GetSectionValue("server", "host")
	serverPort := gocfg.GetSectionValue("server", "port")
	g.StartD(serverName, serverPort)
}

// StartD 启动服务
func (g *GinServer) StartD(serverName string, serverPort string) {
	_ = g.r.Run(serverName + ":" + serverPort)
}

func (g *GinServer) Router(f func(r *gin.Engine) *gin.Engine) *GinServer {
	g.r = f(g.r)
	return g
}

func (g *GinServer) Middleware(middleware ...gin.HandlerFunc) *GinServer {
	g.r.Use(middleware...)
	return g
}
