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
				container1 := tuigo.NewContainer(
					true,
					tuigo.VerticalContainer,
					tuigo.NewButton("button1", tuigo.SimpleBtn, nil),
					tuigo.NewRadio("radio1"),
					tuigo.NewText("label1", tuigo.NormalText),
					tuigo.NewInput("input2", "<default>", "<placeholder>", tuigo.PathInput),
				)

				container2 := tuigo.NewContainer(
					true,
					tuigo.VerticalContainer,
					tuigo.NewText("text2", tuigo.DimmedText),
					tuigo.NewButton("button4", tuigo.SimpleBtn, nil),
					tuigo.NewSelector([]string{"item1", "item2", "item3", "item4", "item5"}, false),
				)

				container3 := tuigo.NewContainer(true, tuigo.HorizontalContainer, container1, container2)

				container := tuigo.NewContainer(
					true,
					tuigo.VerticalContainer,
					tuigo.NewButton("button6", tuigo.SimpleBtn, nil),
					tuigo.NewSelector([]string{"item1", "item2", "item3"}, true),
					tuigo.NewInput("input1", "<default>", "<placeholder>", tuigo.TextInput),
					tuigo.NewButton("button9", tuigo.AcceptBtn, nil),
					container3,
				)
				return container
			},
			"final": func(prev tuigo.Element) tuigo.Element {
				return tuigo.NewContainer(
					true,
					tuigo.VerticalContainer,
					tuigo.NewButton("button9", tuigo.SimpleBtn, nil),
					tuigo.NewInput("input3", "<default>", "<placeholder>", tuigo.TextInput),
					tuigo.NewRadio("radio2"),
				)
			},
		},
		Finalizer: func(containers map[tuigo.AppState]tuigo.Element) {
			fmt.Println("Finalizer Called")
		},
	}
	p := tea.NewProgram(tuigo.NewApp(backend, true))
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
