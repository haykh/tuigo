package radio

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/obj"
)

type TestMsg struct{}

func TestRadio(t *testing.T) {
	radio := Radio{
		obj.NewElementWithID(0), "test", false,
	}
	{
		r1, _ := radio.Update(tea.KeyMsg{Type: tea.KeySpace})
		radio = r1.(Radio)
		if state, ok := radio.Data().(bool); ok {
			if !state {
				t.Fatalf("radio did not capture space key")
			}
		} else {
			t.Fatalf("radio data is not bool")
		}
		radio = radio.Toggle()
		if state, ok := radio.Data().(bool); ok {
			if state {
				t.Fatalf("radio did not toggle")
			}
		} else {
			t.Fatalf("radio data is not bool")
		}
	}
}
