package ui

import (
	"hyprzen/internal/services"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// MenuChoice represents the available menu options
type MenuChoice int

const (
	MenuInstall MenuChoice = iota
	MenuExit
	MenuChoiceCount // This will be 2, representing the total number of menu choices
)

// String returns the string representation of the menu choice
func (m MenuChoice) String() string {
	switch m {
	case MenuInstall:
		return "Install HyprZen"
	case MenuExit:
		return "Exit"
	default:
		return "Unknown"
	}
}

// Message types for the application
type (
	TickMsg       struct{}
	InstallLogMsg struct {
		Message string
	}
	InstallCompleteMsg struct {
		Error error
	}
	StepCompleteMsg string
)

// Model represents the application state
type Model struct {
	Choice       MenuChoice
	Chosen       bool
	Ticks        int
	Quitting     bool
	Installing   bool
	InstallError error
	InstallLogs  []string

	// Installation state
	Steps   []services.InstallStep
	StepIndex int
	PkgIndex  int
	Width     int
	Height    int
	Spinner   spinner.Model
	Progress  progress.Model
	Done      bool
	
	// Retry state
	RetryCount int
	MaxRetries int
}

// Commands
func Tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return TickMsg{}
	})
}

// LogInstallation creates a command that logs installation progress
func LogInstallation(message string) tea.Cmd {
	return func() tea.Msg {
		return InstallLogMsg{Message: message}
	}
}
