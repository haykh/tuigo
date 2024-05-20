package theme

import "github.com/charmbracelet/lipgloss"

var (
	DEBUG_MODE = false
)

var (
	ColorDimmed  = lipgloss.AdaptiveColor{Light: "8", Dark: "8"}
	ColorSuccess = lipgloss.AdaptiveColor{Light: "2", Dark: "2"}
	ColorLight   = lipgloss.AdaptiveColor{Light: "0", Dark: "7"}
	ColorSpecial = lipgloss.AdaptiveColor{Light: "3", Dark: "3"}
	ColorAccent  = lipgloss.AdaptiveColor{Light: "1", Dark: "1"}
	ColorAccent2 = lipgloss.AdaptiveColor{Light: "6", Dark: "6"}
)

var (
	ElementStyle   = lipgloss.NewStyle().MarginLeft(2).MarginBottom(1)
	ContainerStyle = lipgloss.NewStyle().PaddingRight(1)
)
