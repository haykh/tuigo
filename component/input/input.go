package input

import (
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/component"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

type Model struct {
	component.TextInputWrap
	inputtype utils.InputType
}

func New(label, def, placeholder string, inputtype utils.InputType) Model {
	m := textinput.New()
	m.SetValue(def)
	m.Placeholder = placeholder
	m.Prompt = label
	if inputtype == utils.PathInput {
		m.ShowSuggestions = true
		m.KeyMap.AcceptSuggestion = keys.Keys.Right
		m.KeyMap.NextSuggestion = keys.Keys.Down
		m.KeyMap.PrevSuggestion = keys.Keys.Up
	}
	return Model{
		TextInputWrap: component.NewTextInputWrap(m),
		inputtype:     inputtype,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (component.Updater, tea.Cmd) {
	if m.inputtype == utils.PathInput {
		// update suggestions
		var suggestions []string
		entry := m.Model.Value()
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
			m.Model.SetSuggestions(suggestions)
		}
	}

	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return &m, cmd
}

func (m Model) View() string {
	return ui.PathInputView(m.Model)
}

func (m Model) Value() string {
	return m.Model.Value()
}
