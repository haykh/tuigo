package obj

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Accessor interface {
	ID() int
	Data() interface{}
}

type Actor interface {
	Callback() tea.Msg
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

type ElementWithCallback struct {
	callback tea.Msg
}

func NewElementWithCallback(callback tea.Msg) ElementWithCallback {
	return ElementWithCallback{callback: callback}
}

func (e ElementWithCallback) Callback() tea.Msg {
	return e.callback
}
