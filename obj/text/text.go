package text

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/obj/container"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

type Model struct {
	texttype utils.TextType
	txt      string
}

func New(txt string, texttype utils.TextType) obj.Element {
	return container.NewSimpleContainer(false, Model{
		texttype: texttype,
		txt:      txt,
	})
}

func (m Model) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	return m, nil
}

func (m Model) View(bool) string {
	return ui.TextView(false, m.txt, m.texttype)
}
