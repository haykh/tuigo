package debug

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/haykh/tuigo/ui/theme"
	"golang.org/x/term"
)

var (
	style = lipgloss.NewStyle().
		MarginTop(1).
		Border(lipgloss.NormalBorder(), true, false, false, false).
		BorderForeground(theme.ColorDimmed)
	labelStyle = lipgloss.NewStyle().
			Background(theme.ColorSpecial).
			PaddingLeft(1).PaddingRight(1)
	msgStyle = lipgloss.NewStyle().PaddingLeft(1).Background(theme.ColorDimmed)
)

func render(msg string) string {
	w, _, _ := term.GetSize(int(os.Stdout.Fd()))
	label := labelStyle.Render("DEBUG:")
	message := msgStyle.Width(w - 1 - lipgloss.Width(label)).Render(msg)
	return style.Width(w - 1).Render(lipgloss.JoinHorizontal(
		lipgloss.Center,
		label,
		message,
	))
}

func View(enabled bool, dbg string) string {
	if enabled {
		return render(dbg)
	}
	return ""
}
