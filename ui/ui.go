package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/haykh/tuigo/ui/base"
	"github.com/haykh/tuigo/ui/button"
	"github.com/haykh/tuigo/ui/debug"
	"github.com/haykh/tuigo/ui/field"
	"github.com/haykh/tuigo/ui/multiselect"
	"github.com/haykh/tuigo/ui/pathinput"
	"github.com/haykh/tuigo/ui/radio"
	"github.com/haykh/tuigo/utils"
)

// exposing all the views from the subpackages

func BaseView(contents ...string) string {
	return base.View(contents...)
}

func FieldView(label string, contents ...string) string {
	return field.View(label, contents...)
}

func FieldControlView(controls ...string) string {
	return field.ControlView(controls...)
}

func DebugView(enabled bool, dbg string) string {
	return debug.View(enabled, dbg)
}

// components

func MultiSelectView(
	cursor int,
	items []string,
	selected, disabled map[string]struct{},
	focused bool,
) string {
	return multiselect.View(cursor, items, selected, disabled, focused)
}

func ButtonView(label string, focused bool, btntype utils.ButtonType) string {
	return button.View(label, focused, btntype)
}

func RadioView(label string, state, focused bool) string {
	return radio.View(label, state, focused)
}

func PathInputView(ti textinput.Model) string {
	return pathinput.View(ti)
}
