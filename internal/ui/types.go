package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// Message types for the application
type (
	TickMsg       struct{}
	InstallLogMsg struct {
		Message string
	}
	InstallCompleteMsg struct {
		Error error
	}
	InstalledPkgMsg string
)

// Model represents the application state
type Model struct {
	Choice       int
	Chosen       bool
	Ticks        int
	Quitting     bool
	Installing   bool
	InstallError error
	InstallLogs  []string

	// Package manager state
	Packages []string
	Index    int
	Width    int
	Height   int
	Spinner  spinner.Model
	Progress progress.Model
	Done     bool
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
