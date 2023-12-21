package container

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/utils"
)

type TestMsg struct{}

func TestContainer(t *testing.T) {
	container := New(true, utils.SimpleContainer)
	{
		container = container.Focus().(Container)
		if !container.Focused() {
			t.Errorf("expected container to be focused")
		}
		container = container.Blur().(Container)
		if container.Focused() {
			t.Errorf("expected container to be unfocused")
		}
	}
	container.elements = append(container.elements, New(true, utils.SimpleContainer))
	{
		container = container.Focus().(Container)
		subcontainer := container.Containers()[0].(Container)
		if !subcontainer.Focused() {
			t.Errorf("expected subcontainer to be focused")
		}
		container = container.Blur().(Container)
		subcontainer = container.Containers()[0].(Container)
		if subcontainer.Focused() {
			t.Errorf("expected subcontainer to be unfocused")
		}
	}
	container.elements = append(container.elements, New(true, utils.SimpleContainer))
	{
		container = container.Focus().(Container)
		subcontainer0 := container.Containers()[0].(Container)
		if !subcontainer0.Focused() {
			t.Errorf("expected subcontainer #0 to be focused")
		}
		newc, _ := container.FocusNext()
		container = newc.(Container)
		subcontainer0 = container.Containers()[0].(Container)
		subcontainer1 := container.Containers()[1].(Container)
		if !subcontainer1.Focused() {
			t.Errorf("expected subcontainer #1 to be focused")
		}
		if subcontainer0.Focused() {
			t.Errorf("expected subcontainer #0 to be unfocused")
		}
		newc, _ = container.FocusPrev()
		container = newc.(Container)
		subcontainer0 = container.Containers()[0].(Container)
		subcontainer1 = container.Containers()[1].(Container)
		if !subcontainer0.Focused() {
			t.Errorf("expected subcontainer #0 to be focused")
		}
		if subcontainer1.Focused() {
			t.Errorf("expected subcontainer #1 to be unfocused")
		}
	}
	{
		newc, cmd := container.FocusNext()
		container = newc.(Container)
		switch cmd().(type) {
		case utils.FocusNextMsg:
			t.Errorf("expected cmd to be not nil")
		}
		newc, cmd = container.FocusNext()
		container = newc.(Container)
		if cmd == nil {
			t.Errorf("expected cmd to be not nil")
		}
		correct := false
		switch cmd().(type) {
		case utils.FocusNextMsg:
			correct = true
		}
		if !correct {
			t.Errorf("expected cmd to be FocusNextMsg")
		}
		component0 := container.Containers()[0].(Container)
		component1 := container.Containers()[1].(Container)
		if component0.Focused() {
			t.Errorf("expected component #0 to be unfocused")
		}
		if !component1.Focused() {
			t.Errorf("expected component #1 to be focused")
		}
		container = container.Blur().(Container)
	}
	{
		newc := container.Focus()
		container = newc.(Container)
		component0 := container.Containers()[0].(Container)
		component1 := container.Containers()[1].(Container)
		if !component0.Focused() {
			t.Errorf("expected component #0 to be focused")
		}
		if component1.Focused() {
			t.Errorf("expected component #1 to be unfocused")
		}
		var cmd tea.Cmd
		newc, cmd = container.FocusPrev()
		container = newc.(Container)
		if cmd == nil {
			t.Errorf("expected cmd to be not nil")
		}
		correct := false
		switch cmd().(type) {
		case utils.FocusPrevMsg:
			correct = true
		}
		if !correct {
			t.Errorf("expected cmd to be FocusPrevMsg")
		}
	}
}
