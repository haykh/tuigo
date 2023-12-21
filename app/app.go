package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/debug"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/obj/button"
	"github.com/haykh/tuigo/obj/container"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

type Constructor = (func(obj.Element) obj.Element)
type AppState = string

type Backend struct {
	States       []AppState
	Constructors map[AppState]Constructor
	Finalizer    func(map[AppState]obj.Element)
}

type App struct {
	activeState AppState
	debugger    debug.Debugger
	backend     Backend
	containers  map[AppState]obj.Element
}

func New(backend Backend, enable_debug bool) App {
	if len(backend.States) == 0 {
		panic("No states provided")
	}
	dbg := debug.New()
	if enable_debug {
		dbg.Enable()
	}
	if _, ok := backend.Constructors[backend.States[0]]; !ok {
		panic("No constructor for initial state")
	}
	return App{
		activeState: backend.States[0],
		backend:     backend,
		containers:  map[AppState]obj.Element{},
		debugger:    dbg,
	}
}

func (a App) Init() tea.Cmd {
	head_container := a.backend.Constructors[a.backend.States[0]](nil)
	if head, ok := head_container.(obj.Collection); ok {
		is_first := true
		is_last := len(a.backend.States) == 1
		head_container = head.AddElements(a.GenerateControls(is_first, is_last))
	} else {
		panic("Head container must be a collection")
	}
	a.containers[a.activeState] = head_container
	a.containers[a.activeState] = a.containers[a.activeState].(obj.Collection).Focus()
	return nil
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Quit):
			return a, tea.Quit
		case key.Matches(msg, keys.Keys.Tab):
			cont, cmd := a.containers[a.activeState].Update(utils.FocusNextMsg{})
			a.containers[a.activeState] = cont
			return a, cmd
		case key.Matches(msg, keys.Keys.ShiftTab):
			cont, cmd := a.containers[a.activeState].Update(utils.FocusPrevMsg{})
			a.containers[a.activeState] = cont
			return a, cmd
		default:
			cont, cmd := a.containers[a.activeState].Update(msg)
			a.containers[a.activeState] = cont
			return a, cmd
		}
	case utils.SubmitMsg:
		a.debugger.Log("Submit")
		a.backend.Finalizer(a.containers)
		return a, tea.Quit
	case utils.NextStateMsg:
		a.debugger.Log("Next")
		return a.NextState(), nil
	case utils.PrevStateMsg:
		a.debugger.Log("Prev")
		return a.PrevState(), nil
	case utils.DebugMsg:
		a.debugger.Log(msg.String())
		return a, nil
	}
	return a, nil
}

func (a App) View() string {
	containerView := a.containers[a.activeState].View(false)
	debugView := ui.DebugView(a.debugger.Enabled(), a.debugger.Get())
	return ui.AppView(containerView, debugView)
}

func (a App) GenerateControls(isFirst, isLast bool) obj.Element {
	controls := []obj.Element{}
	prevbtn := button.New("< prev", utils.ControlBtn, utils.PrevStateMsg{})
	nextbtn := button.New("next >", utils.ControlBtn, utils.NextStateMsg{})
	submitbtn := button.New("submit", utils.ControlBtn, utils.SubmitMsg{})
	if !isFirst {
		controls = append(controls, prevbtn)
	}
	if !isLast {
		controls = append(controls, nextbtn)
	} else {
		controls = append(controls, submitbtn)
	}
	return container.New(true,
		utils.HorizontalContainer,
		controls...,
	)
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
	currentContainer := a.containers[currentState]
	if _, ok := a.backend.Constructors[newState]; !ok {
		panic(fmt.Sprintf("No constructor for next state: %s", newState))
	}
	newContainer := a.backend.Constructors[newState](currentContainer)
	is_first := false
	is_last := (newState_idx == len(a.backend.States)-1)
	if head, ok := newContainer.(obj.Collection); ok {
		newContainer = head.AddElements(a.GenerateControls(is_first, is_last))
	} else {
		panic("New container must be a collection")
	}
	a.activeState = newState
	a.containers[newState] = newContainer.(obj.Collection).Focus()
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
	if newContainer, ok := a.containers[newState]; !ok {
		panic(fmt.Sprintf("No container for prev state: %s", newState))
	} else {
		a.activeState = newState
		a.containers[newState] = newContainer.(obj.Collection).Focus()
		return a
	}
}
