package radio

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

type TestMsg struct{}

func TestRadio(t *testing.T) {
	radio := New("test")
	{
		r1, _ := radio.Update(tea.KeyMsg{Type: tea.KeySpace})
		radio = *(r1.(*Model))
		if !radio.State() {
			t.Fatalf("radio did not toggle on space")
		}
		radio.Toggle()
		if radio.State() {
			t.Fatalf("radio did not toggle")
		}
	}
	{
		radio.Focus()
		if !radio.Focused() {
			t.Fatalf("radio did not focus")
		}
		radio.Blur()
		if radio.Focused() {
			t.Fatalf("radio did not blur")
		}
	}
}
