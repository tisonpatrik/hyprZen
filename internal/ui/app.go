package ui

import tea "github.com/charmbracelet/bubbletea"

// App represents the main application model
type App struct {
	Model
}

// NewApp creates a new application instance
func NewApp() App {
	return App{
		Model: Model{
			Choice:     0,
			Chosen:     false,
			Ticks:      0,
			Frames:     0,
			Progress:   0.0,
			Loaded:     false,
			Quitting:   false,
			Installing: false,
		},
	}
}

// Init initializes the application
func (a App) Init() tea.Cmd {
	return nil
}

// Update handles application updates
func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, cmd := Update(msg, a.Model)
	a.Model = model
	return a, cmd
}

// View renders the application
func (a App) View() string {
	return MainView(a.Model)
} 