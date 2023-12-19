package button

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/haykh/tuigo/ui/theme"
	"github.com/haykh/tuigo/utils"
)

var (
	style          = lipgloss.NewStyle().MarginLeft(2).MarginBottom(1)
	focusedStyle   = style.Copy().Underline(true)
	unfocusedStyle = style.Copy().Foreground(theme.ColorDimmed)

	controlStyle = lipgloss.NewStyle().Padding(0, 1).Margin(0, 1).
			Background(theme.ColorDimmed)
	focusedControlStyle = controlStyle.Copy().Background(theme.ColorAccent)
)

func View(label string, focused bool, btntype utils.ButtonType) string {
	var btnstyle lipgloss.Style
	if btntype == utils.ControlBtn {
		if focused {
			btnstyle = focusedControlStyle
		} else {
			btnstyle = controlStyle
		}
	} else {
		label = "[ " + label + " ]"
		if focused {
			btnstyle = focusedStyle
		} else {
			btnstyle = unfocusedStyle
		}
	}
	return btnstyle.Render(label)
}
