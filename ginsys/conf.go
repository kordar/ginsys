package ginsys

import "github.com/gin-gonic/gin"

func SetGinLevel(level string) {
	switch level {
	case "debug":
		gin.SetMode(gin.DebugMode)
		return
	case "test":
		gin.SetMode(gin.TestMode)
		return
	default:
		gin.SetMode(gin.ReleaseMode)
	}
}
