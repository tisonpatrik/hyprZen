package ui

import (
	"fmt"
	"hyprzen/internal/services"
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// MenuModal represents the menu selection modal
type MenuModal struct {
	Choice                MenuChoice
	shouldSwitchToInstall bool
	width                 int
	height                int
}

// NewMenuModal creates a new menu modal
func NewMenuModal() *MenuModal {
	return &MenuModal{
		Choice:                MenuInstall,
		shouldSwitchToInstall: false,
		width:                 0,
		height:                0,
	}
}

// Init initializes the menu modal
func (m *MenuModal) Init() tea.Cmd {
	return nil
}

// Update handles menu modal updates
func (m *MenuModal) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
		
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice = m.Choice.Next()
		case "k", "up":
			m.Choice = m.Choice.Previous()
		case "enter":
			switch m.Choice {
			case MenuInstall:
				m.shouldSwitchToInstall = true
			case MenuExit:
				return m, tea.Quit
			}
		}
	}
	
	return m, nil
}

// View renders the menu modal
func (m *MenuModal) View() string {
	title := InstallTitleStyle.Render("🎯 HyprZen Installation Menu")
	
	choices := fmt.Sprintf(
		"%s\n%s",
		Checkbox("Install HyprZen", m.Choice == MenuInstall),
		Checkbox("Exit", m.Choice == MenuExit),
	)
	
	controls := SubtleStyle.Render("j/k, up/down: select") + DotStyle +
		SubtleStyle.Render("enter: choose") + DotStyle +
		SubtleStyle.Render("q, esc: quit")
	
	content := fmt.Sprintf("%s\n\n%s", choices, controls)
	
	return MenuWindowStyle.Render(title + "\n" + content)
}

// InstallModal represents the installation progress modal
type InstallModal struct {
	Installing   bool
	InstallError error
	InstallLogs  []string
	Steps        []services.InstallStep
	StepIndex    int
	Spinner      spinner.Model
	Progress     progress.Model
	Done         bool
	RetryCount   int
	MaxRetries   int
	Ticks        int
	width        int
	height       int
}

// NewInstallModal creates a new installation modal
func NewInstallModal() *InstallModal {
	// Initialize spinner
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))

	// Initialize progress bar
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	
	return &InstallModal{
		Installing:   false,
		InstallError: nil,
		InstallLogs:  []string{},
		Steps:        []services.InstallStep{},
		StepIndex:    0,
		Spinner:      s,
		Progress:     p,
		Done:         false,
		RetryCount:   0,
		MaxRetries:   3,
		Ticks:        0,
		width:        0,
		height:       0,
	}
}

// Init initializes the installation modal
func (m *InstallModal) Init() tea.Cmd {
	return nil
}

// Update handles installation modal updates
func (m *InstallModal) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
		
	case InstallStartMsg:
		// Start installation
		m.Installing = true
		installer := services.NewInstallerService()
		m.Steps = installer.Install()
		return m, tea.Batch(
			ExecuteStep(m.Steps[m.StepIndex]),
			m.Spinner.Tick,
		)
		
	case StepCompleteMsg:
		step := m.Steps[m.StepIndex]
		
		// Reset retry count on success
		m.RetryCount = 0
		m.InstallError = nil
		
		if m.StepIndex >= len(m.Steps)-1 {
			// Everything's been installed. We're done!
			m.Done = true
			return m, tea.Sequence(
				tea.Printf("✓ %s", step.Name),
				tea.Quit,
			)
		}

		// Update progress bar
		m.StepIndex++
		progressCmd := m.Progress.SetPercent(float64(m.StepIndex) / float64(len(m.Steps)))

		return m, tea.Batch(
			progressCmd,
			tea.Printf("✓ %s", step.Name),
			ExecuteStep(m.Steps[m.StepIndex]),
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
				return m, tea.Quit
			}
			return m, Tick()
		}
	}
	
	return m, nil
}

// View renders the installation modal
func (m *InstallModal) View() string {
	if m.Installing && !m.Done {
		return ActiveInstallationView(m)
	}
	
	return InstallationCompleteView(m)
}

// ActiveInstallationView renders the active installation process
func ActiveInstallationView(m *InstallModal) string {
	title := InstallTitleStyle.Render("🚀 HyprZen Installation in Progress")
	
	content := InstallationProgressView(m)
	
	// Add recent logs if available
	if len(m.InstallLogs) > 0 {
		content += "\n" + InstallationLogsView(m)
	}
	
	return InstallWindowStyle.Render(title + "\n" + content)
}

// InstallationProgressView renders the progress bar and current step
func InstallationProgressView(m *InstallModal) string {
	n := len(m.Steps)
	if n == 0 {
		return StatusStyle.Render("Preparing installation...")
	}
	
	// Current step info
	currentStep := m.Steps[m.StepIndex]
	stepName := StepNameStyle.Render(currentStep.Name)
	
	// Progress bar
	progressBar := m.Progress.View()
	
	// Step counter
	stepCounter := fmt.Sprintf("Step %d of %d", m.StepIndex+1, n)
	
	// Retry info if applicable
	retryInfo := ""
	if m.RetryCount > 0 {
		retryInfo = fmt.Sprintf(" (Retry %d/%d)", m.RetryCount, m.MaxRetries)
	}
	
	// Spinner for current activity
	spinner := m.Spinner.View()
	
	// Combine everything
	progressContent := fmt.Sprintf(
		"%s\n%s\n%s %s%s\n%s",
		stepName,
		ProgressContainerStyle.Render(progressBar),
		spinner,
		stepCounter,
		retryInfo,
		StatusStyle.Render("Installing..."),
	)
	
	return progressContent
}

// InstallationLogsView renders recent installation logs
func InstallationLogsView(m *InstallModal) string {
	if len(m.InstallLogs) == 0 {
		return ""
	}
	
	// Show last 5 logs
	start := 0
	if len(m.InstallLogs) > 5 {
		start = len(m.InstallLogs) - 5
	}
	
	var logs []string
	for i := start; i < len(m.InstallLogs); i++ {
		logs = append(logs, "  "+m.InstallLogs[i])
	}
	
	return LogStyle.Render("Recent logs:\n" + strings.Join(logs, "\n"))
}

// InstallationCompleteView renders the completion or error state
func InstallationCompleteView(m *InstallModal) string {
	title := InstallTitleStyle.Render("🏁 Installation Complete")
	
	var content string
	
	if m.InstallError != nil {
		// Error state
		errorMsg := fmt.Sprintf("❌ Installation failed: %s", m.InstallError.Error())
		if m.RetryCount > 0 {
			errorMsg += fmt.Sprintf("\nRetry attempts: %d/%d", m.RetryCount, m.MaxRetries)
		}
		content = ErrorStyle.Render(errorMsg)
	} else if m.Done {
		// Success state
		successMsg := fmt.Sprintf("✅ Successfully completed %d installation steps!", len(m.Steps))
		content = StatusStyle.Render(successMsg)
	} else {
		// Generic completion
		content = StatusStyle.Render("Installation completed successfully!")
	}
	
	// Add countdown if applicable
	if m.Ticks > 0 {
		content += fmt.Sprintf("\n\nExiting in %s seconds...", TicksStyle.Render(fmt.Sprintf("%d", m.Ticks)))
	}
	
	return InstallWindowStyle.Render(title + "\n" + content)
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