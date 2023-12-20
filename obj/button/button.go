package button

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
	label   string
	btntype utils.ButtonType
	action  tea.Msg
}

func New(label string, btntype utils.ButtonType, action tea.Msg) obj.Element {
	return container.NewSimpleContainer(true, Model{
		label:   label,
		btntype: btntype,
		action:  action,
	})
}

func (m Model) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Space) || key.Matches(msg, keys.Keys.Enter):
			cmds = append(cmds, utils.Callback(m.Action()))
			cmds = append(cmds, utils.DebugCmd(fmt.Sprintf("%s called", m.label)))
		}
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View(focused bool) string {
	return ui.ButtonView(focused, m.label, m.btntype)
}

func (m Model) Action() tea.Msg {
	return m.action
}
