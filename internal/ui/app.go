package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// App represents the main application model
type App struct {
	*ModalModel
}

// NewApp creates a new application instance
func NewApp() App {
	return App{
		ModalModel: NewModalModel(),
	}
}

// Init initializes the application
func (a App) Init() tea.Cmd {
	return nil
}

// Update handles application updates
func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, cmd := a.ModalModel.Update(msg)
	if modalModel, ok := model.(*ModalModel); ok {
		a.ModalModel = modalModel
	}
	return a, cmd
}

// View renders the application
func (a App) View() string {
	return a.ModalModel.View()
}
