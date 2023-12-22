package container

import (
	"fmt"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/utils"
)

type element struct{}

var _ obj.Element = (*element)(nil)

func (e element) View(bool) string {
	return "dummy"
}

func (e element) Update(tea.Msg) (obj.Element, tea.Cmd) {
	return e, nil
}

func TestSimpleContainer(t *testing.T) {
	var new_element = func() element {
		return element{}
	}

	var new_simple_container = func() SimpleContainer {
		return NewSimpleContainer(true, new_element())
	}

	var focus = func(c SimpleContainer) SimpleContainer {
		return c.Focus().(SimpleContainer)
	}

	var focus_complex = func(c ComplexContainer) ComplexContainer {
		return c.Focus().(ComplexContainer)
	}

	var blur = func(c SimpleContainer) SimpleContainer {
		return c.Blur().(SimpleContainer)
	}

	var focus_test = func(t *testing.T, c SimpleContainer, want bool) {
		if c.Focused() != want {
			t.Errorf("expected container to be focused")
		}
	}

	var focus_complex_test = func(t *testing.T, c ComplexContainer, want bool, ind int) {
		sc := c.Components()[ind].(SimpleContainer)
		if sc.Focused() != want {
			t.Errorf(fmt.Sprintf("expected container[%d] to be focused", ind))
		}
	}

	container := new_simple_container()
	container = focus(container)
	focus_test(t, container, true)
	container = blur(container)
	focus_test(t, container, false)
	if container.Element().(element) != (element{}) {
		t.Errorf("expected container to have element")
	}

	complex_container := add_components(new_complex_container(),
		new_simple_container(), new_simple_container(),
	)
	complex_container = focus_complex(complex_container)
	complex_container = focus_next(complex_container)
	nelem_test(t, complex_container, 2)
	focus_complex_test(t, complex_container, false, 0)
	focus_complex_test(t, complex_container, true, 1)
	complex_container = focus_prev(complex_container)
	focus_complex_test(t, complex_container, false, 1)
	focus_complex_test(t, complex_container, true, 0)
}

func TestSimpleContainerNested(t *testing.T) {
	var new_element = func() element {
		return element{}
	}

	var new_simple_container = func() SimpleContainer {
		return NewSimpleContainer(true, new_element())
	}

	var focus_complex = func(c ComplexContainer) ComplexContainer {
		return c.Focus().(ComplexContainer)
	}

	var focus_complex_test = func(t *testing.T, c ComplexContainer, want bool, ind ...int) {
		var sc SimpleContainer
		ind_str := ""
		for ii, i := range ind {
			ind_str += fmt.Sprintf("[%d]", i)
			if ii < len(ind)-1 {
				c = c.Components()[i].(ComplexContainer)
			} else {
				sc = c.Components()[i].(SimpleContainer)
			}
		}
		if sc.Focused() != want {
			t.Errorf(fmt.Sprintf("expected container[%s] to be focused", ind_str))
		}
	}
	complex_container := add_components(
		new_complex_container(), // < parent
		new_simple_container(),
		new_simple_container(),
		new_simple_container(),
		add_components(
			NewComplexContainer(true, utils.HorizontalContainer),
			new_simple_container(),
			new_simple_container(),
		),
	)
	complex_container = focus_complex(complex_container)
	focus_complex_test(t, complex_container, true, 0)
	for i := 0; i < 10; i++ {
		complex_container = focus_next(complex_container)
	}
	focus_complex_test(t, complex_container, false, 0)
	focus_complex_test(t, complex_container, true, 3, 1)
	for i := 0; i < 10; i++ {
		complex_container = focus_prev(complex_container)
	}
	focus_complex_test(t, complex_container, true, 0)
	focus_complex_test(t, complex_container, false, 3, 1)
}
