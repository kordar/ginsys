package ginsys

import (
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

func (g *GinServer) OpenValidateAndTranslations(tr ...gotrans.ITranslation) *GinServer {
	gotrans.InitValidateAndTranslations(tr...)
	return g
}

func (g *GinServer) AddValidate(valid ...govalidator.IValidation) *GinServer {
	for i := range valid {
		govalidator.AddValidation(valid[i])
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
