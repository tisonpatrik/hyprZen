package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Message types for the application
type (
	TickMsg  struct{}
	FrameMsg struct{}
)

// Model represents the application state
type Model struct {
	Choice      int
	Chosen      bool
	Ticks       int
	Frames      int
	Progress    float64
	Loaded      bool
	Quitting    bool
	Installing  bool
}

// Commands
func Tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return TickMsg{}
	})
}

func Frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return FrameMsg{}
	})
} 