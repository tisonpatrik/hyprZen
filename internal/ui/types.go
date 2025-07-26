package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

// Next returns the next menu choice (wraps around)
func (m MenuChoice) Next() MenuChoice {
	switch m {
	case MenuInstall:
		return MenuExit
	case MenuExit:
		return MenuInstall
	default:
		return MenuInstall
	}
}

// Previous returns the previous menu choice (wraps around)
func (m MenuChoice) Previous() MenuChoice {
	switch m {
	case MenuInstall:
		return MenuExit
	case MenuExit:
		return MenuInstall
	default:
		return MenuInstall
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
	InstallStartMsg struct{} // New message to signal installation has started
)

// ModalModel represents the main modal UI model
type ModalModel struct {
	modals        map[string]tea.Model
	selectedModal string
	width         int
	height        int
}

// NewModalModel creates a new modal UI model
func NewModalModel() *ModalModel {
	menuModal := NewMenuModal()
	installModal := NewInstallModal()
	
	modals := map[string]tea.Model{
		"menu":    menuModal,
		"install": installModal,
	}
	
	return &ModalModel{
		modals:        modals,
		selectedModal: "menu",
		width:         0,
		height:        0,
	}
}

// Init initializes the modal model
func (m *ModalModel) Init() tea.Cmd {
	return nil
}

// Update handles modal updates
func (m *ModalModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// Update all modals with window size
		for _, modal := range m.modals {
			modal.Update(msg)
		}
		return m, nil
		
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	
	// Update current modal
	currentModal := m.CurrentModal()
	updatedModal, cmd := currentModal.Update(msg)
	
	// Update the modal in our map
	m.modals[m.selectedModal] = updatedModal
	
	// Handle modal switching based on modal state
	if menuModal, ok := updatedModal.(*MenuModal); ok {
		if menuModal.shouldSwitchToInstall {
			m.selectedModal = "install"
			menuModal.shouldSwitchToInstall = false
			m.modals["menu"] = menuModal
			// Send InstallStartMsg to the install modal
			return m, tea.Batch(
				func() tea.Msg { return InstallStartMsg{} },
			)
		}
	}
	
	return m, cmd
}

// View renders the modal UI
func (m *ModalModel) View() string {
	current := m.CurrentModal()
	if current == nil {
		return "No modal available"
	}
	
	// Center the modal content
	modalContent := current.View()
	
	// Create a centered container
	centeredStyle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Height(m.height).
		Width(m.width)
	
	return centeredStyle.Render(modalContent)
}

// CurrentModal returns the currently selected modal
func (m *ModalModel) CurrentModal() tea.Model {
	if current, exists := m.modals[m.selectedModal]; exists {
		return current
	}
	return nil
}

// SwitchModal switches to a different modal
func (m *ModalModel) SwitchModal(name string) {
	if _, exists := m.modals[name]; exists {
		m.selectedModal = name
	}
}
