package util

import (
	"github.com/kordar/gocfg"
	"github.com/spf13/cast"
)

func I18n(locale string, section string) string {
	return gocfg.GetSectionValueM(section, "language", locale)
}

func I18ns(locale string, section string) map[string]interface{} {
	v := gocfg.GetM(section+".language", locale)
	return cast.ToStringMap(v)
}
