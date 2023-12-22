package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo"
)

func main() {
	type Btn1Msg struct{}
	// type StateUpdatedMsg struct{}

	backend := tuigo.Backend{
		States: []tuigo.AppState{"initial", "final"},
		Constructors: map[tuigo.AppState]tuigo.Constructor{
			"initial": func(tuigo.Window) tuigo.Window {
				container1 := tuigo.Container(
					true,
					tuigo.VerticalContainer,
					tuigo.Button("button1", tuigo.SimpleBtn, Btn1Msg{}),
					tuigo.RadioWithID(1, "radio1"),
					tuigo.Text("label1", tuigo.NormalText),
					tuigo.Input("input2", "<default>", "<placeholder>", tuigo.PathInput),
				)

				container2 := tuigo.Container(
					true,
					tuigo.VerticalContainer,
					tuigo.Text("text2", tuigo.DimmedText),
					tuigo.Button("hidden_button4", tuigo.SimpleBtn, nil).Hide(),
					tuigo.Selector([]string{"item1", "item2", "item3", "item4", "item5"}, false),
					tuigo.Text("text3", tuigo.DimmedText),
				)

				container3 := tuigo.Container(true, tuigo.HorizontalContainer, container1, container2)

				container := tuigo.Container(
					true,
					tuigo.VerticalContainer,
					tuigo.Text("label2", tuigo.NormalText),
					tuigo.Button("button6", tuigo.SimpleBtn, nil),
					tuigo.Selector([]string{"item1", "item2", "item3"}, true),
					tuigo.Input("input1", "<default>", "<placeholder>", tuigo.TextInput),
					tuigo.ButtonWithID(9, "button9", tuigo.AcceptBtn, nil),
					container3,
				)
				return container
			},
			"final": func(prev tuigo.Window) tuigo.Window {
				return tuigo.Container(
					true,
					tuigo.VerticalContainer,
					tuigo.Button("button9", tuigo.SimpleBtn, nil),
					tuigo.Input("input3", "<default>", "<placeholder>", tuigo.TextInput),
					tuigo.Radio("radio2"),
				)
			},
		},
		Updaters: map[tuigo.AppState]tuigo.Updater{
			"initial": func(window tuigo.Window, msg tea.Msg) (tuigo.Window, tea.Cmd) {
				// switch msg.(type) {
				// case Btn1Msg:
				// 	radio1_cont, radio1 := container.GetElementByID(1)
				// 	radio1 = radio1.(tuigo.RadioElement).Toggle()
				// 	radio1_cont
				// 	// radio1 = (*(radio1.(radio.Radio))).Toggle()
				// 	// button9 := container.GetElementByID(9)
				// 	return container, tuigo.Callback(StateUpdatedMsg{})
				// }
				return window, nil
			},
		},
		Finalizer: func(containers map[tuigo.AppState]tuigo.Window) tuigo.Window {
			return tuigo.Container(
				false, tuigo.VerticalContainer,
				tuigo.Text("app finalized", tuigo.NormalText),
			)
		},
	}
	p := tea.NewProgram(tuigo.App(backend, false))
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
