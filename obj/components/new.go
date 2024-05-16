package component

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/obj/components/button"
	"github.com/haykh/tuigo/obj/components/input"
	"github.com/haykh/tuigo/obj/components/radio"
	"github.com/haykh/tuigo/obj/components/selector"
	"github.com/haykh/tuigo/obj/components/text"
	"github.com/haykh/tuigo/obj/container"
	"github.com/haykh/tuigo/utils"
)

func NewInput(id int, label, def, placeholder string, inputtype utils.InputType, callback tea.Msg) container.SimpleContainer {
	return container.NewSimpleContainer(true, input.New(id, label, def, placeholder, inputtype, callback))
}

func NewButton(id int, label string, btntype utils.ButtonType, callback tea.Msg) container.SimpleContainer {
	return container.NewSimpleContainer(true, button.New(id, label, btntype, callback))
}

func NewRadio(id int, label string, callback tea.Msg) container.SimpleContainer {
	return container.NewSimpleContainer(true, radio.New(id, label, callback))
}

func NewSelector(id int, options []string, multiselect bool, view_limit int, callback tea.Msg) container.SimpleContainer {
	return container.NewSimpleContainer(true, selector.New(id, options, multiselect, view_limit, callback))
}

func NewText(id int, txt string, texttype utils.TextType) container.SimpleContainer {
	return container.NewSimpleContainer(false, text.New(id, txt, texttype))
}
