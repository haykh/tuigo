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

var _ obj.Element = (*Button)(nil)
var _ obj.Accessor = (*Button)(nil)

type Button struct {
	obj.ElementWithID
	label    string
	npresses int
	btntype  utils.ButtonType
	action   tea.Msg
}

func New(id int, label string, btntype utils.ButtonType, action tea.Msg) obj.Element {
	return container.NewSimpleContainer(true, Button{
		ElementWithID: obj.NewElementWithID(id),
		label:         label,
		npresses:      0,
		btntype:       btntype,
		action:        action,
	})
}

func (b Button) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Space) || key.Matches(msg, keys.Keys.Enter):
			cmds = append(cmds, utils.Callback(b.Action()))
			cmds = append(cmds, utils.DebugCmd(fmt.Sprintf("%s called %d times", b.label, b.npresses)))
			b.npresses++
		}
	}
	return b, tea.Batch(cmds...)
}

func (b Button) View(focused bool) string {
	return ui.ButtonView(focused, b.label, b.btntype)
}

func (b Button) Action() tea.Msg {
	return b.action
}

func (b Button) Data() interface{} {
	return b.npresses
}
