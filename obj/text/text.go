package text

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/obj/container"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

var _ obj.Element = (*Text)(nil)
var _ obj.Accessor = (*Text)(nil)

type Text struct {
	obj.ElementWithID
	texttype utils.TextType
	txt      string
}

func New(id int, txt string, texttype utils.TextType) obj.Element {
	return container.NewSimpleContainer(false, Text{
		ElementWithID: obj.NewElementWithID(id),
		texttype:      texttype,
		txt:           txt,
	})
}

func (t Text) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	return t, nil
}

func (t Text) View(bool) string {
	return ui.TextView(false, t.txt, t.texttype)
}
