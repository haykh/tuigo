package tuigo

import (
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

// constructors
var NewApp = app.New

var NewContainer = container.New
var NewText = text.New
var NewRadio = radio.New
var NewInput = input.New
var NewButton = button.New
var NewSelector = selector.New

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
