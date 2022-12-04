package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func Run() (tea.Model, error) {
	return tea.NewProgram(newModel(), tea.WithAltScreen()).Run()
}
