package views

import (
	"fmt"
	"slices"
	"strings"

	"github.com/golkhandani/taskmate/constants"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattn/go-runewidth"
	"github.com/ttacon/chalk"
)

func UpdateListPage(m AppState, msg tea.KeyMsg) (AppState, tea.Cmd) {
	s := &m.List // Has to be reference otherwise won't update the state!
	switch msg.String() {
	case "up", "k":
		if s.Cursor > 0 {
			s.Cursor--
		}
	case "down", "j":
		if s.Cursor < len(s.Tasks)-1 {
			s.Cursor++
		}
	case " ":
		s.Tasks[s.Cursor].IsDone = !s.Tasks[s.Cursor].IsDone
		m.Version++
	case "ctrl+d":
		s.Tasks = slices.Delete(s.Tasks, s.Cursor, s.Cursor+1)
		s.Cursor = max(0, s.Cursor-1)
		m.Version++

	// Handle navigation
	case "a":
		m.Page = constants.ADD_PAGE
	case "b":
		m.Page = constants.MENU_PAGE
	case "ctrl+c":
		return m, tea.Quit
	}
	return m, nil
}

func RenderListMenu(m AppState) string {
	s := m.List
	msg := chalk.Reset.String() + "Here is your todo list?\n\n"

	for i, task := range s.Tasks {

		style := chalk.White.NewStyle().WithBackground(chalk.ResetColor)

		checked := strings.Repeat(" ", runewidth.StringWidth(constants.CHEKMARK_SYMBOL))

		if task.IsDone {
			style = chalk.Green.NewStyle().WithBackground(chalk.ResetColor)
			checked = constants.CHEKMARK_SYMBOL
		}

		cursor := strings.Repeat(" ", runewidth.StringWidth(constants.BULLET_SYMBOL))
		if s.Cursor == i {
			cursor = constants.BULLET_SYMBOL
			style = chalk.Black.NewStyle().WithBackground(chalk.Blue).WithTextStyle(chalk.Bold)
		}
		msg += fmt.Sprintf(
			style.String()+"%2s %10d - [%s] %s\n"+chalk.Reset.String(),
			cursor, i+1, checked, task.Title,
			// fmt.Sprintf("(ID:%d)", task.ID)
		)
	}

	msg += "\nPress a to add new item; ctrl+d to delete; b to back to menu; ctrl+c to quit.\n\n" + chalk.Reset.String()
	return msg
}
