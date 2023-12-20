package container

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

var _ obj.Element = (*Container)(nil)
var _ obj.Collection = (*Container)(nil)

type Container struct {
	focusable     bool
	focused       bool
	containerType utils.ContainerType
	elements      []obj.Element
	render        func(Container) string
}

func NewSimpleContainer(focusable bool, elements ...obj.Element) Container {
	return NewContainer(focusable, utils.SimpleContainer, elements...)
}

func NewContainer(focusable bool, containerType utils.ContainerType, elements ...obj.Element) Container {
	render := func(self Container) string {
		el_views := []string{}
		for _, element := range self.elements {
			el_views = append(el_views, element.View(self.Focused()))
		}
		return ui.ContainerView(self.focused, self.containerType, el_views...)
	}
	return Container{
		focusable:     focusable,
		focused:       false,
		containerType: containerType,
		elements:      elements,
		render:        render,
	}
}

func (c Container) View(bool) string {
	return c.render(c)
}

func (c Container) Containers() []obj.Element {
	return c.elements
}

func (c Container) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	if c.Focusable() && c.Focused() {
		switch msg.(type) {
		case utils.FocusNextMsg:
			return c.FocusNext()
		case utils.FocusPrevMsg:
			return c.FocusPrev()
		}

		var cmds []tea.Cmd
		for e, element := range c.elements {
			el, cmd := element.Update(msg)
			c.elements[e] = el
			cmds = append(cmds, cmd)
		}
		return c, tea.Batch(cmds...)
	}
	return c, nil
}

func (c Container) Focusable() bool {
	return c.focusable
}

func (c Container) Focused() bool {
	return c.focused
}

func (c Container) Focus() obj.Collection {
	return c.FocusFromStart()
}

func (c Container) FocusFromStart() obj.Collection {
	if c.Focusable() && !c.Focused() {
		c.focused = true
		for e, element := range c.elements {
			if el, ok := element.(obj.Collection); ok {
				c.elements[e] = el.Focus()
				if c.elements[e].(obj.Collection).Focused() {
					break
				}
			}
		}
	}
	return c
}

func (c Container) FocusFromEnd() obj.Collection {
	if c.Focusable() && !c.Focused() {
		c.focused = true
		for e := len(c.elements) - 1; e >= 0; e-- {
			element := c.elements[e]
			if el, ok := element.(obj.Collection); ok {
				c.elements[e] = el.Focus()
				if c.elements[e].(obj.Collection).Focused() {
					break
				}
			}
		}
	}
	return c
}

func (c Container) Blur() obj.Collection {
	if c.Focusable() && c.Focused() {
		c.focused = false
		for e, element := range c.elements {
			if el, ok := element.(obj.Collection); ok {
				c.elements[e] = el.Blur()
			}
		}
	}
	return c
}

func (c Container) FocusNext() (obj.Collection, tea.Cmd) {
	type FocusChangedMsg struct{}
	if c.Focusable() && c.Focused() {
		last_col_id := -1
		for e, element := range c.elements {
			if _, ok := element.(obj.Collection); ok {
				if e > last_col_id {
					last_col_id = e
				}
			}
		}
		focus_next := false
		for e, element := range c.elements {
			if el, ok := element.(obj.Collection); ok {
				if focus_next {
					if el.Focusable() {
						c.elements[e] = el.Focus()
						return c, utils.Callback(FocusChangedMsg{})
					}
				} else if el.Focused() {
					el, cmd := el.FocusNext()
					c.elements[e] = el
					if cmd != nil {
						switch cmd().(type) {
						case utils.FocusNextMsg:
							if e < last_col_id {
								c.elements[e] = el.Blur()
							}
							focus_next = true
						case FocusChangedMsg:
							return c, nil
						default:
							panic("unknown cmd")
						}
					}
				}
			}
		}
		return c, utils.Callback(utils.FocusNextMsg{})
	}
	return c, nil
}

func (c Container) FocusPrev() (obj.Collection, tea.Cmd) {
	type FocusChangedMsg struct{}
	if c.Focusable() && c.Focused() {
		first_col_id := len(c.elements) + 1
		for e, element := range c.elements {
			if _, ok := element.(obj.Collection); ok {
				if e < first_col_id {
					first_col_id = e
				}
			}
		}
		focus_prev := false
		for e := len(c.elements) - 1; e >= 0; e-- {
			element := c.elements[e]
			if el, ok := element.(obj.Collection); ok {
				if focus_prev {
					if el.Focusable() {
						c.elements[e] = el.FocusFromEnd()
						return c, utils.Callback(FocusChangedMsg{})
					}
				} else if el.Focused() {
					el, cmd := el.FocusPrev()
					c.elements[e] = el
					if cmd != nil {
						switch cmd().(type) {
						case utils.FocusPrevMsg:
							if e > first_col_id {
								c.elements[e] = el.Blur()
							}
							focus_prev = true
						case FocusChangedMsg:
							return c, nil
						default:
							panic("unknown cmd")
						}
					}
				}
			}
		}
		return c, utils.Callback(utils.FocusPrevMsg{})
	}
	return c, nil
}
