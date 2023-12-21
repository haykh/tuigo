package radio

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/haykh/tuigo/ui/theme"
)

var (
	style         = theme.ElementStyle.Copy()
	focusedStyle  = lipgloss.NewStyle().Underline(true)
	unfocsedStyle = lipgloss.NewStyle().Foreground(theme.ColorDimmed)
	stateOnStyle  = lipgloss.NewStyle().Foreground(theme.ColorSuccess)
	stateOffStyle = lipgloss.NewStyle()
)

var (
	stateOn  = "●"
	stateOff = " "
)

func View(focused bool, label string, state bool) string {
	var focusstyle lipgloss.Style
	var stateview string
	if focused {
		focusstyle = focusedStyle
	} else {
		focusstyle = unfocsedStyle
	}
	lb := focusstyle.Copy().Underline(false).Render("(")
	rb := focusstyle.Copy().Underline(false).Render(")")
	if state {
		stateview = stateOnStyle.Render(stateOn)
	} else {
		stateview = stateOffStyle.Render(stateOff)
	}
	return style.Render(fmt.Sprintf("%s%s%s %s", lb, stateview, rb, focusstyle.Render(label)))
}
