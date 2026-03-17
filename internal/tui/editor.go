package tui

import "github.com/charmbracelet/lipgloss"

func Editor(m Model) string {
	editorStyle := lipgloss.NewStyle().
		Width(m.width - 20 - 2).
		Height((m.height / 2) - 2).
		Border(lipgloss.NormalBorder())

	return editorStyle.Render("query")

}
