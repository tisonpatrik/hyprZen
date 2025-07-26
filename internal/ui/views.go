package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// MainView renders the main application view
func MainView(m Model) string {
	var s string
	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	if !m.Chosen {
		s = ChoicesView(m)
	} else {
		s = ChosenView(m)
	}
	return MainStyle.Render("\n" + s + "\n\n")
}

// ChoicesView renders the menu selection view
func ChoicesView(m Model) string {
	c := m.Choice

	tpl := TitleStyle.Render("HyprZen Installation Menu") + "\n\n"
	tpl += "%s\n\n"
	tpl += SubtleStyle.Render("j/k, up/down: select") + DotStyle +
		SubtleStyle.Render("enter: choose") + DotStyle +
		SubtleStyle.Render("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s",
		Checkbox("Install HyprZen", c == 0),
		Checkbox("Exit", c == 1),
	)

	return fmt.Sprintf(tpl, choices)
}

// ChosenView renders the installation status view
func ChosenView(m Model) string {
	if m.Installing && !m.Done {
		return PackageManagerView(m)
	}

	var msg string
	if m.InstallError != nil {
		msg = fmt.Sprintf("Installation failed: %s", m.InstallError.Error())
	} else if m.Done {
		msg = fmt.Sprintf("Done! Completed %d installation steps.\n", len(m.Steps))
	} else {
		msg = "Installation completed successfully!"
	}

	if m.Ticks > 0 {
		msg += fmt.Sprintf("\n\nExiting in %s seconds...", TicksStyle.Render(fmt.Sprintf("%d", m.Ticks)))
	}

	return msg
}

// PackageManagerView renders the installation progress interface
func PackageManagerView(m Model) string {
	n := len(m.Steps)
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	stepCount := fmt.Sprintf(" %*d/%*d", w, m.StepIndex, w, n)

	spin := m.Spinner.View() + " "
	prog := m.Progress.View()
	cellsAvail := max(0, m.Width-lipgloss.Width(spin+prog+stepCount))

	stepName := KeywordStyle.Render(m.Steps[m.StepIndex].Name)
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render("Executing " + stepName)

	cellsRemaining := max(0, m.Width-lipgloss.Width(spin+info+prog+stepCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + prog + stepCount
}

// Checkbox renders a checkbox option
func Checkbox(label string, checked bool) string {
	if checked {
		return CheckboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
