package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo"
	"github.com/haykh/tuigo/component/button"
	"github.com/haykh/tuigo/component/multiselect"
	"github.com/haykh/tuigo/component/pathinput"
	"github.com/haykh/tuigo/component/radio"
	"github.com/haykh/tuigo/utils"
)

type testState struct {
	label string
	next  *testState
	prev  *testState
}

func (ts testState) Label() string {
	return ts.label
}

func (ts testState) Next() utils.State {
	return *ts.next
}

func (ts testState) Prev() utils.State {
	return *ts.prev
}

func newTestStates() []testState {
	window1 := testState{label: "window1"}
	window2 := testState{label: "window2"}
	window3 := testState{label: "window3"}
	window1.next = &window2
	window2.next = &window3
	window3.next = &window1
	window1.prev = &window3
	window2.prev = &window1
	window3.prev = &window2
	return []testState{window1, window2, window3}
}

var TestStates []testState = newTestStates()

func newTestFields() map[utils.State]tuigo.Field {
	flds := map[utils.State]tuigo.Field{}
	selector1 := multiselect.New([]string{"option1", "option2", "option3"})
	selector2 := multiselect.New([]string{"option4", "option5", "option6"})
	selector3 := multiselect.New([]string{"option7", "option8", "option9"})
	radio1 := radio.New("radio1")
	flds[TestStates[0]] = tuigo.NewField("window1", true, false).
		AddElement(&selector1).
		AddElement(&selector2).
		AddElement(&radio1).
		AddElement(&selector3)

	btn1 := button.New("click me", utils.SimpleBtn, nil)
	radio2 := radio.New("radio2")

	flds[TestStates[1]] = tuigo.NewField("window2", false, false).
		AddElement(&btn1).
		AddElement(&radio2)

	btn2 := button.New("yes babe", utils.SimpleBtn, nil)

	pathinput1 := pathinput.New("source path", "$HOME/", "<default>")

	flds[TestStates[2]] = tuigo.NewField("window3", false, true).AddElement(&btn2).AddElement(&pathinput1)
	return flds
}

var TestFields map[utils.State]tuigo.Field = newTestFields()

func main() {
	initialState := TestStates[0]
	allFields := TestFields
	p := tea.NewProgram(tuigo.NewApp(initialState, allFields, true))
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
