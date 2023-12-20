package button

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
	label   string
	btntype utils.ButtonType
	action  tea.Msg
}

func New(label string, btntype utils.ButtonType, action tea.Msg) Model {
	return Model{
		Component: component.NewComponent(utils.Button),
		label:     label,
		btntype:   btntype,
		action:    action,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (component.Updater, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Space) || key.Matches(msg, keys.Keys.Enter):
			cmds = append(cmds, utils.Callback(m))
			cmds = append(cmds, utils.DebugCmd(fmt.Sprintf("%s called", m.label)))
		}
	}
	return &m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return ui.ButtonView(m.label, m.Focused(), m.btntype)
}

// access

func (m Model) Action() tea.Msg {
	return m.action
}
