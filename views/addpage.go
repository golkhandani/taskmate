package views

import (
	"fmt"

	"github.com/golkhandani/taskmate/constants"
	"github.com/golkhandani/taskmate/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ttacon/chalk"
)

func UpdateAddPage(m AppState, msg tea.KeyMsg) (AppState, tea.Cmd) {
	s := &m.Add
	switch msg.String() {
	case tea.KeyEscape.String():
		s.Input.Reset()
		m.Page = constants.MENU_PAGE
		return m, nil
	case tea.KeyEnter.String():
		title := s.Input.Value()
		if title != "" {
			newID := int64(len(m.List.Tasks))
			if newID == 0 {
				newID = 1
			} else {
				newID = m.List.Tasks[len(m.List.Tasks)-1].ID + 1
			}
			m.List.Tasks = append(m.List.Tasks, models.Task{
				ID:     newID,
				Title:  title,
				IsDone: false,
			})
		}
		s.Input.Reset()
		m.Page = constants.LIST_PAGE
		m.Version++
		return m, nil
	case "ctrl+c":
		return m, tea.Quit
	}

	var cmd tea.Cmd
	s.Input, cmd = s.Input.Update(msg)
	return m, cmd
}

func RenderAddMenu(m AppState) string {

	return fmt.Sprintf(
		chalk.Reset.String()+
			"Add a new task:\n\n%s\n\nPress Enter to save, Esc to cancel.\n\n",
		m.Add.Input.View())

}
