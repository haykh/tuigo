package button

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/haykh/tuigo/ui/theme"
	"github.com/haykh/tuigo/utils"
)

var (
	style          = theme.ElementStyle.Copy()
	focusedStyle   = style.Copy().Underline(true)
	unfocusedStyle = style.Copy().Foreground(theme.ColorDimmed)

	controlStyle = lipgloss.NewStyle().Padding(0, 1).Margin(0, 1).
			Background(theme.ColorDimmed).Background(theme.ColorAccent)
	focusedControlStyle = controlStyle.Copy().Background(theme.ColorAccent2)
)

func View(focused bool, label string, btntype utils.ButtonType) string {
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
