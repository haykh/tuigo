package container

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/haykh/tuigo/ui/theme"
	"github.com/haykh/tuigo/utils"
)

var (
	style          = theme.ContainerStyle.Copy()
	focusedStyle   = style.Copy().Border(lipgloss.RoundedBorder())
	unfocusedStyle = style.Copy()
)

func View(focused bool, containerType utils.ContainerType, contents ...string) string {
	var focus_style lipgloss.Style
	if focused {
		focus_style = focusedStyle
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
			lipgloss.Top,
			contents...,
		))
	}
}

func ControlView(controls ...string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, controls...)
}
