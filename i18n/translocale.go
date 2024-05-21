package i18n

import "github.com/gin-gonic/gin"

// TransLocaleMiddleware // 转换真实locale值
func TransLocaleMiddleware(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := getlocale(c)
		c.Request.Header.Set(key, GetRealLocale(s))
		// 处理请求
		c.Next() //  处理请求
	}
}
