package utils

import tea "github.com/charmbracelet/bubbletea"

// button types
type ButtonType int

const (
	SimpleBtn ButtonType = iota
	AcceptBtn
	ControlBtn
)

// window states
type State interface {
	Label() string
	Next() State
	Prev() State
}

// actions & messages
type Actor interface {
	Action() tea.Msg
}

type Messenger interface {
	Message() string
}

func Callback(a Actor) tea.Cmd {
	return func() tea.Msg {
		return a.Action()
	}
}

func DebugCmd(msg string) tea.Cmd {
	return func() tea.Msg {
		return DebugMsg{msg}
	}
}

// messages
type NextStateMsg struct{}
type PrevStateMsg struct{}

type DebugMsg struct {
	msg string
}

func (dbg DebugMsg) String() string {
	return dbg.msg
}
