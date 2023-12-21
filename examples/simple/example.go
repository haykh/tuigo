package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo"
)

func main() {
	backend := tuigo.Backend{
		States: []tuigo.AppState{"initial", "final"},
		Constructors: map[tuigo.AppState]tuigo.Constructor{
			"initial": func(tuigo.Element) tuigo.Element {
				container1 := tuigo.Container(
					true,
					tuigo.VerticalContainer,
					tuigo.Button("button1", tuigo.SimpleBtn, nil),
					tuigo.Radio("radio1"),
					tuigo.Text("label1", tuigo.NormalText),
					tuigo.Input("input2", "<default>", "<placeholder>", tuigo.PathInput),
				)

				container2 := tuigo.Container(
					true,
					tuigo.VerticalContainer,
					tuigo.Text("text2", tuigo.DimmedText),
					tuigo.Button("button4", tuigo.SimpleBtn, nil),
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
					tuigo.Button("button9", tuigo.AcceptBtn, nil),
					container3,
				)
				return container
			},
			"final": func(prev tuigo.Element) tuigo.Element {
				return tuigo.Container(
					true,
					tuigo.VerticalContainer,
					tuigo.Button("button9", tuigo.SimpleBtn, nil),
					tuigo.Input("input3", "<default>", "<placeholder>", tuigo.TextInput),
					tuigo.Radio("radio2"),
				)
			},
		},
		Finalizer: func(containers map[tuigo.AppState]tuigo.Element) {
			fmt.Println("Finalizer Called")
		},
	}
	p := tea.NewProgram(tuigo.App(backend, true))
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
