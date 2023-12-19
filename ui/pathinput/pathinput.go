package pathinput

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/haykh/tuigo/ui/theme"
)

var (
	style                = lipgloss.NewStyle().MarginLeft(2).MarginBottom(1)
	focusedStyle         = lipgloss.NewStyle()
	unfocusedStyle       = lipgloss.NewStyle().Foreground(theme.ColorDimmed)
	promptStyle          = lipgloss.NewStyle()
	unfocusedPromptStyle = promptStyle.Copy().Foreground(theme.ColorDimmed)
	focusedPromptStyle   = promptStyle.Copy().Foreground(theme.ColorSpecial)
)

func View(ti textinput.Model) string {
	if ti.Focused() {
		ti.PromptStyle = focusedPromptStyle
		ti.TextStyle = focusedStyle
	} else {
		ti.PromptStyle = unfocusedPromptStyle
		ti.TextStyle = unfocusedStyle
	}
	ti.Prompt += ": "
	return style.Render(ti.View())
}
