package selector

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

type TestMsg struct{}

func TestSelector(t *testing.T) {
	selector := New([]string{"option1", "option2", "option3"}, true)
	{
		selector.Next()
		selector.Toggle()
		if len(selector.Selected()) != 1 || selector.Selected()[0] != "option2" {
			t.Fatalf("selector did not select option2")
		}
		selector.Prev()
		selector.Toggle()
		if len(selector.Selected()) != 2 || selector.Selected()[0] != "option1" || selector.Selected()[1] != "option2" {
			t.Fatalf("selector did not select option1 & option2")
		}
		selector.Next()
		selector.Toggle()
		if len(selector.Selected()) != 1 || selector.Selected()[0] != "option1" {
			t.Fatalf("selector did not deselect option2")
		}
		selector.Prev()
		selector.Disable("option2")
		selector.Next()
		selector.Toggle()
		if (len(selector.Selected()) != 2) || (selector.Selected()[0] != "option1") || (selector.Selected()[1] != "option3") || (selector.Cursor() != 2) {
			t.Fatalf("selector did not properly disable option2")
		}
		selector.Enable("option2")
		selector.Prev()
		selector.Toggle()
		if (len(selector.Selected()) != 3) || (selector.Selected()[0] != "option1") || (selector.Selected()[1] != "option2") || (selector.Selected()[2] != "option3") || (selector.Cursor() != 1) {
			t.Fatalf("selector did not properly enable option2")
		}
	}
	{
		sel, _ := selector.Update(tea.KeyMsg{Type: tea.KeySpace})
		selector = *(sel.(*Model))
		if len(selector.Selected()) != 2 || selector.Selected()[0] != "option1" || selector.Selected()[1] != "option3" {
			t.Fatalf("selector did not deselect option2 with updater")
		}
		sel, _ = selector.Update(tea.KeyMsg{Type: tea.KeyUp})
		selector = *(sel.(*Model))
		if selector.Cursor() != 0 {
			t.Fatalf("selector did not move cursor up")
		}
		sel, _ = selector.Update(tea.KeyMsg{Type: tea.KeyDown})
		selector = *(sel.(*Model))
		if selector.Cursor() != 1 {
			t.Fatalf("selector did not move cursor down")
		}
	}
	{
		selector.Focus()
		if !selector.Focused() {
			t.Fatalf("selector did not focus")
		}
		selector.Blur()
		if selector.Focused() {
			t.Fatalf("selector did not blur")
		}
	}
	{
		newselector := New([]string{"option1", "option2", "option3"}, false)
		newselector.Toggle()
		newselector.Next()
		newselector.Next()
		newselector.Toggle()
		if len(newselector.Selected()) != 1 || newselector.Selected()[0] != "option3" {
			t.Fatalf("selector did not select option3 & deselect option1")
		}
	}
}
