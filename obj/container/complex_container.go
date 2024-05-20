package container

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/ui/theme"
	"github.com/haykh/tuigo/utils"
)

var _ obj.Element = (*ComplexContainer)(nil)
var _ AbstractComponent = (*ComplexContainer)(nil)
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
					if !comp.Hidden() || theme.DEBUG_MODE {
						el_views = append(el_views, comp.View(self.Focused()))
					}
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

func (cc ComplexContainer) HasVisible() bool {
	for _, component := range cc.components {
		if !component.Hidden() {
			return true
		}
	}
	return len(cc.components) == 0
}

// implementing Component
func (cc ComplexContainer) Hidden() bool {
	return cc.hidden || !cc.HasVisible()
}

func (cc ComplexContainer) Hide() Component {
	cc.hidden = true
	return cc
}

func (cc ComplexContainer) Unhide() Component {
	cc.hidden = false
	return cc
}

func (cc ComplexContainer) Enable() Component {
	cc.focusable = true
	return cc
}

func (cc ComplexContainer) Disable() Component {
	cc.focusable = false
	return cc
}

func (cc ComplexContainer) Focus() Component {
	return cc.FocusFromStart()
}

func (cc ComplexContainer) FocusFromStart() Component {
	if cc.Focusable() && !cc.Focused() && !cc.Hidden() {
		if len(cc.components) == 0 {
			cc.focused = true
			return cc
		}
		for c, component := range cc.components {
			if !component.Hidden() && component.Focusable() {
				cc.components[c] = component.FocusFromStart()
				if cc.components[c].Focused() {
					cc.focused = true
					return cc
				}
			}
		}
	}
	return cc
}

func (cc ComplexContainer) FocusFromEnd() Component {
	if cc.Focusable() && !cc.Focused() && !cc.Hidden() {
		if len(cc.components) == 0 {
			cc.focused = true
			return cc
		}
		for c := len(cc.components) - 1; c >= 0; c-- {
			component := cc.components[c]
			if !component.Hidden() && component.Focusable() {
				cc.components[c] = component.FocusFromEnd()
				if cc.components[c].Focused() {
					cc.focused = true
					return cc
				}
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
	if cc.Focusable() && !cc.Hidden() && cc.Focused() {
		last_foc_id := -1
		for c := len(cc.components) - 1; c >= 0; c-- {
			component := cc.components[c]
			if component.Focusable() && !component.Hidden() {
				if c > last_foc_id {
					last_foc_id = c
					break
				}
			}
		}
		focus_next := false
		for c, component := range cc.components {
			if focus_next {
				if component.Focusable() && !component.Hidden() {
					cc.components[c] = component.Focus()
					if cc.components[c].Focused() {
						return cc, utils.Callback(focusChangedMsg{})
					} else {
						return cc, utils.Callback(utils.FocusNextMsg{})
					}
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
	if cc.Focusable() && !cc.Hidden() && cc.Focused() {
		first_foc_id := len(cc.components) + 1
		for c, component := range cc.components {
			if c < first_foc_id {
				if component.Focusable() && !component.Hidden() {
					first_foc_id = c
					break
				}
			}
		}
		focus_prev := false
		for c := len(cc.components) - 1; c >= 0; c-- {
			component := cc.components[c]
			if focus_prev {
				if component.Focusable() && !component.Hidden() {
					cc.components[c] = component.FocusFromEnd()
					if cc.components[c].Focused() {
						return cc, utils.Callback(focusChangedMsg{})
					} else {
						return cc, utils.Callback(utils.FocusPrevMsg{})
					}
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

func (cc ComplexContainer) GetElementByID(id int) obj.Accessor {
	for _, component := range cc.components {
		if acc := component.GetElementByID(id); acc != nil {
			return acc
		}
	}
	return nil
}

func (cc ComplexContainer) GetContainerByID(id int) Component {
	for _, component := range cc.components {
		if comp := component.GetContainerByID(id); comp != nil {
			return comp
		}
	}
	return nil
}
