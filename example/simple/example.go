package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo"
	"github.com/haykh/tuigo/utils"
)

func main() {
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

	p := tea.NewProgram(tuigo.NewApp(container, true))
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
