package radio

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/obj/container"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

var _ obj.Element = (*Model)(nil)

type Model struct {
	label string
	state bool
}

func New(label string) obj.Element {
	return container.NewSimpleContainer(true, Model{
		label: label,
		state: false,
	})
}

func (m Model) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Space):
			m.Toggle()
			cmd = utils.DebugCmd(fmt.Sprintf("%s toggled", m.label))
		}
	}
	return m, cmd
}

func (m Model) View(focused bool) string {
	return ui.RadioView(focused, m.label, m.state)
}

// access

func (m *Model) Toggle() {
	m.state = !m.state
}

func (m Model) State() bool {
	return m.state
}
