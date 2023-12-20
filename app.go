package tuigo

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/debug"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

type App struct {
	state      utils.State
	containers map[utils.State]Container
	debugger   debug.Debugger
}

func NewApp(state utils.State, containers map[utils.State]Container, enable_debug bool) App {
	dbg := debug.New()
	if enable_debug {
		dbg.Enable()
	}
	mod := App{
		state:      state,
		containers: containers,
		debugger:   dbg,
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
	fld, cmd := m.containers[m.state].Update(msg)
	m.containers[m.state] = fld.(Container)
	return m, cmd
}

func (m App) View() string {
	containerView := m.containers[m.state].View()
	debugView := ui.DebugView(m.debugger.Enabled(), m.debugger.Get())
	return ui.AppView(containerView, debugView)
}
