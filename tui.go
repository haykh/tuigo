package tuigo

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/component"
	"github.com/haykh/tuigo/component/button"
	"github.com/haykh/tuigo/component/input"
	"github.com/haykh/tuigo/component/radio"
	"github.com/haykh/tuigo/component/selector"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

type State interface {
	Next() State
	Prev() State
}

type Actor interface {
	Action() tea.Msg
}

type Messenger interface {
	Message() string
}

// Container

type Container struct {
	elements  []component.Viewer
	label     string
	ncontrols int
}

func NewContainer(label string, isFirstContainer bool, isLastContainer bool) Container {
	f := Container{
		elements:  []component.Viewer{},
		label:     label,
		ncontrols: 0,
	}
	if !isFirstContainer {
		f.ncontrols++
		prev := button.New("< back", utils.ControlBtn, utils.PrevStateMsg{})
		f = f.AddElement(&prev)
	}
	if !isLastContainer {
		f.ncontrols++
		next := button.New("next >", utils.ControlBtn, utils.NextStateMsg{})
		f = f.AddElement(&next)
	}
	f.FocusLast()
	return f
}

func (f Container) Init() tea.Cmd {
	return nil
}

func (f Container) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// process switching focus
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Tab):
			f.FocusNext()
			return f, nil
		case key.Matches(msg, keys.Keys.ShiftTab):
			f.FocusPrev()
			return f, nil
		}
	}

	// process individual elements
	var cmds []tea.Cmd
	for e, element := range f.elements {
		switch el := element.(type) {
		case component.Updater:
			if el.Focused() {
				el, cmd := el.Update(msg)
				f.elements[e] = el.(component.Viewer)
				cmds = append(cmds, cmd)
			}
		}
	}
	return f, tea.Batch(cmds...)
}

func (f Container) View() string {
	viewers := f.GetViewersWithoutControls()
	controls := f.GetControls()
	elementViews := []string{}
	for _, element := range viewers {
		elementViews = append(elementViews, element.View())
	}
	controlViews := []string{}
	for _, control := range controls {
		controlViews = append(controlViews, control.View())
	}
	return ui.ContainerView(
		f.label,
		append(
			elementViews,
			ui.ContainerControlView(controlViews...),
		)...,
	)
}

func (f Container) AddElement(element component.Viewer) Container {
	f.elements = append(f.elements, element)
	f.FocusNext()
	f.FocusPrev()
	return f
}

func (f Container) GetViewersWithoutControls() []component.Viewer {
	return f.elements[f.ncontrols:]
}

func (f Container) GetControls() []component.Viewer {
	return f.elements[:f.ncontrols]
}

func (f Container) GetFocusers() []component.Focuser {
	var focusers []component.Focuser
	elements := append(f.GetViewersWithoutControls(), f.GetControls()...)
	for _, element := range elements {
		switch el := element.(type) {
		case component.Focuser:
			focusers = append(focusers, el)
		}
	}
	return focusers
}

func (f *Container) FocusNext() {
	focusers := f.GetFocusers()
	for e, focuser := range focusers {
		if focuser.Focused() {
			focuser.Blur()
			if e+1 < len(focusers) {
				focusers[e+1].Focus()
				return
			} else {
				focusers[0].Focus()
				return
			}
		}
	}
	if len(focusers) > 0 {
		focusers[0].Focus()
	}
}

func (f *Container) FocusPrev() {
	focusers := f.GetFocusers()
	for e, focuser := range focusers {
		if focuser.Focused() {
			focuser.Blur()
			if e-1 >= 0 {
				focusers[e-1].Focus()
				return
			} else {
				focusers[len(focusers)-1].Focus()
				return
			}
		}
	}
	if len(focusers) > 0 {
		focusers[0].Focus()
	}
}

func (f *Container) FocusLast() {
	focusers := f.GetFocusers()
	if len(focusers) > 0 {
		for _, f := range focusers {
			f.Blur()
		}
		focusers[len(focusers)-1].Focus()
	}
}

func (f *Container) FocusFirst() {
	focusers := f.GetFocusers()
	if len(focusers) > 0 {
		for _, foc := range focusers {
			foc.Blur()
		}
		focusers[0].Focus()
	}
}

// Components

func NewRadio(label string) radio.Model {
	return radio.New(label)
}

func NewInput(label, def, placeholder string, inputtype utils.InputType) input.Model {
	return input.New(label, def, placeholder, inputtype)
}

func NewButton(label string, btnType utils.ButtonType, msg tea.Msg) button.Model {
	return button.New(label, btnType, msg)
}

func NewSelector(options []string, multiselect bool) selector.Model {
	return selector.New(options, multiselect)
}
