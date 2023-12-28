package input

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

var _ obj.Accessor = (*Input)(nil)
var _ obj.Actor = (*Input)(nil)
var _ obj.Element = (*Input)(nil)

type Input struct {
	obj.ElementWithID
	obj.ElementWithCallback
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

func New(id int, label, def, placeholder string, inputtype utils.InputType, callback tea.Msg) Input {
	m := NewTextinputModel(label, def, placeholder)
	if inputtype == utils.PathInput {
		m.ShowSuggestions = true
		m.KeyMap.AcceptSuggestion = keys.Keys.Right
		m.KeyMap.NextSuggestion = keys.Keys.Down
		m.KeyMap.PrevSuggestion = keys.Keys.Up
	}
	return Input{
		ElementWithID:       obj.NewElementWithID(id),
		ElementWithCallback: obj.NewElementWithCallback(callback),
		inputtype:           inputtype,
		model:               m,
	}
}

// implementing Element
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
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		var cmd tea.Cmd
		ti.model, cmd = ti.model.Update(msg)
		cmds = append(cmds, cmd)
		cmds = append(cmds, utils.Callback(ti.Callback()))
		cmds = append(cmds, utils.DebugCmd(fmt.Sprintf("input: %s", ti.model.Value())))
	}
	return ti, tea.Batch(cmds...)
}

func (ti Input) View(focused bool) string {
	if focused {
		ti.model.Focus()
	} else {
		ti.model.Blur()
	}
	return ui.PathInputView(focused, ti.model)
}

// implementing Accessor
func (ti Input) Data() interface{} {
	return ti.model.Value()
}
