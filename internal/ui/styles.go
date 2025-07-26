package ui

import "github.com/charmbracelet/lipgloss"

// Style constants
const (
	DotChar = " • "
)

// Style definitions
var (
	KeywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	SubtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	TicksStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
	CheckboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	DotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(DotChar)
	MainStyle     = lipgloss.NewStyle().MarginLeft(2)
	TitleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Bold(true)
	
	// Installation interface styles
	InstallTitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("99")).
		Bold(true).
		Align(lipgloss.Center).
		MarginBottom(1)
	
	InstallWindowStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(2, 3).
		Margin(2, 4).
		Background(lipgloss.Color("235")).
		Foreground(lipgloss.Color("252"))
	
	MenuWindowStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("99")).
		Padding(2, 3).
		Margin(2, 4).
		Background(lipgloss.Color("235")).
		Foreground(lipgloss.Color("252"))
	
	StepNameStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("211")).
		Bold(true)
	
	ProgressContainerStyle = lipgloss.NewStyle().
		MarginTop(1).
		MarginBottom(1)
	
	LogStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Italic(true).
		MarginTop(1)
	
	StatusStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("79")).
		Bold(true).
		Align(lipgloss.Center)
	
	ErrorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("203")).
		Bold(true)
)
