package app

import (
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
		head_container = head.AddElements(a.GenerateControls(true, false))
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
	case utils.DebugMsg:
		a.debugger.Log(msg.String())
	}
	return a, nil
}

func (a App) View() string {
	containerView := a.containers[a.activeState].View(false)
	debugView := ui.DebugView(a.debugger.Enabled(), a.debugger.Get())
	return ui.AppView(containerView, debugView)
}

func (a App) GenerateControls(isFirst, isLast bool) obj.Collection {
	var controls obj.Collection
	prevbtn := button.New("< prev", utils.ControlBtn, utils.PrevStateMsg{})
	nextbtn := button.New("next >", utils.ControlBtn, utils.NextStateMsg{})
	if isFirst {
		controls = container.New(true,
			utils.HorizontalContainer,
			nextbtn,
		)
	} else if isLast {
		controls = container.New(true,
			utils.HorizontalContainer,
			prevbtn,
		)
	} else {
		controls = container.New(true,
			utils.HorizontalContainer,
			prevbtn,
			nextbtn,
		)
	}
	return controls
}

func (a *App) Next() {}
