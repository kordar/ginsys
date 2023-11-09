package ginsys

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kordar/ginsys/valid"
	"reflect"
)

func DefaultGetValidParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	// 获取翻译器
	validator := valid.GetValidator()
	if validator == nil {
		return errors.New("validator nil")
	}
	err := validator.Struct(params)
	if err != nil {
		return err
	}
	refParams := reflect.ValueOf(params) // 需要传入指针，后面再解析
	validMethod := refParams.MethodByName("Valid")
	if validMethod.IsValid() {
		v := validMethod.Call(make([]reflect.Value, 0))
		if e := v[0].Interface(); e != nil {
			return e.(error)
		}
	}
	return nil
}

func CtxGetValidParams(c *gin.Context, params interface{}, targetService interface{}, methods ...string) error {
	// 获取验证器
	validator := valid.GetValidator()
	if validator == nil {
		return errors.New("validator nil")
	}

	err := validator.Struct(params)
	if err != nil {
		return err
	}

	if targetService == nil {
		return nil
	}

	refParams := reflect.ValueOf(targetService) // 需要传入指针，后面再解析
	for _, method := range methods {
		validMethod := refParams.MethodByName(method)
		if validMethod.IsValid() {
			P := make([]reflect.Value, 2)
			P[0] = reflect.ValueOf(c)
			P[1] = reflect.ValueOf(params)
			v := validMethod.Call(P)
			if e := v[0].Interface(); e != nil {
				return e.(error)
			}
		}
	}
	return nil
}
