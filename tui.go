package tuigo

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo/app"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/obj/button"
	"github.com/haykh/tuigo/obj/container"
	"github.com/haykh/tuigo/obj/input"
	"github.com/haykh/tuigo/obj/radio"
	"github.com/haykh/tuigo/obj/selector"
	"github.com/haykh/tuigo/obj/text"
	"github.com/haykh/tuigo/utils"
)

// aliases
type AppState = app.AppState
type Constructor = app.Constructor
type Backend = app.Backend
type Element = obj.Element
type Collection = obj.Collection

// constructors
var App = app.New

var TextWithID = text.New
var RadioWithID = radio.New
var InputWithID = input.New
var ButtonWithID = button.New
var SelectorWithID = selector.New

var Container = container.New

var Text = func(txt string, txttype TextType) Element {
	return TextWithID(-1, txt, txttype)
}
var Radio = func(lbl string) Element {
	return RadioWithID(-1, lbl)
}
var Input = func(lbl, def, plc string, inptype InputType) Element {
	return InputWithID(-1, lbl, def, plc, inptype)
}
var Button = func(lbl string, btntype ButtonType, act tea.Msg) Element {
	return ButtonWithID(-1, lbl, btntype, act)
}
var Selector = func(opt []string, mult bool) Element {
	return SelectorWithID(-1, opt, mult)
}

// component types
type ContainerType = utils.ContainerType

var SimpleContainer = utils.SimpleContainer
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
