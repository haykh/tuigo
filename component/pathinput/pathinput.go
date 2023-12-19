package pathinput

import (
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/component"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/ui"
)

type Model struct {
	component.TextInputWrap
}

func New(label, def, placeholder string) Model {
	m := textinput.New()
	m.SetValue(def)
	m.Placeholder = placeholder
	m.Prompt = label
	m.ShowSuggestions = true
	m.KeyMap.AcceptSuggestion = keys.Keys.Right
	m.KeyMap.NextSuggestion = keys.Keys.Down
	m.KeyMap.PrevSuggestion = keys.Keys.Up
	return Model{
		TextInputWrap: component.TextInputWrap{Model: m},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (component.Updater, tea.Cmd) {
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

	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return &m, cmd
}

func (m Model) View() string {
	return ui.PathInputView(m.Model)
}
