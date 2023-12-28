package button

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

var _ obj.Accessor = (*Button)(nil)
var _ obj.Actor = (*Button)(nil)
var _ obj.Element = (*Button)(nil)

type Button struct {
	obj.ElementWithID
	obj.ElementWithCallback
	label    string
	npresses int
	btntype  utils.ButtonType
}

func New(id int, label string, btntype utils.ButtonType, callback tea.Msg) Button {
	return Button{
		ElementWithID:       obj.NewElementWithID(id),
		ElementWithCallback: obj.NewElementWithCallback(callback),
		label:               label,
		npresses:            0,
		btntype:             btntype,
	}
}

// implementing Element
func (b Button) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Space) || key.Matches(msg, keys.Keys.Enter):
			b.npresses++
			cmds = append(cmds, utils.Callback(b.Callback()))
			cmds = append(cmds, utils.DebugCmd(fmt.Sprintf("%s called %d times", b.label, b.npresses)))
		}
	}
	return b, tea.Batch(cmds...)
}

func (b Button) View(focused bool) string {
	return ui.ButtonView(focused, b.label, b.btntype)
}

// implementing Accessor
func (b Button) Data() interface{} {
	return b.npresses
}

// special
func (b Button) Set(lbl string) Button {
	b.label = lbl
	return b
}
