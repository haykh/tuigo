package input

import (
	"testing"

	"github.com/haykh/tuigo/utils"
)

type TestMsg struct{}

func TestPathInput(t *testing.T) {
	input := Input{
		model:     NewTextinputModel("test", "<default>", "<placeholder>"),
		inputtype: utils.TextInput,
	}
	{
		if input.model.View() != "test<default> " {
			t.Fatalf("input did not set default value")
		}
		input.model.SetValue("test_dir")
		if val, ok := input.Data().(string); ok {
			if val != "test_dir" {
				t.Fatalf("input did not set value")
			}
		} else {
			t.Fatalf("input data is not string")
		}
		input.model.Reset()
		if input.model.View() != "test<placeholder>" {
			t.Fatalf("input did not reset value")
		}
	}
}
