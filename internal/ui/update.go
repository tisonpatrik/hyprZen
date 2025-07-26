package ui

import (
	"hyprzen/internal"

	tea "github.com/charmbracelet/bubbletea"
)

// Update handles all application state updates
func Update(msg tea.Msg, m Model) (Model, tea.Cmd) {
	// Handle quit keys
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}

	// Route to appropriate update function based on state
	if !m.Chosen {
		return UpdateChoices(msg, m)
	}
	return UpdateChosen(msg, m)
}

// UpdateChoices handles updates for the menu selection state
func UpdateChoices(msg tea.Msg, m Model) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice++
			if m.Choice > 1 {
				m.Choice = 1
			}
		case "k", "up":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = 0
			}
		case "enter":
			m.Chosen = true
			if m.Choice == 0 {
				// Install option selected
				m.Installing = true
				return m, Frame()
			} else {
				// Exit option selected
				m.Quitting = true
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

// UpdateChosen handles updates for the installation progress state
func UpdateChosen(msg tea.Msg, m Model) (Model, tea.Cmd) {
	switch msg.(type) {
	case FrameMsg:
		if !m.Loaded {
			m.Frames++
			m.Progress = float64(m.Frames) / float64(100)
			if m.Progress >= 1 {
				m.Progress = 1
				m.Loaded = true
				// Run the actual installation
				go func() {
					internal.PreInstallSetup()
					internal.InstallSystem()
					internal.InstallAps()
					internal.AddConfigs()
				}()
				m.Ticks = 3
				return m, Tick()
			}
			return m, Frame()
		}

	case TickMsg:
		if m.Loaded {
			if m.Ticks == 0 {
				m.Quitting = true
				return m, tea.Quit
			}
			m.Ticks--
			return m, Tick()
		}
	}

	return m, nil
} 