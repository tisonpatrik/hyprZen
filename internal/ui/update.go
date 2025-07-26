package ui

import (
	"hyprzen/internal/services"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
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
				// Install option selected - start installation
				m.Installing = true
				installer := services.NewInstallerService()
				m.Steps = installer.Install()

				// Start the installation process
				return m, tea.Batch(
					ExecuteStep(m.Steps[m.StepIndex]),
					m.Spinner.Tick,
				)
			} else {
				// Exit option selected
				m.Quitting = true
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

// UpdateChosen handles updates for the installation state
func UpdateChosen(msg tea.Msg, m Model) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height

	case StepCompleteMsg:
		step := m.Steps[m.StepIndex]
		if m.StepIndex >= len(m.Steps)-1 {
			// Everything's been installed. We're done!
			m.Done = true
			return m, tea.Sequence(
				tea.Printf("✓ %s", step.Name), // print the last success message
				tea.Quit,                      // exit the program
			)
		}

		// Update progress bar
		m.StepIndex++
		progressCmd := m.Progress.SetPercent(float64(m.StepIndex) / float64(len(m.Steps)))

		return m, tea.Batch(
			progressCmd,
			tea.Printf("✓ %s", step.Name),           // print success message above our program
			ExecuteStep(m.Steps[m.StepIndex]),       // execute the next step
		)

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd

	case progress.FrameMsg:
		newModel, cmd := m.Progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.Progress = newModel
		}
		return m, cmd

	case InstallLogMsg:
		// Add log message to the installation logs
		m.InstallLogs = append(m.InstallLogs, msg.Message)
		return m, nil

	case InstallCompleteMsg:
		// Installation completed
		m.InstallError = msg.Error
		m.Ticks = 3
		return m, Tick()

	case TickMsg:
		if m.Ticks > 0 {
			m.Ticks--
			if m.Ticks == 0 {
				m.Quitting = true
				return m, tea.Quit
			}
			return m, Tick()
		}
	}

	return m, nil
}

// ExecuteStep creates a command that executes an installation step
func ExecuteStep(step services.InstallStep) tea.Cmd {
	return func() tea.Msg {
		// Execute the step action
		if err := step.Action(); err != nil {
			// Return an error message instead of ignoring it
			return InstallCompleteMsg{Error: err}
		}
		// Simulate some processing time
		time.Sleep(time.Millisecond * 1000)
		return StepCompleteMsg(step.Name)
	}
}
