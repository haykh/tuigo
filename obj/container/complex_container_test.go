package container

import (
	"fmt"
	"testing"

	"github.com/haykh/tuigo/utils"
)

var new_complex_container = func() ComplexContainer {
	return NewComplexContainer(true, utils.VerticalContainer)
}

var new_unfocusable_complex_container = func() ComplexContainer {
	return NewComplexContainer(false, utils.VerticalContainer)
}

var new_hidden_complex_container = func() ComplexContainer {
	return NewComplexContainer(true, utils.VerticalContainer).Hide().(ComplexContainer)
}

var get_child = func(c ComplexContainer, ind ...int) ComplexContainer {
	for _, i := range ind {
		c = c.Components()[i].(ComplexContainer)
	}
	return c
}

var add_components = func(c ComplexContainer, elements ...Component) ComplexContainer {
	return c.AddComponents(elements...).(ComplexContainer)
}

var focus_next = func(c ComplexContainer) ComplexContainer {
	newc, _ := c.FocusNext()
	return newc.(ComplexContainer)
}

var focus_prev = func(c ComplexContainer) ComplexContainer {
	newc, _ := c.FocusPrev()
	return newc.(ComplexContainer)
}

var nelem_test = func(t *testing.T, c ComplexContainer, want int, ind ...int) {
	c = get_child(c, ind...)
	ind_str := ""
	for _, i := range ind {
		ind_str += fmt.Sprintf("[%d]", i)
	}
	if len(c.Components()) != want {
		t.Errorf(fmt.Sprintf("expected container%s to have %d components", ind_str, want))
	}
}

func TestComplexContainer(t *testing.T) {
	var focus = func(c ComplexContainer) ComplexContainer {
		return c.Focus().(ComplexContainer)
	}

	var blur = func(c ComplexContainer) ComplexContainer {
		return c.Blur().(ComplexContainer)
	}

	var focus_test = func(t *testing.T, c ComplexContainer, want bool, ind ...int) {
		c = get_child(c, ind...)
		ind_str := ""
		for _, i := range ind {
			ind_str += fmt.Sprintf("[%d]", i)
		}
		if c.Focused() != want {
			t.Errorf(fmt.Sprintf("expected container%s to be focused", ind_str))
		}
	}

	{
		container := new_complex_container()
		container = focus(container)
		focus_test(t, container, true)
		container = blur(container)
		focus_test(t, container, false)
	}
	{
		container := add_components(
			new_complex_container(), // < parent
			add_components(
				new_complex_container(), // < parent
				add_components(
					new_complex_container(), // < parent
					new_complex_container(),
					new_unfocusable_complex_container(),
					new_complex_container(),
				),
				new_complex_container(),
				add_components(
					new_complex_container(), // < parent
					new_complex_container(),
					new_complex_container(),
					new_hidden_complex_container(),
				),
			),
			new_complex_container(),
			new_complex_container(),
		)
		nelem_test(t, container, 3)
		nelem_test(t, container, 3, 0)
		nelem_test(t, container, 3, 0, 0)
		nelem_test(t, container, 3, 0, 2)
		container = focus(container)
		focus_test(t, container, true, 0)
		focus_test(t, container, true, 0, 0)
		focus_test(t, container, true, 0, 0, 0)
		container = focus_next(container)
		focus_test(t, container, false, 0, 0, 0)
		focus_test(t, container, true, 0, 0, 2)
		container = focus_next(container)
		focus_test(t, container, false, 0, 0)
		focus_test(t, container, true, 0, 1)
		container = focus_prev(container)
		focus_test(t, container, false, 0, 0, 0)
		focus_test(t, container, true, 0, 0, 2)
		container = focus_prev(container)
		container = focus_prev(container)
		focus_test(t, container, true, 0)
		focus_test(t, container, true, 0, 0)
		focus_test(t, container, true, 0, 0, 0)
		for i := 0; i < 10; i++ {
			container = focus_next(container)
		}
		focus_test(t, container, true, 2)
	}
}
