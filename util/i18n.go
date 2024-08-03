package util

import "github.com/kordar/gocfg"

func I18n(locale string, section string) string {
	return gocfg.GetSectionValue(locale, section, "language")
}
