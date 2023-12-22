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

var _ obj.Element = (*Input)(nil)
var _ obj.Accessor = (*Input)(nil)

type Input struct {
	obj.ElementWithID
	inputtype utils.InputType
	model     textinput.Model
}

func NewTextinputModel(label, def, placeholder string) textinput.Model {
	m := textinput.New()
	m.Focus()
	m.SetValue(def)
	m.Placeholder = placeholder
	m.Prompt = label
	return m
}

func New(id int, label, def, placeholder string, inputtype utils.InputType) obj.Collection {
	m := NewTextinputModel(label, def, placeholder)
	if inputtype == utils.PathInput {
		m.ShowSuggestions = true
		m.KeyMap.AcceptSuggestion = keys.Keys.Right
		m.KeyMap.NextSuggestion = keys.Keys.Down
		m.KeyMap.PrevSuggestion = keys.Keys.Up
	}
	return container.NewSimpleContainer(true, Input{
		ElementWithID: obj.NewElementWithID(id),
		inputtype:     inputtype,
		model:         m,
	})
}

func (ti Input) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	if ti.inputtype == utils.PathInput {
		// update suggestions
		var suggestions []string
		entry := ti.model.Value()
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
			ti.model.SetSuggestions(suggestions)
		}
	}

	var cmd tea.Cmd
	ti.model, cmd = ti.model.Update(msg)
	return ti, cmd
}

func (ti Input) View(focused bool) string {
	if focused {
		ti.model.Focus()
	} else {
		ti.model.Blur()
	}
	return ui.PathInputView(focused, ti.model)
}

func (ti Input) Data() interface{} {
	return ti.model.Value()
}
