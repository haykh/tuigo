package pathinput

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

var (
	style              = lipgloss.NewStyle()
	focusedStyle       = style.Copy()
	promptStyle        = lipgloss.NewStyle()
	focusedPromptStyle = promptStyle.Copy()
)

func View(ti textinput.Model) string {
	if ti.Focused() {
		ti.PromptStyle = focusedPromptStyle
		ti.TextStyle = focusedStyle
	} else {
		ti.PromptStyle = promptStyle
		ti.TextStyle = style
	}
	ti.Prompt += ": "
	return ti.View()
}
