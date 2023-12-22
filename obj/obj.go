package obj

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Accessor interface {
	ID() int
	Data() interface{}
}

type Element interface {
	View(bool) string
	Update(tea.Msg) (Element, tea.Cmd)
}

type Collection interface {
	Element
	Elements() []Element
	AddElements(...Element) Collection
	Hidden() bool
	Hide() Collection
	Unhide() Collection
	Focusable() bool
	Focused() bool
	Focus() Collection
	FocusFromStart() Collection
	FocusFromEnd() Collection
	Blur() Collection
	FocusNext() (Collection, tea.Cmd)
	FocusPrev() (Collection, tea.Cmd)
	GetElementByID(int) Accessor
}

type ElementWithID struct {
	id int
}

func NewElementWithID(id int) ElementWithID {
	return ElementWithID{id: id}
}

func (e ElementWithID) ID() int {
	return e.id
}
