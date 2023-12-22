package text

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/obj/container"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

var _ obj.Accessor = (*Text)(nil)
var _ obj.Actor = (*Text)(nil)
var _ obj.Element = (*Text)(nil)

type Text struct {
	obj.ElementWithID
	obj.ElementWithCallback
	texttype utils.TextType
	txt      string
}

func New(id int, txt string, texttype utils.TextType) container.SimpleContainer {
	return container.NewSimpleContainer(false, Text{
		ElementWithID:       obj.NewElementWithID(id),
		ElementWithCallback: obj.NewElementWithCallback(nil),
		texttype:            texttype,
		txt:                 txt,
	})
}

// implementing Element
func (t Text) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	return t, nil
}

func (t Text) View(bool) string {
	return ui.TextView(false, t.txt, t.texttype)
}

// implementing Accessor
func (t Text) Data() interface{} {
	return t.txt
}

// special
func (t Text) Set(txt string) Text {
	t.txt = txt
	return t
}
