package utils

import (
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
)

type ComponentType = string

const (
	Container ComponentType = "container"
	Button    ComponentType = "button"
	Input     ComponentType = "input"
	Radio     ComponentType = "radio"
	Selector  ComponentType = "selector"
)

// button types
type ButtonType int

const (
	SimpleBtn ButtonType = iota
	AcceptBtn
	ControlBtn
)

// input types

type InputType int

const (
	PathInput InputType = iota
	TextInput
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

// for testing

func CheckCmd(cmd tea.Cmd, want tea.Msg) bool {
	if cmd == nil {
		return want == nil
	}
	found := false
	switch cmd := cmd().(type) {
	case tea.BatchMsg:
		for _, c := range cmd {
			found = found || CheckCmd(c, want)
			if found {
				return true
			}
		}
	case tea.Msg:
		if reflect.TypeOf(cmd) == reflect.TypeOf(want) {
			found = true
		}
	}
	return found
}
