package utils

import (
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
)

// container types
type ContainerType int

const (
	SimpleContainer ContainerType = iota
	VerticalContainer
	HorizontalContainer
)

// text types
type TextType int

const (
	NormalText TextType = iota
	LabelText
	DimmedText
)

// item types
type ItemType = string

const (
	Container ItemType = "container"
	Button    ItemType = "button"
	Input     ItemType = "input"
	Radio     ItemType = "radio"
	Selector  ItemType = "selector"
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

// text types

func Callback(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

func DebugCmd(msg string) tea.Cmd {
	return func() tea.Msg {
		return DebugMsg{msg}
	}
}

// messages
type DebugMsg struct {
	msg string
}

func (dbg DebugMsg) String() string {
	return dbg.msg
}

type FocusNextMsg struct{}
type FocusPrevMsg struct{}
type NextStateMsg struct{}
type PrevStateMsg struct{}
type SubmitMsg struct{}

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
