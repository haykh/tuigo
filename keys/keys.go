package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Up        key.Binding
	Down      key.Binding
	Left      key.Binding
	Right     key.Binding
	Space     key.Binding
	Enter     key.Binding
	Tab       key.Binding
	ShiftTab  key.Binding
	HelpOpen  key.Binding
	HelpClose key.Binding
	Quit      key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.HelpOpen, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Left, k.Space, k.Enter, k.Tab},
		{k.HelpOpen, k.Quit},
	}
}

var Keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑/↓", "move up/down"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("⇆/[⇧]⇆", "cycle"),
	),
	ShiftTab: key.NewBinding(
		key.WithKeys("shift+tab"),
	),
	Left: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("←/→", "move left/right"),
	),
	Right: key.NewBinding(
		key.WithKeys("right"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↵", "confirm"),
	),
	Space: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("⎵", "select"),
	),
	HelpOpen: key.NewBinding(
		key.WithKeys("shift+up"),
		key.WithHelp("[⇧]↑/↓", "toggle help"),
	),
	HelpClose: key.NewBinding(
		key.WithKeys("shift+down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("Esc/Ctrl+c", "quit"),
	),
}
