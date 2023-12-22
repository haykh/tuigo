package container

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

var _ obj.Element = (*SimpleContainer)(nil)
var _ Component = (*SimpleContainer)(nil)
var _ Wrapper = (*SimpleContainer)(nil)

type SimpleContainer struct {
	Container
	element obj.Element
}

func NewSimpleContainer(focusable bool, element obj.Element) SimpleContainer {
	render := func(self Component) string {
		var el_view string
		switch self := self.(type) {
		case SimpleContainer:
			el_view = self.Element().View(self.Focused())
		default:
			panic("unknown container type in ComplexContainer::render")
		}
		return ui.SimpleContainerView(self.Focused(), el_view)
	}
	return SimpleContainer{
		Container: Container{
			hidden:    false,
			focusable: focusable,
			focused:   false,
			render:    render,
		},
		element: element,
	}
}

// implementing Element
func (sc SimpleContainer) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	if sc.Focusable() && sc.Focused() {
		switch msg.(type) {
		case utils.FocusNextMsg:
			return sc.FocusNext()
		case utils.FocusPrevMsg:
			return sc.FocusPrev()
		}

		el, cmd := sc.element.Update(msg)
		sc.element = el
		return sc, cmd
	}
	return sc, nil
}

func (sc SimpleContainer) View(bool) string {
	return sc.render(sc)
}

// implementing Wrapper
func (sc SimpleContainer) Element() obj.Element {
	return sc.element
}

// implementing Component
func (sc SimpleContainer) Hide() Component {
	sc = sc.Blur().(SimpleContainer)
	sc.hidden = true
	return sc
}

func (sc SimpleContainer) Unhide() Component {
	sc.hidden = false
	return sc
}

func (sc SimpleContainer) Focus() Component {
	return sc.FocusFromStart()
}

func (sc SimpleContainer) FocusFromStart() Component {
	if sc.Focusable() && !sc.Focused() {
		sc.focused = true
	}
	return sc
}

func (sc SimpleContainer) FocusFromEnd() Component {
	if sc.Focusable() && !sc.Focused() {
		sc.focused = true
	}
	return sc
}

func (sc SimpleContainer) Blur() Component {
	if sc.Focusable() && sc.Focused() {
		sc.focused = false
	}
	return sc
}

func (sc SimpleContainer) FocusNext() (Component, tea.Cmd) {
	if sc.Focusable() {
		if sc.Focused() {
			// sc = sc.Blur().(SimpleContainer)
			return sc, utils.Callback(utils.FocusNextMsg{})
		} else {
			sc = sc.Focus().(SimpleContainer)
			return sc, utils.Callback(focusChangedMsg{})
		}
	}
	return sc, nil
}

func (sc SimpleContainer) FocusPrev() (Component, tea.Cmd) {
	if sc.Focusable() {
		if sc.Focused() {
			// sc = sc.Blur().(SimpleContainer)
			return sc, utils.Callback(utils.FocusPrevMsg{})
		} else {
			sc = sc.Focus().(SimpleContainer)
			return sc, utils.Callback(focusChangedMsg{})
		}
	}
	return sc, nil
}

func (sc SimpleContainer) GetElementByID(id int) (Component, obj.Accessor) {
	if acc, ok := sc.element.(obj.Accessor); ok && acc.ID() == id {
		return sc, acc
	}
	return nil, nil
}
