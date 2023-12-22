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

type ElementWithID struct {
	id int
}

func NewElementWithID(id int) ElementWithID {
	return ElementWithID{id: id}
}

func (e ElementWithID) ID() int {
	return e.id
}
