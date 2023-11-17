package i18n

import "github.com/kordar/goi18n"

func InitI18n(dir string, lang string) {
	handle := goi18n.NewIniHandler(dir, lang)

	// 配置 yaml 支持
	goi18n.InitLang(handle)

	// 配置response国际化支持
	initResponse()
}
