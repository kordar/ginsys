package middleware

import (
	"github.com/gin-gonic/gin"
	response "github.com/kordar/goframework_resp"
	"runtime"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if val, ok := err.(string); ok {
					response.Error(c, val, nil)
				}
				if val, ok := err.(runtime.Error); ok {
					response.Error(c, val.Error(), nil)
				}
				return
			}
		}()
		c.Next()
	}
}
