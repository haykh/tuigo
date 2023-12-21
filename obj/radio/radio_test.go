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
		if !radio.State() {
			t.Fatalf("radio did not toggle on space")
		}
		radio = radio.Toggle()
		if radio.State() {
			t.Fatalf("radio did not toggle")
		}
	}
}
