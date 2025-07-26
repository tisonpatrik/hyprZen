package ui

import (
	"fmt"
	"math"
	"strings"
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

// ChosenView renders the installation progress view
func ChosenView(m Model) string {
	msg := "Installing HyprZen...\n\n"
	msg += "Setting up your system with HyprZen configuration..."

	label := "Preparing installation..."
	if m.Loaded {
		label = fmt.Sprintf("Installation completed! Exiting in %s seconds...", TicksStyle.Render(fmt.Sprintf("%d", m.Ticks)))
	}

	return msg + "\n\n" + label + "\n" + ProgressBar(m.Progress) + "%"
}

// Checkbox renders a checkbox option
func Checkbox(label string, checked bool) string {
	if checked {
		return CheckboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}

// ProgressBar renders a progress bar
func ProgressBar(percent float64) string {
	w := float64(ProgressBarWidth)

	fullSize := int(percent * w)
	var fullCells string
	for i := 0; i < fullSize; i++ {
		fullCells += KeywordStyle.Render(ProgressFullChar)
	}

	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(ProgressEmpty, emptySize)

	return fmt.Sprintf("%s%s %3.0f", fullCells, emptyCells, math.Round(percent*100))
} 