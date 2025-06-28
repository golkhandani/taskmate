package main

import (
	"os"

	"github.com/golkhandani/taskmate/exceptions"
	"github.com/golkhandani/taskmate/models"
	"github.com/golkhandani/taskmate/utils"
	"github.com/golkhandani/taskmate/views"

	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	views.AppState
	StateChan chan views.AppStateChange
}

func (m App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	prevVersion := m.AppState.Version
	newState, cmd := m.AppState.Update(msg)
	m.AppState = newState.(views.AppState)

	m.StateChan <- views.AppStateChange{
		// send only changed versions or other minimal metadata
		PreviousVersion: prevVersion,
		CurrentVersion:  m.AppState.Version,
		State:           m.AppState,
	}

	return m, cmd
}

func main() {

	file, content := utils.ReadDataFile()
	defer file.Close()
	initialTasks := models.LoadTasksFromBytes(content)

	initial := views.InitState(initialTasks)
	stateChan := make(chan views.AppStateChange, 1)
	p := tea.NewProgram(App{
		AppState:  initial,
		StateChan: stateChan,
	})

	// Start a goroutine to listen to state updates
	go stateWatcher(stateChan, file)

	_, err := p.Run()
	exceptions.HandleErr(err)
}

func stateWatcher(stateChan chan views.AppStateChange, file *os.File) {
	for stateChange := range stateChan {
		if stateChange.CurrentVersion != stateChange.PreviousVersion {
			state := stateChange.State
			utils.SaveDataFile(file, &state.List.Tasks)
		}

	}
}
