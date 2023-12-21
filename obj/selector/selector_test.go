package selector

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

type TestMsg struct{}

func TestSelector(t *testing.T) {
	selector := Selector{
		multiselect: true,
		cursor:      0,
		options:     []string{"option1", "option2", "option3"},
		selected:    map[string]struct{}{},
		disabled:    map[string]struct{}{},
	}
	{
		selector = selector.Next()
		selector = selector.Toggle()
		if len(selector.Selected()) != 1 || selector.Selected()[0] != "option2" {
			t.Fatalf("selector did not select option2")
		}
		selector = selector.Prev()
		selector = selector.Toggle()
		if len(selector.Selected()) != 2 || selector.Selected()[0] != "option1" || selector.Selected()[1] != "option2" {
			t.Fatalf("selector did not select option1 & option2")
		}
		selector = selector.Next()
		selector = selector.Toggle()
		if len(selector.Selected()) != 1 || selector.Selected()[0] != "option1" {
			t.Fatalf("selector did not deselect option2")
		}
		selector = selector.Prev()
		selector = selector.Disable("option2")
		selector = selector.Next()
		selector = selector.Toggle()
		if (len(selector.Selected()) != 2) || (selector.Selected()[0] != "option1") || (selector.Selected()[1] != "option3") || (selector.Cursor() != 2) {
			t.Fatalf("selector did not properly disable option2")
		}
		selector = selector.Enable("option2")
		selector = selector.Prev()
		selector = selector.Toggle()
		if (len(selector.Selected()) != 3) || (selector.Selected()[0] != "option1") || (selector.Selected()[1] != "option2") || (selector.Selected()[2] != "option3") || (selector.Cursor() != 1) {
			t.Fatalf("selector did not properly enable option2")
		}
	}
	{
		sel, _ := selector.Update(tea.KeyMsg{Type: tea.KeySpace})
		selector = sel.(Selector)
		if len(selector.Selected()) != 2 || selector.Selected()[0] != "option1" || selector.Selected()[1] != "option3" {
			t.Fatalf("selector did not deselect option2 with updater")
		}
		sel, _ = selector.Update(tea.KeyMsg{Type: tea.KeyUp})
		selector = sel.(Selector)
		if selector.Cursor() != 0 {
			t.Fatalf("selector did not move cursor up")
		}
		sel, _ = selector.Update(tea.KeyMsg{Type: tea.KeyDown})
		selector = sel.(Selector)
		if selector.Cursor() != 1 {
			t.Fatalf("selector did not move cursor down")
		}
	}
}

func TestSelectorToggle(t *testing.T) {
	newselector := Selector{
		multiselect: false,
		cursor:      0,
		options:     []string{"option1", "option2", "option3"},
		selected:    map[string]struct{}{},
		disabled:    map[string]struct{}{},
	}
	newselector = newselector.Toggle()
	newselector = newselector.Next()
	newselector = newselector.Next()
	newselector = newselector.Toggle()
	if len(newselector.Selected()) != 1 || newselector.Selected()[0] != "option3" {
		t.Fatalf("selector did not select option3 & deselect option1")
	}
}
