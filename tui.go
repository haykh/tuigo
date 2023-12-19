package tuigo

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/component"
	"github.com/haykh/tuigo/component/button"
	"github.com/haykh/tuigo/component/input"
	"github.com/haykh/tuigo/component/radio"
	"github.com/haykh/tuigo/component/selector"
	"github.com/haykh/tuigo/debug"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

type State interface {
	Label() string
	Next() State
	Prev() State
}

type Actor interface {
	Action() tea.Msg
}

type Messenger interface {
	Message() string
}

// App

type App struct {
	state    utils.State
	fields   map[utils.State]Field
	debugger debug.Debugger
}

func NewApp(state utils.State, fields map[utils.State]Field, enable_debug bool) App {
	dbg := debug.New()
	if enable_debug {
		dbg.Enable()
	}
	mod := App{
		state:    state,
		fields:   fields,
		debugger: dbg,
	}
	return mod
}

func (m App) Init() tea.Cmd {
	return nil
}

func (m App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.debugger.Log(msg.String())
		switch {
		case key.Matches(msg, keys.Keys.Quit):
			return m, tea.Quit
		}
	case utils.NextStateMsg:
		m.state = m.state.Next()
		m.debugger.Log(fmt.Sprintf("state - %s", m.state.Label()))
		return m, nil
	case utils.PrevStateMsg:
		m.state = m.state.Prev()
		m.debugger.Log(fmt.Sprintf("state - %s", m.state.Label()))
		return m, nil
	case utils.DebugMsg:
		m.debugger.Log(msg.String())
	}
	fld, cmd := m.fields[m.state].Update(msg)
	m.fields[m.state] = fld.(Field)
	return m, cmd
}

func (m App) View() string {
	fieldView := m.fields[m.state].View()
	debugView := ui.DebugView(m.debugger.Enabled(), m.debugger.Get())
	return ui.AppView(fieldView, debugView)
}

func (m App) AddField(st utils.State, f Field) App {
	m.fields[st] = f
	return m
}

// Field

type Field struct {
	elements  []component.Viewer
	label     string
	ncontrols int
}

func NewField(label string, isFirstField bool, isLastField bool) Field {
	f := Field{
		elements:  []component.Viewer{},
		label:     label,
		ncontrols: 0,
	}
	if !isFirstField {
		f.ncontrols++
		prev := button.New("< back", utils.ControlBtn, utils.PrevStateMsg{})
		f = f.AddElement(&prev)
	}
	if !isLastField {
		f.ncontrols++
		next := button.New("next >", utils.ControlBtn, utils.NextStateMsg{})
		f = f.AddElement(&next)
	}
	f.FocusLast()
	return f
}

func (f Field) Init() tea.Cmd {
	return nil
}

func (f Field) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (f Field) View() string {
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
	return ui.FieldView(
		f.label,
		append(
			elementViews,
			ui.FieldControlView(controlViews...),
		)...,
	)
}

func (f Field) AddElement(element component.Viewer) Field {
	f.elements = append(f.elements, element)
	f.FocusNext()
	f.FocusPrev()
	return f
}

func (f Field) GetViewersWithoutControls() []component.Viewer {
	return f.elements[f.ncontrols:]
}

func (f Field) GetControls() []component.Viewer {
	return f.elements[:f.ncontrols]
}

func (f Field) GetFocusers() []component.Focuser {
	var focusers []component.Focuser
	for _, element := range f.elements {
		switch el := element.(type) {
		case component.Focuser:
			focusers = append(focusers, el)
		}
	}
	return focusers
}

func (f *Field) FocusNext() {
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

func (f *Field) FocusPrev() {
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

func (f *Field) FocusLast() {
	focusers := f.GetFocusers()
	if len(focusers) > 0 {
		for _, f := range focusers {
			f.Blur()
		}
		focusers[len(focusers)-1].Focus()
	}
}

func (f *Field) FocusFirst() {
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
