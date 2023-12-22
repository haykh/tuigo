package container

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

var _ obj.Element = (*ComplexContainer)(nil)
var _ Component = (*ComplexContainer)(nil)
var _ Collection = (*ComplexContainer)(nil)

type ComplexContainer struct {
	Container
	containerType utils.ContainerType
	components    []Component
}

func NewComplexContainer(
	focusable bool,
	containerType utils.ContainerType,
	components ...Component,
) ComplexContainer {
	render := func(self Component) string {
		if self.Hidden() {
			return ""
		}
		el_views := []string{}
		switch self := self.(type) {
		case ComplexContainer:
			if len(self.Components()) == 0 {
				el_views = append(el_views, "")
			} else {
				for _, comp := range self.Components() {
					// if !comp.Hidden() {
					el_views = append(el_views, comp.View(self.Focused()))
					// }
				}
			}
		default:
			panic("unknown container type in ComplexContainer::render")
		}
		return ui.ComplexContainerView(self.Focused(), self.(Collection).Type(), el_views...)
	}
	return ComplexContainer{
		Container: Container{
			hidden:    false,
			focusable: focusable,
			focused:   false,
			render:    render,
		},
		containerType: containerType,
		components:    components,
	}
}

// implementing Element
func (cc ComplexContainer) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	// targeted message ignores focus
	switch msg := msg.(type) {
	case utils.TargetedMsg:
		var cmds []tea.Cmd
		for c, component := range cc.components {
			comp, cmd := component.Update(msg)
			cc.components[c] = comp.(Component)
			cmds = append(cmds, cmd)
		}
		return cc, tea.Batch(cmds...)
	}
	// messages that require focus
	if cc.Focusable() && cc.Focused() {
		switch msg.(type) {
		case utils.FocusNextMsg:
			return cc.FocusNext()
		case utils.FocusPrevMsg:
			return cc.FocusPrev()
		}

		var cmds []tea.Cmd
		for c, component := range cc.components {
			comp, cmd := component.Update(msg)
			cc.components[c] = comp.(Component)
			cmds = append(cmds, cmd)
		}
		return cc, tea.Batch(cmds...)
	}
	return cc, nil
}

func (cc ComplexContainer) View(bool) string {
	return cc.render(cc)
}

// implementing Collection
func (cc ComplexContainer) Type() utils.ContainerType {
	return cc.containerType
}

func (cc ComplexContainer) Components() []Component {
	return cc.components
}

func (cc ComplexContainer) AddComponents(components ...Component) Collection {
	cc.components = append(cc.components, components...)
	return cc
}

// implementing Component
func (cc ComplexContainer) Hide() Component {
	cc = cc.Blur().(ComplexContainer)
	cc.hidden = true
	return cc
}

func (cc ComplexContainer) Unhide() Component {
	cc.hidden = false
	return cc
}

func (cc ComplexContainer) Focus() Component {
	return cc.FocusFromStart()
}

func (cc ComplexContainer) FocusFromStart() Component {
	if cc.Focusable() && !cc.Focused() {
		cc.focused = true
		for c, component := range cc.components {
			cc.components[c] = component.FocusFromStart()
			if cc.components[c].Focused() {
				break
			}
		}
	}
	return cc
}

func (cc ComplexContainer) FocusFromEnd() Component {
	if cc.Focusable() && !cc.Focused() {
		cc.focused = true
		for c := len(cc.components) - 1; c >= 0; c-- {
			component := cc.components[c]
			cc.components[c] = component.FocusFromEnd()
			if cc.components[c].Focused() {
				break
			}
		}
	}
	return cc
}

func (cc ComplexContainer) Blur() Component {
	if cc.Focusable() && cc.Focused() {
		cc.focused = false
		for c, component := range cc.components {
			cc.components[c] = component.Blur()
		}
	}
	return cc
}

func (cc ComplexContainer) FocusNext() (Component, tea.Cmd) {
	if cc.Focusable() && cc.Focused() {
		last_foc_id := -1
		for c := len(cc.components) - 1; c >= 0; c-- {
			component := cc.components[c]
			if component.Focusable() {
				if c > last_foc_id {
					last_foc_id = c
					break
				}
			}
		}
		focus_next := false
		for c, component := range cc.components {
			if focus_next {
				if component.Focusable() {
					cc.components[c] = component.Focus()
					return cc, utils.Callback(focusChangedMsg{})
				}
			} else if component.Focused() {
				comp, cmd := component.FocusNext()
				cc.components[c] = comp
				if cmd != nil {
					switch cmd().(type) {
					case utils.FocusNextMsg:
						if c < last_foc_id {
							cc.components[c] = comp.Blur()
						}
						focus_next = true
					case focusChangedMsg:
						return cc, cmd
					default:
						panic("unknown cmd")
					}
				}
			}
		}
		return cc, utils.Callback(utils.FocusNextMsg{})
	}
	return cc, nil
}

func (cc ComplexContainer) FocusPrev() (Component, tea.Cmd) {
	if cc.Focusable() && cc.Focused() {
		first_foc_id := len(cc.components) + 1
		for c, component := range cc.components {
			if c < first_foc_id {
				if component.Focusable() {
					first_foc_id = c
					break
				}
			}
		}
		focus_prev := false
		for c := len(cc.components) - 1; c >= 0; c-- {
			component := cc.components[c]
			if focus_prev {
				if component.Focusable() {
					cc.components[c] = component.FocusFromEnd()
					return cc, utils.Callback(focusChangedMsg{})
				}
			} else if component.Focused() {
				el, cmd := component.FocusPrev()
				cc.components[c] = el
				if cmd != nil {
					switch cmd().(type) {
					case utils.FocusPrevMsg:
						if c > first_foc_id {
							cc.components[c] = el.Blur()
						}
						focus_prev = true
					case focusChangedMsg:
						return cc, cmd
					default:
						panic("unknown cmd")
					}
				}
			}
		}
		return cc, utils.Callback(utils.FocusPrevMsg{})
	}
	return cc, nil
}

func (cc ComplexContainer) GetElementByID(id int) (Component, obj.Accessor) {
	for _, component := range cc.components {
		if comp, acc := component.GetElementByID(id); comp != nil {
			return comp, acc
		}
	}
	return nil, nil
}
