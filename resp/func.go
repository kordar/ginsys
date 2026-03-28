package resp

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kordar/gocfg"
	"github.com/kordar/gocrud"
)

func InitCrudLangFn() {
	gocrud.MessageFn = func(c context.Context, message string) string {
		ctx := c.(*gin.Context)
		locale := getlocale(ctx)
		return gocfg.GetSectionValueM(fmt.Sprintf("%s.gocrud.message", locale), message, "language")
	}
}

func GetLocale(c context.Context) string {
	ctx := c.(*gin.Context)
	return getlocale(ctx)
}
