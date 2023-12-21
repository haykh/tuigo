package app

import "github.com/charmbracelet/lipgloss"

var (
	style = lipgloss.NewStyle().MarginTop(1)
)

func View(contents ...string) string {
	return style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			contents...,
		),
	)
}
