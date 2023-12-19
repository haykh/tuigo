package input

import (
	"testing"

	"github.com/haykh/tuigo/utils"
)

type TestMsg struct{}

func TestPathInput(t *testing.T) {
	input := New("test", "<default>", "<placeholder>", utils.TextInput)
	{
		input.Focus()
		if !input.Focused() {
			t.Fatalf("input did not focus")
		}
		input.Blur()
		if input.Focused() {
			t.Fatalf("input did not blur")
		}
	}
	{
		if input.Model.View() != "test<default> " {
			t.Fatalf("input did not set default value")
		}
		input.Model.SetValue("test_dir")
		if input.Value() != "test_dir" {
			t.Fatalf("input did not set value")
		}
		input.Model.Reset()
		if input.Model.View() != "test<placeholder>" {
			t.Fatalf("input did not reset value")
		}
	}
}
