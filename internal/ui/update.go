package ui

import (
	"hyprzen/internal/services"
	"log"

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
			if m.Choice == MenuInstall {
				m.Choice = MenuExit
			}
		case "k", "up":
			if m.Choice == MenuExit {
				m.Choice = MenuInstall
			}
		case "enter":
			m.Chosen = true
			switch m.Choice {
			case MenuInstall:
				// Install option selected - start installation
				m.Installing = true
				installer := services.NewInstallerService()
				m.Steps = installer.Install()

				// Start the installation process
				return m, tea.Batch(
					ExecuteStep(m.Steps[m.StepIndex]),
					m.Spinner.Tick,
				)
			case MenuExit:
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
		
		// Reset retry count on success
		m.RetryCount = 0
		m.InstallError = nil
		
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
		// Installation completed with error
		m.InstallError = msg.Error
		
		// Check if we should retry
		if m.RetryCount < m.MaxRetries {
			m.RetryCount++
			// Retry the current step
			return m, tea.Batch(
				tea.Printf("⚠️  %s failed, retrying (%d/%d)...", m.Steps[m.StepIndex].Name, m.RetryCount, m.MaxRetries),
				ExecuteStep(m.Steps[m.StepIndex]),
			)
		} else {
			// Max retries exceeded, show error and exit
			m.Ticks = 3
			return m, tea.Batch(
				tea.Printf("❌ %s failed after %d retries: %s", m.Steps[m.StepIndex].Name, m.MaxRetries, msg.Error),
				Tick(),
			)
		}

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
		log.Printf("ExecuteStep: Starting step: %s", step.Name)
		// Execute the step action
		if err := step.Action(); err != nil {
			log.Printf("ExecuteStep: Step failed: %v", err)
			// Return an error message
			return InstallCompleteMsg{Error: err}
		}
		log.Printf("ExecuteStep: Step completed successfully: %s", step.Name)
		// Return success message
		return StepCompleteMsg(step.Name)
	}
}
