package base

import "github.com/charmbracelet/lipgloss"

func View(contents ...string) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		contents...,
	)
}
