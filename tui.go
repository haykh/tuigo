package tuigo

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/app"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/obj/components/button"
	"github.com/haykh/tuigo/obj/components/input"
	"github.com/haykh/tuigo/obj/components/radio"
	"github.com/haykh/tuigo/obj/components/selector"
	"github.com/haykh/tuigo/obj/components/text"
	"github.com/haykh/tuigo/obj/container"
	"github.com/haykh/tuigo/utils"
)

// aliases
type AppState = app.AppState
type Window = app.Window
type Constructor = app.Constructor
type Updater = app.Updater
type Backend = app.Backend
type Accessor = obj.Accessor
type Element = obj.Element
type Wrapper = container.Wrapper

// constructors
var App = app.New

var TextWithID = text.New
var RadioWithID = radio.New
var InputWithID = input.New
var ButtonWithID = button.New
var SelectorWithID = selector.New

var Container = container.NewComplexContainer

var Text = func(txt string, txttype TextType) container.SimpleContainer {
	return TextWithID(-1, txt, txttype)
}
var Radio = func(lbl string, callback tea.Msg) container.SimpleContainer {
	return RadioWithID(-1, lbl, callback)
}
var Input = func(lbl, def, plc string, inptype InputType, callback tea.Msg) container.SimpleContainer {
	return InputWithID(-1, lbl, def, plc, inptype, callback)
}
var Button = func(lbl string, btntype ButtonType, callback tea.Msg) container.SimpleContainer {
	return ButtonWithID(-1, lbl, btntype, callback)
}
var Selector = func(opt []string, mult bool, callback tea.Msg) container.SimpleContainer {
	return SelectorWithID(-1, opt, mult, callback)
}

// components & elements
type TextElement = text.Text
type RadioElement = radio.Radio
type InputElement = input.Input
type ButtonElement = button.Button
type SelectorElement = selector.Selector

type ComplexContainerElement = container.ComplexContainer
type SimpleContainerElement = container.SimpleContainer

// component types
type ContainerType = utils.ContainerType

var VerticalContainer = utils.VerticalContainer
var HorizontalContainer = utils.HorizontalContainer

type ButtonType = utils.ButtonType

var SimpleBtn = utils.SimpleBtn
var AcceptBtn = utils.AcceptBtn
var ControlBtn = utils.ControlBtn

type InputType = utils.InputType

var PathInput = utils.PathInput
var TextInput = utils.TextInput

type TextType = utils.TextType

var NormalText = utils.NormalText
var LabelText = utils.LabelText
var DimmedText = utils.DimmedText

// messages
var Callback = utils.Callback
var DbgCmd = utils.DebugCmd
var TgtCmd = utils.TargetCmd
