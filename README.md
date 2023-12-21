# `tuigo` 

## a terminal UI framework written in Go using the `bubbletea` library.

### building an application

the backbone of any application is the `tuigo.Backend` object, which contains all the states of the app, and describes the workflow of changing the states and finalizing the app. the backend object contains the constructors for each "state" of the application (returning an `Element`), and can optionally depend on the previous state (passed as an argument). most of the functions for making new elements are exposed through `tuigo.%Element%`, which return an `Element`. `Container` elements can be nested, while each display element (buttons, selectors, texts, etc) is contained within a parent container. the backend also contains a finalizer function, which is called when the app finishes successfully.

the `tuigo.NewApp` function takes a `Backend` and a boolean value, which determines whether the app should be run in debug mode. the backend is then used to create an app, which is the passed to the `tea.NewProgram`.

see [`examples/`](examples/) for usage examples of `tuigo` for building applications. the scheme below shows roughly the structure of the API. 

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
    Container :: Func[bool, ContainerType, ...Element] -> Element
    Button :: Func[string, ButtonType, Msg] -> Element
    Selector :: Func[List~string~] -> Element
    Input :: Func[string, string, string, InputType] -> Element
    Radio :: Func[string] -> Element
  }

  class Backend["tuigo.Backend"]{
    States :: List~AppState~
    Constructors :: Map[AppState]~Func[Collection]->Collection~
    Finalizer :: Func[Map[AppState]~Collection~]
  }

  class NewApp["tuigo"] {
    <<package>>
    App :: Func[Backend, bool] -> app.App
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
