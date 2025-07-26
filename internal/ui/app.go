package ui

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// App represents the main application model
type App struct {
	Model
}

// NewApp creates a new application instance
func NewApp() App {
	// Initialize spinner
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	
	// Initialize progress bar
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	
	return App{
		Model: Model{
			Choice:      0,
			Chosen:      false,
			Ticks:       0,
			Quitting:    false,
			Installing:  false,
			InstallError: nil,
			InstallLogs:  []string{},
			Packages:    []string{},
			Index:       0,
			Width:       0,
			Height:      0,
			Spinner:     s,
			Progress:    p,
			Done:        false,
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