package container

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/haykh/tuigo/ui/theme"
	"github.com/haykh/tuigo/utils"
)

var (
	style          = theme.ContainerStyle.Copy()
	focusedStyle   = style.Copy()
	unfocusedStyle = style.Copy()
)

func ViewComplex(focused bool, containerType utils.ContainerType, contents ...string) string {
	var focus_style lipgloss.Style
	if focused {
		focus_style = focusedStyle
		if theme.DEBUG_MODE {
			focus_style = focus_style.Copy().Border(lipgloss.NormalBorder())
		}
	} else {
		focus_style = unfocusedStyle
	}
	if containerType == utils.VerticalContainer {
		return focus_style.Render(lipgloss.JoinVertical(
			lipgloss.Left,
			contents...,
		))
	} else {
		return focus_style.Render(lipgloss.JoinHorizontal(
			lipgloss.Center,
			contents...,
		))
	}
}

func ViewSimple(focused bool, content string) string {
	var focus_style lipgloss.Style
	if focused {
		focus_style = focusedStyle
		if theme.DEBUG_MODE {
			focus_style = focus_style.Copy().Border(lipgloss.NormalBorder())
		}
	} else {
		focus_style = unfocusedStyle
	}
	return focus_style.Render(content)
}

func ControlView(controls ...string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, controls...)
}
