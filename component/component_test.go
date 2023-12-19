package component

import (
	"testing"

	"github.com/charmbracelet/bubbles/textinput"
)

type TestMsg struct{}

func TestComponent(t *testing.T) {
	{
		comp := Component{}
		comp.Focus()
		if !comp.Focused() {
			t.Errorf("expected component to be focused")
		}
		comp.Blur()
		if comp.Focused() {
			t.Errorf("expected component to be unfocused")
		}
	}
	{
		ti := TextInputWrap{Model: textinput.New()}
		ti.Focus()
		if !ti.Focused() {
			t.Errorf("expected textinput to be focused")
		}
		ti.Blur()
		if ti.Focused() {
			t.Errorf("expected textinput to be unfocused")
		}
	}
}
