package views

import (
	"github.com/golkhandani/taskmate/constants"
	"github.com/golkhandani/taskmate/models"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type MenuState struct {
	// empty for now, but useful for future settings or highlighting
}

type AddState struct {
	Input textinput.Model
}

func InitAddState() AddState {
	ti := textinput.New()
	ti.Placeholder = "Enter task description"
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 60
	ti.Prompt = ""
	return AddState{
		Input: ti,
	}
}

type ListState struct {
	Tasks  []models.Task
	Cursor int
}

func InitListState(tasks []models.Task) ListState {
	return ListState{
		Tasks:  tasks,
		Cursor: 0,
	}
}

type AppState struct {
	Version int
	Page    string
	Menu    MenuState
	Add     AddState
	List    ListState
}

type AppStateChange struct {
	PreviousVersion int
	CurrentVersion  int
	State           AppState
}

func InitState(tasks []models.Task) AppState {
	ti := textinput.New()
	ti.Placeholder = "Enter task description"
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 60
	ti.Prompt = ""
	return AppState{
		Page: constants.MENU_PAGE,
		Add: AddState{
			Input: ti,
		},
		List: InitListState(tasks),
	}
}

func (m AppState) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

// View implements tea.Model.
func (m AppState) View() string {
	switch m.Page {
	case constants.MENU_PAGE:
		return RenderOptionMenu()
	case constants.ADD_PAGE:
		return RenderAddMenu(m)
	case constants.LIST_PAGE:
		return RenderListMenu(m)
	default:
		return "404 Page Not Found"
	}
}

func (m AppState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.Page {
		case constants.MENU_PAGE:
			return UpdateMenuPage(m, msg)
		case constants.ADD_PAGE:
			return UpdateAddPage(m, msg)
		case constants.LIST_PAGE:
			return UpdateListPage(m, msg)
		}
	}
	return m, nil
}
