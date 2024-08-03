package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/kordar/gocrud"
)

func InitCrudLangFn() {
	gocrud.SetLangFn(func(args ...interface{}) string {
		if args == nil || len(args) == 0 {
			return ""
		}
		c := args[0].(*gin.Context)
		locale := c.GetHeader(headerKey)
		if locale == "" {
			return defaultLocale
		}
		return locale
	})
}
