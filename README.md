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
    Container :: Func[bool, ContainerType, ...Element] -> ComplexContainer
    Button :: Func[string, ButtonType, Msg] -> SimpleContainer
    Selector :: Func[List~string~] -> SimpleContainer
    Input :: Func[string, string, string, InputType] -> SimpleContainer
    Radio :: Func[string] -> SimpleContainer
  }

  class Backend["tuigo.Backend"]{
    States :: List~AppState~
    Constructors :: Map[AppState]:Func[Window]->Window
    Updaters :: Map[AppState]:Func[Window,tea.Msg]->Window,tea.Cmd
    Finalizer :: Func[Map[AppState]:Window]->Window
  }

  class NewApp["tuigo"] {
    <<package>>
    App :: Func[Backend, bool] -> app.App
  }
  note for Backend "AppState = string"
  note for Backend "Window = Collection"

  class tea["tea 'github.com/charmbracelet/bubbletea'"] {
    <<package>>
    NewProgram :: Func[Model, ...ProgramOption] -> *Program
  }

  tuigo --o Backend
  NewApp --o tea
  Backend --o NewApp
```
### containers

```mermaid
%%{
  init: {
    'theme': 'base',
    'themeVariables': {
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

  class AbstractComponent {
    <<interface>>
    Hidden() bool
    Focusable() bool
    Focused() bool
	  Disabled() bool
  }

  class Component {
    <<interface>>
    Hide() Component
    Unhide() Component
	  Enable() Component
	  Disable() Component
    Focus() Component
    FocusFromStart() Component
    FocusFromEnd() Component
    Blur() Component
    FocusNext() (Component, tea.Cmd)
    FocusPrev() (Component, tea.Cmd)
    GetContainerByID(int) Component
    GetElementByID(int) Accessor
  }
  Element <|-- Component
  AbstractComponent <|-- Component

  class Collection {
    <<interface>>
    Type() utils.ContainerType
    Components() List~Component~
    AddComponents(...Component) Collection
  }
  Component <|-- Collection

  class Wrapper {
    <<interface>>
    Element() Element
  }
  Component <|-- Wrapper

  class Container {
    -bool hidden
    -bool focusable
    -bool focused
    -Func[Container] -> string render
  }
  AbstractComponent <|.. Container

  class SimpleContainer {
    -Element element
  }

  class ComplexContainer {
    -ContainerType containerType
    -List~Component~ components
  }

  Wrapper <|.. SimpleContainer
  Container <|-- SimpleContainer
  Collection <|.. ComplexContainer
  Container <|-- ComplexContainer
```

### elements

```mermaid
%%{
  init: {
    'theme': 'base',
    'themeVariables': {
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

  class Actor {
    <<interface>>
    Callback()
  }

  class Element {
    <<interface>>
    View(bool) string
    Update(tea.Msg) (Element, tea.Cmd)
  }

  class ElementWithID {
    -int id
  }

  class ElementWithCallback {
    -tea.Msg callback
  }
  ElementWithID ..|> Accessor
  ElementWithCallback ..|> Actor

  class Button {
    -string label
    -npresses int
    -ButtonType btntype
    -tea.Msg action
    +Data() -> Button::npresses
  }
  ElementWithID <|-- Button
  ElementWithCallback <|-- Button
  Element <|.. Button

  class TextInput {
    -InputType inputtype
    -textinput.Model model    
    +Data() -> TextInput::model.Value -> string
  }
  ElementWithID <|-- TextInput
  ElementWithCallback <|-- TextInput
  Element <|.. TextInput

  class Radio {
	  -string label
	  -bool state
    +Toggle() -> Radio
    +Data() -> Radio::state
  }
  ElementWithID <|-- Radio
  ElementWithCallback <|-- Radio
  Element <|.. Radio

  class Selector {
    -bool multiselect
    -int cursor
    -List~string~ options
    -Map~string~ selected
    -Map~string~ disabled
    -int view_limit
    +Enable(string) -> Selector
    +Disable(string) -> Selector
    +Disabled(string) -> bool
    +Toggle(string) -> Selector
    +SetViewLimit(int) -> Selector
    +Next() -> Selector
    +Prev() -> Selector
    +Selected() -> List~string~
    +Cursor() -> int
    +Data() -> Selector::Selected -> List~string~
  }
  Element <|.. Selector
  ElementWithID <|-- Selector
  ElementWithCallback <|-- Selector

  class Text {
    -string text
    -TextType texttype
    +Data() -> Text::text
    +Set(string) -> Text
  }
  ElementWithID <|-- Text
  ElementWithCallback <|-- Text
  Element <|.. Text
```

## TODO

- [x] app backend
- [x] grid structure
- [x] easily accessible components
- [x] update components based on others
- [ ] unit tests
  - [x] elements
  - [x] containers
  - [ ] callbacks
  - [ ] backend
  - [ ] app
- [ ] customizable theme
- [ ] more components
- [ ] key help menu
- [ ] validators
