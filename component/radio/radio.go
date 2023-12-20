package radio

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/component"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

type Model struct {
	component.Component
	label string
	state bool
}

func New(label string) Model {
	return Model{
		Component: component.NewComponent(utils.Radio),
		label:     label,
		state:     false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (component.Updater, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Space):
			m.Toggle()
			cmd = utils.DebugCmd(fmt.Sprintf("%s toggled", m.label))
		}
	}
	return &m, cmd
}

func (m Model) View() string {
	return ui.RadioView(m.label, m.state, m.Focused())
}

// access

func (m *Model) Toggle() {
	m.state = !m.state
}

func (m Model) State() bool {
	return m.state
}
