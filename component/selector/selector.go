package selector

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/component"
	"github.com/haykh/tuigo/keys"
	"github.com/haykh/tuigo/ui"
	"github.com/haykh/tuigo/utils"
)

type Model struct {
	component.Component
	multiselect bool
	cursor      int
	options     []string
	selected    map[string]struct{}
	disabled    map[string]struct{}
}

func New(options []string, multiselect bool) Model {
	return Model{
		multiselect: multiselect,
		cursor:      0,
		options:     options,
		selected:    map[string]struct{}{},
		disabled:    map[string]struct{}{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (component.Updater, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Up):
			m.Prev()
		case key.Matches(msg, keys.Keys.Down):
			m.Next()
		case key.Matches(msg, keys.Keys.Space):
			m.Toggle()
			cmd = utils.DebugCmd(fmt.Sprintf("%s toggled", m.options[m.cursor]))
		}
	}
	return &m, cmd
}

func (m Model) View() string {
	return ui.SelectorView(m.multiselect, m.cursor, m.options, m.selected, m.disabled, m.Focused())
}

// access

func (m *Model) Disable(opt string) {
	m.disabled[opt] = struct{}{}
	delete(m.selected, opt)
}

func (m *Model) Enable(opt string) {
	delete(m.disabled, opt)
}

func (m *Model) Toggle() {
	if _, ok := m.selected[m.options[m.cursor]]; ok {
		delete(m.selected, m.options[m.cursor])
	} else {

		if !m.multiselect {
			m.selected = map[string]struct{}{m.options[m.cursor]: {}}
		} else {
			m.selected[m.options[m.cursor]] = struct{}{}
		}
	}
}

func (m *Model) Next() {
	m.cursor = (m.cursor + 1 + len(m.options)) % len(m.options)
	if _, ok := m.disabled[m.options[m.cursor]]; ok {
		m.Next()
	}
}

func (m *Model) Prev() {
	m.cursor = (m.cursor - 1 + len(m.options)) % len(m.options)
	if _, ok := m.disabled[m.options[m.cursor]]; ok {
		m.Prev()
	}
}

func (m Model) Selected() []string {
	var selected []string
	for _, o := range m.options {
		if _, ok := m.selected[o]; ok {
			selected = append(selected, o)
		}
	}
	return selected
}

func (m Model) Cursor() int {
	return m.cursor
}
