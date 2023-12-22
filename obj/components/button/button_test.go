package button

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/utils"
)

type TestMsg struct{}

func TestButton(t *testing.T) {
	button := Button{
		ElementWithID:       obj.NewElementWithID(0),
		ElementWithCallback: obj.NewElementWithCallback(TestMsg{}),
		label:               "test",
		btntype:             utils.SimpleBtn,
	}
	{
		btn, cmd := button.Update(tea.KeyMsg{Type: tea.KeySpace})
		if cmd == nil || !utils.CheckCmd(cmd, TestMsg{}) {
			t.Fatalf("button did not capture space key")
		}
		button = btn.(Button)
	}
	{
		btn, cmd := button.Update(tea.KeyMsg{Type: tea.KeyEnter})
		if cmd == nil || !utils.CheckCmd(cmd, TestMsg{}) {
			t.Fatalf("button did not capture enter key")
		}
		button = btn.(Button)
	}
	{
		if npresses, ok := button.Data().(int); ok {
			if npresses != 2 {
				t.Fatalf("button did not capture enter key")
			}
		} else {
			t.Fatalf("button data is not int")
		}
	}
}
