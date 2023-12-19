package debug

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	labelStyle = lipgloss.NewStyle()
	msgStyle   = lipgloss.NewStyle()
)

func render(msg string) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		labelStyle.Render("DEBUG: "),
		msgStyle.Render(msg),
	)
}

func View(enabled bool, dbg string) string {
	if enabled {
		return render(dbg)
	}
	return ""
}
