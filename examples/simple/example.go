package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo"
)

func main() {
	type Btn1Msg struct{}
	type Btn9Msg struct{}
	type Inp1Msg struct{}

	backend := tuigo.Backend{
		States: []tuigo.AppState{"initial", "final"},
		Constructors: map[tuigo.AppState]tuigo.Constructor{
			"initial": func(tuigo.Window) tuigo.Window {
				container1 := tuigo.Container(
					tuigo.Focusable,
					tuigo.VerticalContainer,
					tuigo.Button("button1", tuigo.SimpleBtn, Btn1Msg{}),
					tuigo.RadioWithID(1, "radio1", tuigo.NoCallback),
					tuigo.Text("label1", tuigo.NormalText),
					tuigo.Input("input2", "<default>", "<placeholder>", tuigo.PathInput, tuigo.NoCallback),
				)

				container2 := tuigo.Container(
					tuigo.Focusable,
					tuigo.VerticalContainer,
					tuigo.Text("very long\nlist of items", tuigo.DimmedText),
					tuigo.Button("hidden_button4", tuigo.SimpleBtn, tuigo.NoCallback).Hide(),
					tuigo.SelectorWithID(
						5,
						[]string{
							"item1",
							"item2",
							"item3",
							"item4",
							"item5",
							"item6",
							"item7",
							"item8",
							"item9",
							"item10",
							"item11",
							"item12",
							"item13",
						},
						tuigo.SelectOne,
						5,
						tuigo.NoCallback,
					),
					tuigo.Text("text3", tuigo.DimmedText),
				)

				container3 := tuigo.Container(
					tuigo.Focusable,
					tuigo.HorizontalContainer,
					container1,
					container2,
				)
				hidden_container := tuigo.Container(
					tuigo.Focusable,
					tuigo.HorizontalContainer,
					tuigo.Button("hidden_button7", tuigo.SimpleBtn, tuigo.NoCallback).Hide(),
					tuigo.Text("hidden_text", tuigo.DimmedText).Hide(),
				)

				container := tuigo.Container(
					tuigo.Focusable,
					tuigo.VerticalContainer,
					tuigo.TextWithID(2, "label2", tuigo.NormalText),
					tuigo.Button("button6", tuigo.SimpleBtn, tuigo.NoCallback),
					tuigo.Selector(
						[]string{"item1", "item2", "item3"},
						tuigo.MultiSelect,
						tuigo.NoViewLimit,
						tuigo.NoCallback,
					),
					hidden_container,
					tuigo.InputWithID(3, "input1", "<default>", "<placeholder>", tuigo.TextInput, Inp1Msg{}),
					tuigo.ButtonWithID(9, "button9", tuigo.AcceptBtn, Btn9Msg{}),
					container3,
				)
				return container
			},
			"final": func(prev tuigo.Window) tuigo.Window {
				return tuigo.Container(
					tuigo.Focusable,
					tuigo.VerticalContainer,
					tuigo.Button("button9", tuigo.SimpleBtn, tuigo.NoCallback),
					tuigo.Input("input3", "<default>", "<placeholder>", tuigo.TextInput, tuigo.NoCallback),
					tuigo.Radio("radio2", tuigo.NoCallback),
				)
			},
		},
		Updaters: map[tuigo.AppState]tuigo.Updater{
			"initial": func(window tuigo.Window, msg tea.Msg) (tuigo.Window, tea.Cmd) {
				toggle_radio1 := tuigo.TgtCmd(
					1,
					func(cont tuigo.Wrapper, radio tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
						return cont, radio.(tuigo.RadioElement).Toggle()
					})
				en_dis_it3_in_sel5 := tuigo.TgtCmd(
					5,
					func(cont tuigo.Wrapper, selector tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
						el := selector.(tuigo.SelectorElement)
						if el.Disabled("item3") {
							return cont, el.Enable("item3")
						} else {
							return cont, el.Disable("item3")
						}
					})

				switch msg.(type) {
				case Btn1Msg:
					return window, toggle_radio1
				case Btn9Msg:
					return window, tea.Batch(toggle_radio1, en_dis_it3_in_sel5)
				case Inp1Msg:
					typed_text := window.GetElementByID(3).Data().(string)
					return window, tuigo.TgtCmd(
						2,
						func(cont tuigo.Wrapper, text tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
							el := text.(tuigo.TextElement)
							return cont, el.Set("you typed: " + typed_text)
						})
				}
				return window, nil
			},
		},
		Finalizer: func(containers map[tuigo.AppState]tuigo.Window) tuigo.Window {
			return tuigo.Container(
				tuigo.NonFocusable, tuigo.VerticalContainer,
				tuigo.Text("app finalized", tuigo.NormalText),
			)
		},
	}
	p := tea.NewProgram(tuigo.App(backend, true))
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
