package ginsys

import "github.com/kordar/goi18n"

type ValidateError interface {
	Translate(lang string, message string) string
	Error() string
}

func New(text string) error {
	return &errorString{text, "response.errors", "ini"}
}

func New2(text string, section string, t string) error {
	return &errorString{text, section, t}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s       string
	section string
	t       string
}

func (e *errorString) Error() string {
	return e.s
}

func (e *errorString) Translate(lang string, message string) string {
	if s := goi18n.GetSectionValue(lang, e.section, message, e.t).(string); s == "" {
		return message
	} else {
		return s
	}
}
