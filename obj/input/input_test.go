package input

import (
	"testing"

	"github.com/haykh/tuigo/utils"
)

type TestMsg struct{}

func TestPathInput(t *testing.T) {
	input := Model{
		model:     NewTextinput("test", "<default>", "<placeholder>"),
		inputtype: utils.TextInput,
	}
	{
		if input.model.View() != "test<default> " {
			t.Fatalf("input did not set default value")
		}
		input.model.SetValue("test_dir")
		if input.Value() != "test_dir" {
			t.Fatalf("input did not set value")
		}
		input.model.Reset()
		if input.model.View() != "test<placeholder>" {
			t.Fatalf("input did not reset value")
		}
	}
}
