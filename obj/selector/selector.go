package selector

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/obj/container"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

var _ obj.Element = (*Selector)(nil)
var _ obj.Accessor = (*Selector)(nil)

type Selector struct {
	obj.ElementWithID
	multiselect bool
	cursor      int
	options     []string
	selected    map[string]struct{}
	disabled    map[string]struct{}
}

func New(id int, options []string, multiselect bool) obj.Element {
	return container.NewSimpleContainer(true, Selector{
		ElementWithID: obj.NewElementWithID(id),
		multiselect:   multiselect,
		cursor:        0,
		options:       options,
		selected:      map[string]struct{}{},
		disabled:      map[string]struct{}{},
	})
}

func (s Selector) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Up):
			s = s.Prev()
		case key.Matches(msg, keys.Keys.Down):
			s = s.Next()
		case key.Matches(msg, keys.Keys.Space):
			s = s.Toggle()
			cmd = utils.DebugCmd(fmt.Sprintf("%s toggled", s.options[s.cursor]))
		}
	}
	return s, cmd
}

func (s Selector) View(focused bool) string {
	return ui.SelectorView(focused, s.multiselect, s.cursor, s.options, s.selected, s.disabled)
}

// access

func (s Selector) Disable(opt string) Selector {
	s.disabled[opt] = struct{}{}
	delete(s.selected, opt)
	return s
}

func (s Selector) Enable(opt string) Selector {
	delete(s.disabled, opt)
	return s
}

func (s Selector) Toggle() Selector {
	if _, ok := s.selected[s.options[s.cursor]]; ok {
		delete(s.selected, s.options[s.cursor])
	} else {

		if !s.multiselect {
			s.selected = map[string]struct{}{s.options[s.cursor]: {}}
		} else {
			s.selected[s.options[s.cursor]] = struct{}{}
		}
	}
	return s
}

func (s Selector) Next() Selector {
	s.cursor = (s.cursor + 1 + len(s.options)) % len(s.options)
	if _, ok := s.disabled[s.options[s.cursor]]; ok {
		s = s.Next()
	}
	return s
}

func (s Selector) Prev() Selector {
	s.cursor = (s.cursor - 1 + len(s.options)) % len(s.options)
	if _, ok := s.disabled[s.options[s.cursor]]; ok {
		s.Prev()
	}
	return s
}

func (m Selector) Selected() []string {
	var selected []string
	for _, o := range m.options {
		if _, ok := m.selected[o]; ok {
			selected = append(selected, o)
		}
	}
	return selected
}

func (m Selector) Cursor() int {
	return m.cursor
}
