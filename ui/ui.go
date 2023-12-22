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
	"github.com/haykh/tuigo/ui/theme"
	"github.com/haykh/tuigo/utils"
)

// exposing all the views from the subpackages

func AppView(contents ...string) string {
	return app.View(contents...)
}

func ComplexContainerView(focused bool, containerType utils.ContainerType, contents ...string) string {
	for i, content := range contents {
		if content == "" {
			if theme.DEBUG_MODE {
				contents[i] = "dummy"
			}
		}
	}
	return container.ViewComplex(focused, containerType, contents...)
}

func SimpleContainerView(focused bool, content string) string {
	if content == "" {
		if theme.DEBUG_MODE {
			content = "dummy"
		}
	}
	return container.ViewSimple(focused, content)
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
