package middleware

import (
	"github.com/gin-gonic/gin"
)

// TransLocaleMiddleware //// 转换真实locale值
func TransLocaleMiddleware(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := i18n.getlocale(c)
		c.Request.Header.Set(key, i18n.GetRealLocale(s))
		// 处理请求
		c.Next() //  处理请求
	}
}
