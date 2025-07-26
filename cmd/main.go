package main

import (
	"fmt"
	"hyprzen/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	app := ui.NewApp()
	p := tea.NewProgram(app)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}
