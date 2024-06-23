package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kordar/govalidator"
	"reflect"
)

// DefaultGetValidParams 手动触发翻译器
func DefaultGetValidParams(c *gin.Context, params interface{}) error {
	// 1、触发gin自带的翻译组件进行参数验证
	if err := c.ShouldBind(params); err != nil {
		return err
	}

	/**
	 * 2、获取系统中现有的翻译组件（一般服务启动将gin默认翻译组件注入）
	 *    validate := binding.Validator.Engine().(*validator.Validate)
	 *    govalidator.LoadValidate(validate)
	 *
	 */
	validate := govalidator.GetValidate()
	if validate == nil {
		return errors.New("no active \"validate\" found")
	}

	// 3、手动触发验证
	err := validate.Struct(params)
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
	validate := govalidator.GetValidate()
	if validate == nil {
		return errors.New("no active \"validate\" found")
	}

	err := validate.Struct(params)
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
