# Terminal Task Manager

This is a small project I built while learning Go. I'm using simple hands-on projects like this to explore Go's ecosystem, improve my skills, and document my learning. I plan to share more of these projects on GitHub as I go.

This particular project is a terminal-based to-do list built using [Bubble Tea](https://github.com/charmbracelet/bubbletea), a powerful TUI framework for Go.

## Features

- Navigate with arrow keys or `j/k`
- Add tasks with a clean input form
- Mark tasks as done (`space`)
- Delete tasks (`ctrl+d`)
- Return to menu (`b`)
- Persistent state (optional file storage)
- Unicode-aware layout with color support

## Tech Stack

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) – TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) – UI components like `textinput`
- [Chalk](https://github.com/ttacon/chalk) – Terminal colors
- [RuneWidth](https://github.com/mattn/go-runewidth) – Unicode width accuracy

## Usage

```bash
go run main.go
```

![Task Manager Demo](assets/demo.gif)
