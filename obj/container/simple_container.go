package container

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/obj/components/button"
	"github.com/haykh/tuigo/obj/components/input"
	"github.com/haykh/tuigo/obj/components/radio"
	"github.com/haykh/tuigo/obj/components/selector"
	"github.com/haykh/tuigo/obj/components/text"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

var _ obj.Element = (*SimpleContainer)(nil)
var _ AbstractComponent = (*SimpleContainer)(nil)
var _ Component = (*SimpleContainer)(nil)
var _ Wrapper = (*SimpleContainer)(nil)

type SimpleContainer struct {
	Container
	element obj.Element
}

func NewSimpleContainer(focusable bool, element obj.Element) SimpleContainer {
	render := func(self Component) string {
		if self.Hidden() {
			return ""
		}
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
	// targeted message ignores focus
	switch msg := msg.(type) {
	case utils.TargetedMsg:
		if acc, ok := sc.element.(obj.Accessor); ok {
			if acc.ID() == msg.ID() {
				action := msg.Action().(func(Wrapper, obj.Accessor) (Wrapper, obj.Accessor))
				newsc, acc := action(sc, acc)
				sc = newsc.(SimpleContainer)
				sc.element = acc.(obj.Element)
				return sc, nil
			}
		}
		return sc, nil
	}
	// messages that require focus
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

/**
 * Set works differently for different components
 *   Button: Set label
 *   Input: Set value
 *   Radio: Toggle
 *   Selector: Toggle specific option
 *   Text: Set text
 */
func (sc SimpleContainer) Set(args ...interface{}) Wrapper {
	switch el := sc.Element().(type) {
	case button.Button:
		if len(args) == 1 {
			arg := args[0]
			if str, ok := arg.(string); ok {
				el = el.Set(str)
				sc.element = el
				return sc
			} else {
				panic("invalid argument type for Button.Set")
			}
		} else {
			panic("invalid number of arguments for Button.Set")
		}
	case input.Input:
		if len(args) == 1 {
			arg := args[0]
			if str, ok := arg.(string); ok {
				el = el.Set(str)
				sc.element = el
				return sc
			} else {
				panic("invalid argument type for Input.Set")
			}
		} else {
			panic("invalid number of arguments for Input.Set")
		}
	case radio.Radio:
		if len(args) != 0 {
			panic("invalid number of arguments for Radio.Set")
		} else {
			el = el.Toggle()
			sc.element = el
			return sc
		}
	case selector.Selector:
		if len(args) == 1 {
			arg := args[0]
			if str, ok := arg.(string); ok {
				el = el.ToggleSpecific(str)
				sc.element = el
				return sc
			} else {
				panic("invalid argument type for Selector.ToggleSpecific")
			}
		} else {
			panic("invalid number of arguments for Selector.ToggleSpecific")
		}
	case text.Text:
		if len(args) == 1 {
			arg := args[0]
			if str, ok := arg.(string); ok {
				el = el.Set(str)
				sc.element = el
				return sc
			} else {
				panic("invalid argument type for Text.Set")
			}
		} else {
			panic("invalid number of arguments for Text.Set")
		}
	}
	return sc
}

// implementing Component
func (sc SimpleContainer) Hide() Component {
	sc.hidden = true
	return sc
}

func (sc SimpleContainer) Unhide() Component {
	sc.hidden = false
	return sc
}

func (sc SimpleContainer) Enable() Component {
	sc.focusable = true
	return sc
}

func (sc SimpleContainer) Disable() Component {
	sc.focusable = false
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
