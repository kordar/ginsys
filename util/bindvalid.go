package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kordar/govalidator"
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

	// TODO 通过反射执行验证暂关闭
	//refParams := reflect.ValueOf(params) // 需要传入指针，后面再解析
	//validMethod := refParams.MethodByName("Valid")
	//if validMethod.IsValid() {
	//	v := validMethod.Call(make([]reflect.Value, 0))
	//	if e := v[0].Interface(); e != nil {
	//		return e.(error)
	//	}
	//}
	return nil
}
