# `tuigo` 

a terminal UI framework written in Go using the `bubbletea` library.

![order](examples/order.gif)

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
    Finalizer :: Func[Map[AppState]~Collection~]->Collection
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
### elements

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
  class Accessor {
    <<interface>>
    ID() int
    Data() interface<>
  }
  class Element {
    <<interface>>
    View(bool) string
    Update(tea.Msg) (Element, tea.Cmd)
  }
  class Collection {
    <<interface>>
    Elements() []Element
    AddElements(...Element) Collection
    Focusable() bool
    Focused() bool
    Focus() Collection
    FocusFromStart() Collection
    FocusFromEnd() Collection
    Blur() Collection
    FocusNext() (Collection, tea.Cmd)
    FocusPrev() (Collection, tea.Cmd)
    GetElementByID(int) Accessor
  }

  class Button {
    -string label
    -npresses int
    -ButtonType btntype
    -tea.Msg action
    +Data() -> Button::npresses
  }

  class TextInput {
    -InputType inputtype
    -textinput.Model model    
    +Data() -> TextInput::model.Value -> string
  }

  class Radio {
	  -string label
	  -bool state
    +Toggle() -> Radio
    +Data() -> Radio::state
  }

  class Selector {
    -bool multiselect
    -int cursor
    -List~string~ options
    -Map~string~ selected
    -Map~string~ disabled
    +Disable(string) -> Selector
    +Enable(string) -> Selector
    +Toggle(string) -> Selector
    +Next() -> Selector
    +Prev() -> Selector
    +Selected() -> List~string~
    +Cursor() -> int
    +Data() -> Selector::Selected -> List~string~
  }

  class Text {
    -string text
    -TextType texttype
    +Data() -> Text::text
  }

  class Container {
    -bool focusable
    -bool focused
    -ContainerType conttype
    -List~Element~ elements
    -Func[Container] -> string render
  }

  Element <|-- Button
  Accessor <|-- Button
  Element <|-- TextInput
  Accessor <|-- TextInput
  Element <|-- Radio
  Accessor <|-- Radio
  Element <|-- Selector
  Accessor <|-- Selector
  Element <|-- Text
  Accessor <|-- Text
  Element <|-- Container
  Collection <|-- Container
```

## TODO

- [x] app backend
- [x] grid structure
- [x] easily accessible components
- [ ] update components based on others
- [ ] unit tests
  - [x] elements
  - [ ] backend
  - [ ] app
- [ ] customizable theme
- [ ] more components
- [ ] key help menu
- [ ] validators
