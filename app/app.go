package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/debug"
	"github.com/haykh/tuigo/keys"
	component "github.com/haykh/tuigo/obj/components"
	"github.com/haykh/tuigo/obj/container"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/ui/theme"
	"github.com/haykh/tuigo/utils"
)

type Window = container.Collection
type Constructor = (func(Window) Window)
type Updater = (func(Window, tea.Msg) (Window, tea.Cmd))
type AppState = string

type Backend struct {
	States       []AppState
	Constructors map[AppState]Constructor
	Updaters     map[AppState]Updater
	Finalizer    func(map[AppState]Window) Window
}

type App struct {
	activeState AppState
	debugger    debug.Debugger
	backend     Backend
	windows     map[AppState]Window
}

func New(backend Backend, enable_debug bool) App {
	if len(backend.States) == 0 {
		panic("No states provided")
	}
	dbg := debug.New()
	if enable_debug {
		dbg.Enable()
		theme.DEBUG_MODE = true
	}
	if _, ok := backend.Constructors[backend.States[0]]; !ok {
		panic("No constructor for initial state")
	}
	return App{
		activeState: backend.States[0],
		backend:     backend,
		windows:     map[AppState]Window{},
		debugger:    dbg,
	}
}

func NewControls(isFirst, isLast bool) container.Collection {
	controls := []container.Component{}
	prevbtn := component.NewButton(-100, "< prev", utils.ControlBtn, utils.PrevStateMsg{})
	nextbtn := component.NewButton(-200, "next >", utils.ControlBtn, utils.NextStateMsg{})
	submitbtn := component.NewButton(-300, "submit", utils.ControlBtn, utils.SubmitMsg{})
	if !isFirst {
		controls = append(controls, prevbtn)
	}
	if !isLast {
		controls = append(controls, nextbtn)
	} else {
		controls = append(controls, submitbtn)
	}
	return container.NewComplexContainer(true,
		utils.HorizontalContainer,
		controls...,
	)
}

func (a App) Init() tea.Cmd {
	head_container := a.backend.Constructors[a.backend.States[0]](nil)
	is_first := true
	is_last := len(a.backend.States) == 1
	parentContainer := container.NewComplexContainer(
		true,
		utils.VerticalContainer,
		head_container,
		NewControls(is_first, is_last),
	)
	parentContainer = parentContainer.Focus().(container.ComplexContainer)
	a.windows[a.activeState] = parentContainer
	return nil
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if updater, ok := a.backend.Updaters[a.activeState]; ok {
		cont, cmd := updater(a.windows[a.activeState], msg)
		if cmd != nil {
			a.windows[a.activeState] = cont
			return a, cmd
		}
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Quit):
			return a, tea.Quit
		case key.Matches(msg, keys.Keys.Tab):
			cont, cmd := a.windows[a.activeState].Update(utils.FocusNextMsg{})
			a.windows[a.activeState] = cont.(Window)
			return a, cmd
		case key.Matches(msg, keys.Keys.ShiftTab):
			cont, cmd := a.windows[a.activeState].Update(utils.FocusPrevMsg{})
			a.windows[a.activeState] = cont.(Window)
			return a, cmd
		default:
			cont, cmd := a.windows[a.activeState].Update(msg)
			a.windows[a.activeState] = cont.(Window)
			return a, cmd
		}
	case utils.SubmitMsg:
		a.debugger.Log("Submit")
		a.activeState = "FINAL"
		a.windows[a.activeState] = a.backend.Finalizer(a.windows)
		return a, tea.Quit
	case utils.NextStateMsg:
		a.debugger.Log("Next")
		return a.NextState(), nil
	case utils.PrevStateMsg:
		a.debugger.Log("Prev")
		return a.PrevState(), nil
	case utils.TargetedMsg:
		cont, cmd := a.windows[a.activeState].Update(msg)
		a.windows[a.activeState] = cont.(Window)
		return a, cmd
	case utils.DebugMsg:
		a.debugger.Log(msg.String())
		return a, nil
	}
	return a, nil
}

func (a App) View() string {
	containerView := a.windows[a.activeState].View(false)
	debugView := ui.DebugView(a.debugger.Enabled(), a.debugger.Get())
	return ui.AppView(containerView, debugView)
}

func (a App) NextState() App {
	currentState := a.activeState
	var newState AppState
	var newState_idx int
	for si, s := range a.backend.States {
		if s == currentState {
			if si+1 < len(a.backend.States) {
				newState = a.backend.States[si+1]
				newState_idx = si + 1
				break
			} else {
				panic("No next state")
			}
		}
	}
	currentContainer := a.windows[currentState]
	if _, ok := a.backend.Constructors[newState]; !ok {
		panic(fmt.Sprintf("No constructor for next state: %s", newState))
	}
	newContainer := a.backend.Constructors[newState](currentContainer)
	is_first := false
	is_last := (newState_idx == len(a.backend.States)-1)
	parentContainer := container.NewComplexContainer(
		true,
		utils.VerticalContainer,
		newContainer,
		NewControls(is_first, is_last),
	)
	a.activeState = newState
	parentContainer = parentContainer.Focus().(container.ComplexContainer)
	a.windows[newState] = parentContainer
	return a
}

func (a App) PrevState() App {
	currentState := a.activeState
	var newState AppState
	for si, s := range a.backend.States {
		if s == currentState {
			if si-1 >= 0 {
				newState = a.backend.States[si-1]
				break
			} else {
				panic("No prev state")
			}
		}
	}
	if newContainer, ok := a.windows[newState]; !ok {
		panic(fmt.Sprintf("No container for prev state: %s", newState))
	} else {
		a.activeState = newState
		newContainer = newContainer.Focus().(container.ComplexContainer)
		a.windows[newState] = newContainer
		return a
	}
}
