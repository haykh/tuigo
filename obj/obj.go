package obj

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Element interface {
	View(bool) string
	Update(tea.Msg) (Element, tea.Cmd)
}

type Collection interface {
	Element
	Elements() []Element
	AddElements(...Element) Collection
	Focusable() bool
	Focused() bool
	Focus() Collection
	FocusFromStart() Collection
	FocusFromEnd() Collection
	Blur() Collection
	FocusNext() (Collection, tea.Cmd)
	FocusPrev() (Collection, tea.Cmd)
}
