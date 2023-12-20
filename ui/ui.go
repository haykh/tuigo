package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/haykh/tuigo/ui/app"
	"github.com/haykh/tuigo/ui/button"
	"github.com/haykh/tuigo/ui/container"
	"github.com/haykh/tuigo/ui/debug"
	"github.com/haykh/tuigo/ui/pathinput"
	"github.com/haykh/tuigo/ui/radio"
	"github.com/haykh/tuigo/ui/selector"
	"github.com/haykh/tuigo/ui/text"
	"github.com/haykh/tuigo/utils"
)

// exposing all the views from the subpackages

func AppView(contents ...string) string {
	return app.View(contents...)
}

func ContainerView(focused bool, containerType utils.ContainerType, contents ...string) string {
	if containerType == utils.SimpleContainer {
		if len(contents) != 1 {
			panic("SimpleContainer must have exactly one element")
		}
	}
	return container.View(focused, containerType, contents...)
}

func ContainerControlView(controls ...string) string {
	return container.ControlView(controls...)
}

func DebugView(enabled bool, dbg string) string {
	return debug.View(enabled, dbg)
}

// components

func TextView(focused bool, txt string, texttype utils.TextType) string {
	return text.View(focused, txt, texttype)
}

func SelectorView(
	focused bool,
	multiselect bool,
	cursor int,
	items []string,
	selected, disabled map[string]struct{},
) string {
	return selector.View(focused, multiselect, cursor, items, selected, disabled)
}

func ButtonView(focused bool, label string, btntype utils.ButtonType) string {
	return button.View(focused, label, btntype)
}

func RadioView(focused bool, label string, state bool) string {
	return radio.View(focused, label, state)
}

func PathInputView(focused bool, ti textinput.Model) string {
	return pathinput.View(focused, ti)
}
