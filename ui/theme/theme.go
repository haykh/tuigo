package theme

import "github.com/charmbracelet/lipgloss"

var (
	DEBUG_STYLE = false
)

var (
	ColorDimmed  = lipgloss.AdaptiveColor{Light: "#272a31", Dark: "#31353d"}
	ColorSuccess = lipgloss.AdaptiveColor{Light: "#3eb86d", Dark: "#67cc8e"}
	ColorLight   = lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#f2f2f2"}
	ColorSpecial = lipgloss.AdaptiveColor{Light: "#ee9723", Dark: "#F3B562"}
	ColorAccent  = lipgloss.AdaptiveColor{Light: "#ea2323", Dark: "#F06060"}
)

var (
	ElementStyle   = lipgloss.NewStyle().MarginLeft(2).MarginBottom(1)
	ContainerStyle = lipgloss.NewStyle().PaddingRight(1)
)
