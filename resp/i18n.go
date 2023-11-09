package resp

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// ---------- i18n success --------------

type SuccessJsonI18n struct {
	I18nMessage func(message string, t string, c *gin.Context) string
	SuccessJson
}

func (s SuccessJsonI18n) Result(c *gin.Context, message interface{}, data interface{}, count int64) {
	if value, ok := message.(string); ok && value != "" {
		if tmessage := s.I18nMessage(value, "success", c); tmessage != "" {
			jsonCall(c, s.HttpStatus(), Code("success"), tmessage, data, count)
			return
		}
	}
	jsonCall(c, s.HttpStatus(), success, s.I18nMessage("success", "success", c), data, count)
}

// ----------- i18n error --------------

type ErrorJsonI18n struct {
	GetTrans    func(c *gin.Context) (trans ut.Translator, found bool)
	I18nMessage func(message string, t string, c *gin.Context) string
	ErrorJson
}

func (e ErrorJsonI18n) Result(c *gin.Context, message interface{}, data interface{}, count int64) {

	// 处理文本message
	if err, ok := message.(string); ok && err != "" {
		if tmessage := e.I18nMessage(err, "error", c); tmessage != "" {
			jsonCall(c, e.HttpStatus(), Code("error"), tmessage, data, count)
		} else {
			jsonCall(c, e.HttpStatus(), Code("error"), err, data, count)
		}
		return
	}

	// 处理validate error
	if err, ok := message.(validator.ValidationErrors); ok {
		nMessage := e.I18nMessage("error.valid", "error", c)
		if trans, found := e.GetTrans(c); found {
			jsonCall(c, e.HttpStatus(), Code("valid"), nMessage, err.Translate(trans), count)
		} else {
			jsonCall(c, e.HttpStatus(), Code("valid"), nMessage, data, count)
		}
		return
	}

	// 处理error
	if value, ok := message.(error); ok {

		if tmessage := e.I18nMessage(value.Error(), "error", c); tmessage != "" {
			jsonCall(c, e.HttpStatus(), fail, tmessage, data, count)
		} else {
			jsonCall(c, e.HttpStatus(), fail, value.Error(), data, count)
		}

		return
	}

	jsonCall(c, e.HttpStatus(), fail, e.I18nMessage("error", "error", c), data, count)
}

// ----------- i18n output -------------------

type OutputResponseI18n struct {
	I18nMessage func(message string, t string, c *gin.Context) string
	OutputResponse
}

func (o OutputResponseI18n) Result(c *gin.Context, message interface{}, data interface{}, count int64) {
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

		msg := "output.success"
		if tmessage, found := message.(string); found && tmessage != "" {
			msg = tmessage
			if tmessage2 := o.I18nMessage(tmessage, "success", c); tmessage2 != "" {
				msg = tmessage2
			}
		}

		c.JSON(o.HttpStatus(), map[string]interface{}{
			"code": success, "message": o.I18nMessage(msg, "success", c), "data": value.Data(), "params": value.Params(),
		})

		return
	}

	jsonCall(c, o.HttpStatus(), fail, o.I18nMessage("output.fail", "error", c), nil, -1)

}

// ------------ i18n unauthorized --------------

type UnauthorizedJsonI18n struct {
	I18nMessage func(message string, t string, c *gin.Context) string
	UnauthorizedResponse
}

func (s UnauthorizedJsonI18n) Result(c *gin.Context, message interface{}, data interface{}, count int64) {
	if value, ok := message.(string); ok && value != "" {
		if tmessage := s.I18nMessage(value, "unauthorized", c); tmessage != "" {
			jsonCall(c, s.HttpStatus(), Code("unauthorized"), tmessage, data, count)
			return
		}
	}
	jsonCall(c, s.HttpStatus(), Code("unauthorized"), s.I18nMessage("unauthorized", "unauthorized", c), data, count)
}
