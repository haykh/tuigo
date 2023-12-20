package container

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/haykh/tuigo/ui/theme"
	"golang.org/x/term"
)

var (
	labelStyle = lipgloss.NewStyle().Foreground(theme.ColorAccent).MarginBottom(1)
)

func View(label string, contents ...string) string {
	w, _, _ := term.GetSize(int(os.Stdout.Fd()))
	labelview := labelStyle.Width(w - 1).Align(lipgloss.Center).Render(label)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		append([]string{labelview}, contents...)...,
	)
}

func ControlView(controls ...string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, controls...)
}
