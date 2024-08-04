package util

import (
	"github.com/kordar/gocfg"
	"github.com/spf13/cast"
)

func I18n(locale string, section string) string {
	return gocfg.GetSectionValue(locale, section, "language")
}

func I18ns(locale string, section string) map[string]interface{} {
	v := gocfg.Get(locale, section, "language")
	return cast.ToStringMap(v)
}
