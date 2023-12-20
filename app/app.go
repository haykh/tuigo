package app

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/debug"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

type App struct {
	Container obj.Element
	Debugger  debug.Debugger
}

func (a App) Init() tea.Cmd {
	return nil
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Quit):
			return a, tea.Quit
		case key.Matches(msg, keys.Keys.Tab):
			cont, cmd := a.Container.Update(utils.FocusNextMsg{})
			a.Container = cont
			return a, cmd
		case key.Matches(msg, keys.Keys.ShiftTab):
			cont, cmd := a.Container.Update(utils.FocusPrevMsg{})
			a.Container = cont
			return a, cmd
		default:
			cont, cmd := a.Container.Update(msg)
			a.Container = cont
			return a, cmd
		}
	case utils.DebugMsg:
		a.Debugger.Log(msg.String())
	}
	return a, nil
}

func (a App) View() string {
	containerView := a.Container.View(false)
	debugView := ui.DebugView(a.Debugger.Enabled(), a.Debugger.Get())
	return ui.AppView(containerView, debugView)
}
