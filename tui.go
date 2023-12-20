package tuigo

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/app"
	"github.com/haykh/tuigo/debug"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/obj/button"
	"github.com/haykh/tuigo/obj/container"
	"github.com/haykh/tuigo/obj/input"
	"github.com/haykh/tuigo/obj/radio"
	"github.com/haykh/tuigo/obj/selector"
	"github.com/haykh/tuigo/obj/text"
	"github.com/haykh/tuigo/utils"
)

func NewApp(container obj.Element, enable_debug bool) app.App {
	dbg := debug.New()
	if enable_debug {
		dbg.Enable()
	}
	app := app.App{
		Container: container.(obj.Collection).Focus(),
		Debugger:  dbg,
	}
	return app
}

// Components

func NewContainer(focusable bool, containerType utils.ContainerType, elements ...obj.Element) obj.Element {
	return container.NewContainer(focusable, containerType, elements...)
}

func NewText(txt string, texttype utils.TextType) obj.Element {
	return text.New(txt, texttype)
}

func NewRadio(label string) obj.Element {
	return radio.New(label)
}

func NewInput(label, def, placeholder string, inputtype utils.InputType) obj.Element {
	return input.New(label, def, placeholder, inputtype)
}

func NewButton(label string, btnType utils.ButtonType, msg tea.Msg) obj.Element {
	return button.New(label, btnType, msg)
}

func NewSelector(options []string, multiselect bool) obj.Element {
	return selector.New(options, multiselect)
}
