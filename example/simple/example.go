package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo"
	"github.com/haykh/tuigo/app"
	"github.com/haykh/tuigo/obj"
	"github.com/haykh/tuigo/utils"
)

func main() {
	backend := app.Backend{
		States: []app.AppState{"initial", "final"},
		Constructors: map[app.AppState]app.Constructor{
			"initial": func(obj.Element) obj.Element {
				container1 := tuigo.NewContainer(
					true,
					utils.VerticalContainer,
					tuigo.NewButton("button1", utils.SimpleBtn, nil),
					tuigo.NewRadio("radio1"),
					tuigo.NewText("label1", utils.NormalText),
					tuigo.NewInput("input2", "<default>", "<placeholder>", utils.PathInput),
				)

				container2 := tuigo.NewContainer(
					true,
					utils.VerticalContainer,
					tuigo.NewText("text2", utils.DimmedText),
					tuigo.NewButton("button4", utils.SimpleBtn, nil),
					tuigo.NewSelector([]string{"item1", "item2", "item3", "item4", "item5"}, false),
				)

				container3 := tuigo.NewContainer(true, utils.HorizontalContainer, container1, container2)

				container := tuigo.NewContainer(
					true,
					utils.VerticalContainer,
					tuigo.NewButton("button6", utils.SimpleBtn, nil),
					tuigo.NewSelector([]string{"item1", "item2", "item3"}, true),
					tuigo.NewInput("input1", "<default>", "<placeholder>", utils.TextInput),
					tuigo.NewButton("button9", utils.AcceptBtn, nil),
					container3,
				)
				return container
			},
			"final": func(prev obj.Element) obj.Element {
				return tuigo.NewContainer(
					true,
					utils.VerticalContainer,
					tuigo.NewButton("button9", utils.SimpleBtn, nil),
					tuigo.NewInput("input3", "<default>", "<placeholder>", utils.TextInput),
					tuigo.NewRadio("radio2"),
				)
			},
		},
		Finalizer: func(containers map[app.AppState]obj.Element) {
			fmt.Println("Finalizer Called")
		},
	}
	p := tea.NewProgram(tuigo.NewApp(backend, true))
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
