package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var container = map[string]IResponse{
	"success": SuccessJson{},
	"error":   ErrorJson{},
}

func RegResponseFunc(name string, response IResponse) {
	container[name] = response
}

func ToJson(c *gin.Context, t string, message interface{}, data interface{}, count int64) {
	response := container[t]
	if response == nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 0, "message": "not found \"response object!\""})
		return
	}

	response.Result(c, message, data, count)
}

func Data(c *gin.Context, message string, data interface{}, count int64) {
	ToJson(c, SuccessType, message, data, count)
}

func Success(c *gin.Context, message string, data interface{}) {
	ToJson(c, SuccessType, message, data, -1)
}

func Tenant(c *gin.Context, message string, data interface{}) {
	ToJson(c, TenantType, message, data, -1)
}

func Error(c *gin.Context, err interface{}, data interface{}) {
	ToJson(c, ErrorType, err, data, -1)
}

func ErrorV(c *gin.Context, err interface{}, data interface{}) {
	ToJson(c, ValidErrorType, err, data, -1)
}

func Unauthorized(c *gin.Context, message string, data interface{}) {
	ToJson(c, UnauthorizedType, message, data, -1)
}

func SuccessOrError(c *gin.Context, flag bool, successMessage string, err interface{}) {
	if flag {
		ToJson(c, SuccessType, successMessage, nil, -1)
	} else {
		ToJson(c, ErrorType, err, nil, -1)
	}
}

func ErrorOrSuccess(c *gin.Context, err error) {
	if err == nil {
		ToJson(c, SuccessType, "success", nil, -1)
	} else {
		ToJson(c, ErrorType, err, nil, -1)
	}
}
