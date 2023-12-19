# `tuigo` 

## a terminal UI framework written in Go using the `bubbletea` library.

see `example/` for an example usage of `tuigo`. the scheme below shows roughly the structure of the API.

```mermaid
%%{
  init: {
    'theme': 'base', 
    'themeVariables': { 
        'fontFamily': 'monospace',
        'primaryColor': '#4C2FAD',
        'primaryTextColor': '#FFFFFF',
        'lineColor': '#E840E0',
        'primaryBorderColor': '#E840E0'
      }
    }
}%%

classDiagram
class tuigo["tuigo"]
class tui_utils["tuigo/utils"]
class tui_field["tuigo"]
class tui_component["tuigo"]

tuigo <-- tui_utils : initial state
tuigo <-- tui_field : mapping from states to fields
tuigo : NewApp(utils.State, map[utils.State]Field, bool) App
tuigo : App

tui_utils : Label() string
tui_utils : Next() utils.State
tui_utils : Prev() utils.State
tui_utils : utils.State

tui_field : NewField(string, bool, bool) Field
tui_field : Field

tui_component : Field.AddElement(&component) Field
tui_component : NewButton(string, utils.ButtonType, tea.Msg) button.Model
tui_component : NewSelector([]string, bool) selector.Model
tui_component : NewRadio(string) radio.Model
tui_component : NewPathInput(string, string, string) button.Model

tui_field <-- tui_component : populate the field with interactive components
```

## TODO

- [ ] customizable theme
- [ ] more components
- [ ] grid structure
- [ ] key help menu
- [ ] field validator