package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IResponse interface {
	HttpStatus() int
	Result(c *gin.Context, message interface{}, data interface{}, count int64)
}

// SuccessJson 成功
type SuccessJson struct {
}

func (s SuccessJson) HttpStatus() int {
	return http.StatusOK
}

func (s SuccessJson) Result(c *gin.Context, message interface{}, data interface{}, count int64) {
	if value, ok := message.(string); ok && value != "" {
		jsonCall(c, s.HttpStatus(), Code("success"), value, data, count)
		return
	}

	jsonCall(c, s.HttpStatus(), Code("success"), "success", data, count)
}

// ----------- Error ------------

type ErrorJson struct {
}

func (e ErrorJson) HttpStatus() int {
	return http.StatusOK
}

func (e ErrorJson) Result(c *gin.Context, message interface{}, data interface{}, count int64) {
	//TODO implement me
	if value, ok := message.(string); ok && value != "" {
		jsonCall(c, e.HttpStatus(), Code("error"), value, data, count)
		return
	}

	if value, ok := message.(error); ok {
		jsonCall(c, e.HttpStatus(), Code("error"), value.Error(), data, count)
		return
	}

	jsonCall(c, e.HttpStatus(), Code("error"), "error", data, count)

}

// -------------- Output Excel -------------------------

type OutputResponse struct {
}

func (o OutputResponse) HttpStatus() int {
	return http.StatusOK
}

func (o OutputResponse) Result(c *gin.Context, message interface{}, data interface{}, count int64) {
	// TODO implement me
	if value, ok := data.(IOutput); ok && value != nil {
		if value.Type() == "browser" {
			// output web
			for k, v := range value.Header() {
				c.Header(k, v)
			}
			_, _ = c.Writer.Write(value.Data())
			return
		}

		msg := "output success!"
		if tmessage, found := message.(string); found && tmessage != "" {
			msg = tmessage
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"code": Code("success"), "message": msg, "data": value.Data(), "params": value.Params(),
		})
		return
	}

	jsonCall(c, o.HttpStatus(), Code("fail"), "output fail!", nil, -1)

}

// -------------- Unauthorized Excel -------------------------

type UnauthorizedResponse struct {
}

func (o UnauthorizedResponse) HttpStatus() int {
	return http.StatusOK
}

func (o UnauthorizedResponse) Result(c *gin.Context, message interface{}, data interface{}, count int64) {
	jsonCall(c, o.HttpStatus(), Code("unauthorized"), "unauthorized", nil, -1)
}

// -------------- MultiTenant -------------------------

type MultiTenantResponse struct {
}

func (o MultiTenantResponse) HttpStatus() int {
	return http.StatusOK
}

func (o MultiTenantResponse) Result(c *gin.Context, message interface{}, data interface{}, count int64) {
	jsonCall(c, o.HttpStatus(), Code("multitenant"), "multiTenant", data, -1)
}
