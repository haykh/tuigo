package field

import "github.com/charmbracelet/lipgloss"

func View(label string, contents ...string) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		append([]string{label}, contents...)...,
	)
}

func ControlView(controls ...string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, controls...)
}
