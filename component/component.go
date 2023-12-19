package component

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Viewer interface {
	View() string
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
	focused bool
}

func (f Component) Focused() bool {
	return f.focused
}

func (f *Component) Focus() tea.Cmd {
	f.focused = true
	return nil
}

func (f *Component) Blur() {
	f.focused = false
}

type TextInputWrap struct {
	Model textinput.Model
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
