package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Sidebar(m Model) string {

	sideBarStyle := lipgloss.NewStyle().
		Width(20 - 2).
		Height(m.height - 2).
		Border(lipgloss.NormalBorder())

	activeStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#7aa2f7"))
	normalStyle := lipgloss.NewStyle()

	var tableList strings.Builder
	var columnsList strings.Builder

	for i, value := range m.tables {
		if i == m.activateTable {
			tableList.WriteString(activeStyle.Render("▶ " + value.Name + "\n"))
		} else {
			tableList.WriteString(normalStyle.Render(" " + value.Name + "\n"))
		}
	}

	for _, value := range m.tablesColumns {
		columnsList.WriteString(" " + value.Name + "\n")
	}

	separatorStyle := lipgloss.NewStyle().
		SetString("\n" + strings.Repeat("─", 18) + "\n" + "columns" + "\n")

	sidebar := sideBarStyle.Render(tableList.String() + separatorStyle.String() + columnsList.String())

	return sidebar
}
