package ginsys

import (
	"github.com/gin-gonic/gin"
	"github.com/kordar/ginsys/trans"
	"github.com/kordar/ginsys/valid"
	"github.com/kordar/gocfg"
)

type GinServer struct {
	r *gin.Engine
}

func (g *GinServer) R() *gin.Engine {
	return g.r
}

func NewGinServer() *GinServer {
	return &GinServer{gin.Default()}
}

func (g *GinServer) OpenValidateAndTranslations(tr ...trans.ITranslation) *GinServer {
	InitValidateAndTranslations(tr...)
	return g
}

func (g *GinServer) AddValidate(valid ...valid.IValidation) *GinServer {
	for i := range valid {
		RegValidation(valid[i])
	}
	return g
}

func (g *GinServer) Start() {
	serverName := gocfg.GetSectionValue("server", "host")
	serverPort := gocfg.GetSectionValue("server", "port")
	// 启动服务
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
