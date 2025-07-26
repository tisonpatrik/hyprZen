package ui

import "github.com/charmbracelet/lipgloss"

// Style constants
const (
	ProgressBarWidth  = 71
	ProgressFullChar  = "█"
	ProgressEmptyChar = "░"
	DotChar           = " • "
)

// Style definitions
var (
	KeywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	SubtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	TicksStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
	CheckboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	ProgressEmpty = SubtleStyle.Render(ProgressEmptyChar)
	DotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(DotChar)
	MainStyle     = lipgloss.NewStyle().MarginLeft(2)
	TitleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Bold(true)
) 