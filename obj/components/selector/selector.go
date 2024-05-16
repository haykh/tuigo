package selector

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

var _ obj.Accessor = (*Selector)(nil)
var _ obj.Actor = (*Selector)(nil)
var _ obj.Element = (*Selector)(nil)

type Selector struct {
	obj.ElementWithID
	obj.ElementWithCallback
	multiselect bool
	cursor      int
	options     []string
	selected    map[string]struct{}
	disabled    map[string]struct{}
	view_limit  int
}

func New(id int, options []string, multiselect bool, view_limit int, callback tea.Msg) Selector {
	return Selector{
		ElementWithID:       obj.NewElementWithID(id),
		ElementWithCallback: obj.NewElementWithCallback(callback),
		multiselect:         multiselect,
		cursor:              0,
		options:             options,
		selected:            map[string]struct{}{},
		disabled:            map[string]struct{}{},
		view_limit:          view_limit,
	}
}

// implementing Element
func (s Selector) View(focused bool) string {
	return ui.SelectorView(
		focused,
		s.multiselect,
		s.cursor,
		s.options,
		s.selected,
		s.disabled,
		s.view_limit,
	)
}

func (s Selector) Update(msg tea.Msg) (obj.Element, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Up):
			s = s.Prev()
		case key.Matches(msg, keys.Keys.Down):
			s = s.Next()
		case key.Matches(msg, keys.Keys.Space):
			s = s.Toggle()
			cmds = append(cmds, utils.Callback(s.Callback()))
			cmds = append(cmds, utils.DebugCmd(fmt.Sprintf("%s toggled", s.options[s.cursor])))
		}
	}
	return s, tea.Batch(cmds...)
}

// implementing Accessor
func (m Selector) Data() interface{} {
	if m.multiselect {
		return m.Selected()
	} else {
		if len(m.Selected()) == 0 {
			return nil
		} else {
			return m.Selected()[0]
		}
	}
}

// special
func (s Selector) Disable(opt string) Selector {
	s.disabled[opt] = struct{}{}
	delete(s.selected, opt)
	return s
}

func (s Selector) Enable(opt string) Selector {
	delete(s.disabled, opt)
	return s
}

func (s Selector) Disabled(opt string) bool {
	_, ok := s.disabled[opt]
	return ok
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

func (s Selector) ToggleSpecific(opt string) Selector {
	prev_cursor := s.cursor
	s.cursor = -1
	for i, o := range s.options {
		if o == opt {
			s.cursor = i
			break
		}
	}
	if s.cursor != -1 {
		s = s.Toggle()
	}
	s.cursor = prev_cursor
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

func (m Selector) SetViewLimit(limit int) Selector {
	m.view_limit = limit
	return m
}
