package component

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/utils"
)

type Viewer interface {
	View() string
	String() string
}

type Focuser interface {
	Focused() bool
	Focus() tea.Cmd
	Blur()
}

type Updater interface {
	Focuser
	Update(tea.Msg) (Updater, tea.Cmd)
}

type Component struct {
	componentType utils.ComponentType
	focused       bool
}

func NewComponent(componentType utils.ComponentType) Component {
	return Component{
		componentType: componentType,
		focused:       false,
	}
}

func (f Component) Focused() bool {
	return f.focused
}

func (f Component) String() string {
	return f.componentType
}

func (f *Component) Focus() tea.Cmd {
	f.focused = true
	return nil
}

func (f *Component) Blur() {
	f.focused = false
}

type TextInputWrap struct {
	componentType utils.ComponentType
	Model         textinput.Model
}

func (t TextInputWrap) String() string {
	return t.componentType
}

func NewTextInputWrap(model textinput.Model) TextInputWrap {
	return TextInputWrap{
		componentType: utils.Input,
		Model:         model,
	}
}

func (t TextInputWrap) Focused() bool {
	return t.Model.Focused()
}

func (t *TextInputWrap) Focus() tea.Cmd {
	return t.Model.Focus()
}

func (t *TextInputWrap) Blur() {
	t.Model.Blur()
}
