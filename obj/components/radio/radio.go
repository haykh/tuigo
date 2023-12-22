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

var _ obj.Element = (*Radio)(nil)
var _ obj.Actor = (*Radio)(nil)
var _ obj.Accessor = (*Radio)(nil)

type Radio struct {
	obj.ElementWithID
	obj.ElementWithCallback
	label string
	state bool
}

func New(id int, label string, callback tea.Msg) container.SimpleContainer {
	return container.NewSimpleContainer(true, Radio{
		ElementWithID:       obj.NewElementWithID(id),
		ElementWithCallback: obj.NewElementWithCallback(callback),
		label:               label,
		state:               false,
	})
}

// implementing Element
func (r Radio) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Space):
			r = r.Toggle()
			cmds = append(cmds, utils.Callback(r.Callback()))
			cmds = append(cmds, utils.DebugCmd(fmt.Sprintf("%s toggled", r.label)))
		}
	}
	return r, tea.Batch(cmds...)
}

func (r Radio) View(focused bool) string {
	return ui.RadioView(focused, r.label, r.state)
}

// implementing Accessor
func (r Radio) Data() interface{} {
	return r.state
}

// special
func (r Radio) Toggle() Radio {
	r.state = !r.state
	return r
}
