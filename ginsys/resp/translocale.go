package resp

import "github.com/gin-gonic/gin"

// TransLocaleMiddleware // 转换真实locale值
func TransLocaleMiddleware(key string, defaultLocale string) gin.HandlerFunc {
	SetHeaderKey(key)
	SetDefaultLocale(defaultLocale)
	return func(c *gin.Context) {
		s := getlocale(c)
		// 重新设置header头
		c.Request.Header.Set(key, GetRealLocale(s))
		// 处理请求
		c.Next() //  处理请求
	}
}
