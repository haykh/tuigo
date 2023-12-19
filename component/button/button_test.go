package button

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/utils"
)

type TestMsg struct{}

func TestButton(t *testing.T) {
	button := New("test", utils.SimpleBtn, TestMsg{})
	{
		_, cmd := button.Update(tea.KeyMsg{Type: tea.KeySpace})
		if cmd == nil || !utils.CheckCmd(cmd, TestMsg{}) {
			t.Fatalf("button did not capture space key")
		}
	}
	{
		_, cmd := button.Update(tea.KeyMsg{Type: tea.KeyEnter})
		if cmd == nil || !utils.CheckCmd(cmd, TestMsg{}) {
			t.Fatalf("button did not capture enter key")
		}
	}
	{
		button.Focus()
		if !button.Focused() {
			t.Fatalf("button did not focus")
		}
		button.Blur()
		if button.Focused() {
			t.Fatalf("button did not blur")
		}
	}
}
