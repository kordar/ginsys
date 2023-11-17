package ginsys

import "github.com/gin-gonic/gin"

func SetGinLevel(level string) {
	switch level {
	case "debug":
		gin.SetMode(gin.DebugMode)
		break
	case "test":
		gin.SetMode(gin.TestMode)
		break
	default:
		gin.SetMode(gin.ReleaseMode)
	}
}
