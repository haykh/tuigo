# `tuigo` 

## a terminal UI framework written in Go using the `bubbletea` library.

see [`example/`](example/) for an example usage of `tuigo`. the scheme below shows roughly the structure of the API.

```mermaid
%%{
  init: {
    'theme': 'base',
    'themeVariables': {
        'fontFamily': 'JetBrainsMono Nerd Font, BlexMono Nerd Font, Roboto Mono, Source Code Pro, monospace',
        'primaryColor': '#4C2FAD',
        'primaryTextColor': '#FFFFFF',
        'lineColor': '#E840E0',
        'primaryBorderColor': '#E840E0'
      }
    }
}%%

classDiagram
  class tuigo {
    <<package>>
    NewContainer :: Func[bool, ContainerType, ...Element] -> Element
    NewButton :: Func[string, ButtonType, Msg] -> Element
    NewSelector :: Func[List~string~] -> Element
    NewInput :: Func[string, string, string, InputType] -> Element
    NewRadio :: Func[string] -> Element
  }

  class Backend["app.Backend"]{
    States :: List~AppState~
    Constructors :: Map[AppState]~Func[Element]->Element~
    Finalizer :: Func[Map[AppState]~Element~]
  }

  class NewApp["tuigo"] {
    <<package>>
    NewApp :: Func[Backend, bool] -> App
  }
  note for Backend "AppState = string"

  class tea["tea 'github.com/charmbracelet/bubbletea'"] {
    <<package>>
    NewProgram :: Func[Model, ...ProgramOption] -> *Program
  }

  tuigo --o Backend
  NewApp --o tea
  Backend --o NewApp

```


## TODO

- [x] app backend
- [x] grid structure
- [ ] unit tests
  - [x] elements
  - [ ] backend
  - [ ] app
- [ ] customizable theme
- [ ] more components
- [ ] easily accessible components
- [ ] key help menu
- [ ] validators
