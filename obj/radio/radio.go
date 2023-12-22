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
var _ obj.Accessor = (*Radio)(nil)

type Radio struct {
	obj.ElementWithID
	label string
	state bool
}

func New(id int, label string) obj.Collection {
	return container.NewSimpleContainer(true, Radio{
		ElementWithID: obj.NewElementWithID(id),
		label:         label,
		state:         false,
	})
}

func (r Radio) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Space):
			r = r.Toggle()
			cmd = utils.DebugCmd(fmt.Sprintf("%s toggled", r.label))
		}
	}
	return r, cmd
}

func (r Radio) View(focused bool) string {
	return ui.RadioView(focused, r.label, r.state)
}

func (r Radio) Toggle() Radio {
	r.state = !r.state
	return r
}

func (r Radio) Data() interface{} {
	return r.state
}
