package utils

import (
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
)

// container types
type ContainerType int

const (
	VerticalContainer ContainerType = iota
	VerticalContainerCenter
	VerticalContainerRight
	HorizontalContainer
	HorizontalContainerTop
	HorizontalContainerBottom
)

// text types
type TextType int

const (
	NormalText TextType = iota
	LabelText
	DimmedText
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

// messages
func Callback(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

type TargetedMsg struct {
	id     int
	action interface{}
}

func (msg TargetedMsg) ID() int {
	return msg.id
}

func (msg TargetedMsg) Action() interface{} {
	return msg.action
}

func TargetCmd(id int, action interface{}) tea.Cmd {
	return func() tea.Msg {
		return TargetedMsg{id, action}
	}
}

type DebugMsg struct {
	msg string
}

func (dbg DebugMsg) String() string {
	return dbg.msg
}

func DebugCmd(msg string) tea.Cmd {
	return func() tea.Msg {
		return DebugMsg{msg}
	}
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
