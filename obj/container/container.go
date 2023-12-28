package container

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/utils"
)

type AbstractComponent interface {
	Hidden() bool
	Focusable() bool
	Focused() bool
	Disabled() bool
}

type Component interface {
	obj.Element
	AbstractComponent
	Hide() Component
	Unhide() Component
	Enable() Component
	Disable() Component
	Focus() Component
	FocusFromStart() Component
	FocusFromEnd() Component
	Blur() Component
	FocusNext() (Component, tea.Cmd)
	FocusPrev() (Component, tea.Cmd)
	GetElementByID(int) (Component, obj.Accessor)
}

type Collection interface {
	Component
	Type() utils.ContainerType
	Components() []Component
	AddComponents(...Component) Collection
}

type Wrapper interface {
	Component
	Element() obj.Element
	Set(...interface{}) Wrapper
}

var _ AbstractComponent = (*Container)(nil)

type Container struct {
	hidden    bool
	focusable bool
	focused   bool
	render    func(Component) string
}

func (c Container) Hidden() bool {
	return c.hidden
}

func (c Container) Focusable() bool {
	return c.focusable && !c.hidden
}

func (c Container) Focused() bool {
	return c.focused
}

func (c Container) Disabled() bool {
	return !c.focusable
}

type focusChangedMsg struct{}
