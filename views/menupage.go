package views

import (
	"fmt"
	"strings"

	"github.com/golkhandani/taskmate/constants"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattn/go-runewidth"
	"github.com/ttacon/chalk"
)

func RenderOptionMenu() string {
	return fmt.Sprint(
		banner() + chalk.Reset.String() +
			"Select the option:\n\n" +
			"1. View Tasks\n" +
			"2. Add Task\n" +
			"\nPress 1 or 2, or ctrl+c to quit.\n\n",
	)
}

func UpdateMenuPage(m AppState, msg tea.KeyMsg) (AppState, tea.Cmd) {

	switch msg.String() {
	case "1":
		m.Page = constants.LIST_PAGE
	case "2":
		m.Page = constants.ADD_PAGE
	case "ctrl+c":
		return m, tea.Quit
	default:

	}
	return m, nil
}

func banner() string {
	msg := ""
	lines := []string{
		"Welcome to TaskMate CLI",
		"Your simple terminal task tracker!",
	}

	// Calculate max visual width
	maxWidth := 0
	for _, line := range lines {
		w := runewidth.StringWidth(line)
		if w > maxWidth {
			maxWidth = w
		}
	}

	offset := 2
	margin := 40
	maxWidth += margin
	// Build top border
	top := "┌" + strings.Repeat("─", maxWidth+offset) + "┐"
	msg += fmt.Sprintln(chalk.Green.Color(top))

	// Build content lines with padding
	for _, line := range lines {
		visibleWidth := runewidth.StringWidth(line)
		padding := maxWidth - visibleWidth - offset
		leftPadding := padding / 2
		rightPadding := padding / 2
		if padding != rightPadding+leftPadding {
			leftPadding += padding - rightPadding*2
		}
		msg += fmt.Sprintln(
			chalk.Green.Color("│ "),
			chalk.Cyan.Color(
				strings.Repeat(" ", leftPadding)+
					line+
					strings.Repeat(" ", rightPadding),
			),
			chalk.Green.Color(" │"),
		)
	}

	// Build bottom border
	bottom := "└" + strings.Repeat("─", maxWidth+offset) + "┘"
	msg += fmt.Sprintln(chalk.Green.Color(bottom))
	msg += fmt.Sprintln()

	return msg
}
