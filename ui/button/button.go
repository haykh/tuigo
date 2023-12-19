package button

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/haykh/tuigo/utils"
)

var (
	style        = lipgloss.NewStyle()
	focusedStyle = style.Copy().Underline(true)

	controlStyle        = lipgloss.NewStyle().Padding(0, 1).Margin(0, 1).Background(lipgloss.Color("240"))
	focusedControlStyle = controlStyle.Copy().Background(lipgloss.Color("205"))
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
		if focused {
			btnstyle = focusedStyle
		} else {
			btnstyle = style
		}
	}
	return btnstyle.Render(label)
}
