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
  namespace tuigo {
    class tuigo_pkg[" "] {
      NewContainer :: Func[bool, ContainerType, ...Element] -> Element
      NewButton :: Func[string, ButtonType, Msg] -> Element
      NewSelector :: Func[List~string~] -> Element
      NewInput :: Func[string, string, string, InputType] -> Element
      NewRadio :: Func[string] -> Element
    }
  }
  
  namespace app {
    class Backend["Backend"]{
      States :: List~AppState~
      Constructors :: Map[AppState]~Func[Element] -> Element~
      Finalizer :: Func[Map[AppState]~Element~]
    }

    class NewApp[" "] {
      NewApp :: Func[Backend, bool] -> App
    }
  }
  note for Backend "AppState = string"

  namespace tea {
    class tea_pkg[" "] {
      NewProgram :: Func[Model, ...ProgramOption] -> *Program
    }
  }

  tuigo_pkg --o Backend
  NewApp --o tea_pkg
  Backend --o NewApp
end
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