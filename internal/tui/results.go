package tui

import "github.com/charmbracelet/lipgloss"

func Results(m Model) string {
	resultStyle := lipgloss.NewStyle().
		Width(m.width - 20 - 2).
		Height(m.height/2 - 2).
		Border(lipgloss.NormalBorder())

	return resultStyle.Render("result")
}
