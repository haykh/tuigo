package input

import (
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/obj/container"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

var _ obj.Element = (*Model)(nil)

type Model struct {
	inputtype utils.InputType
	model     textinput.Model
}

func NewTextinput(label, def, placeholder string) textinput.Model {
	m := textinput.New()
	m.Focus()
	m.SetValue(def)
	m.Placeholder = placeholder
	m.Prompt = label
	return m
}

func New(label, def, placeholder string, inputtype utils.InputType) obj.Element {
	m := NewTextinput(label, def, placeholder)
	if inputtype == utils.PathInput {
		m.ShowSuggestions = true
		m.KeyMap.AcceptSuggestion = keys.Keys.Right
		m.KeyMap.NextSuggestion = keys.Keys.Down
		m.KeyMap.PrevSuggestion = keys.Keys.Up
	}
	return container.NewSimpleContainer(true, Model{
		inputtype: inputtype,
		model:     m,
	})
}

func (m Model) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	if m.inputtype == utils.PathInput {
		// update suggestions
		var suggestions []string
		entry := m.model.Value()
		for i := len(entry) - 1; i >= 0; i-- {
			if entry[i] == '/' {
				entry = entry[:i+1]
				break
			}
		}
		resolved := os.Expand(entry, func(v string) string {
			return os.Getenv(v)
		})

		if entries, err := os.ReadDir(resolved); err == nil {
			for _, e := range entries {
				suggestions = append(suggestions, entry+e.Name())
			}
			m.model.SetSuggestions(suggestions)
		}
	}

	var cmd tea.Cmd
	m.model, cmd = m.model.Update(msg)
	return m, cmd
}

func (m Model) View(focused bool) string {
	if focused {
		m.model.Focus()
	} else {
		m.model.Blur()
	}
	return ui.PathInputView(focused, m.model)
}

func (m Model) Value() string {
	return m.model.Value()
}
