package radio

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	style         = lipgloss.NewStyle()
	focusedStyle  = style.Copy().Underline(true)
	stateOnStyle  = lipgloss.NewStyle()
	stateOffStyle = lipgloss.NewStyle()
)

var (
	stateOn  = "●"
	stateOff = " "
)

func View(label string, state, focused bool) string {
	var focusstyle lipgloss.Style
	var stateview string
	if focused {
		focusstyle = focusedStyle
	} else {
		focusstyle = style
	}
	if state {
		stateview = stateOnStyle.Render(stateOn)
	} else {
		stateview = stateOffStyle.Render(stateOff)
	}
	return fmt.Sprintf("(%s) %s", stateview, focusstyle.Render(label))
}
