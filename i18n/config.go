package i18n

import (
	"github.com/gin-gonic/gin"
	"github.com/kordar/gocfg"
	"github.com/kordar/goframework_resp"
)

func InitI18n(dir string) {
	gocfg.InitConfigWithSubDir(dir, "ini", "toml", "yaml")
	// 配置response国际化支持
	goframework_resp.RegResultCallFunc(func(c interface{}, httpStatus int, code int, message string, data interface{}, count int64) {
		ctx := c.(*gin.Context)
		if data == nil {
			ctx.JSON(httpStatus, map[string]interface{}{"code": code, "message": message})
			return
		}

		if count < 0 {
			ctx.JSON(httpStatus, map[string]interface{}{"code": code, "message": message, "data": data})
			return
		}

		ctx.JSON(httpStatus, map[string]interface{}{"code": code, "message": message, "data": data, "count": count})
	})
}
