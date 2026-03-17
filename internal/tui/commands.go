package tui

import (
	"github.com/1-AkM-0/sqliteTUI/internal/db"
	tea "github.com/charmbracelet/bubbletea"
)

type tablesLoadedMsg struct {
	tables []db.Table
	err    error
}

type columnsLoadedMsg struct {
	columns []db.Column
	err     error
}

func loadTables(client *db.Client) tea.Cmd {
	return func() tea.Msg {
		tables, err := client.Tables()
		return tablesLoadedMsg{tables: tables, err: err}
	}
}

func loadColumns(client *db.Client, tableName string) tea.Cmd {
	return func() tea.Msg {
		columns, err := client.Columns(tableName)
		return columnsLoadedMsg{columns: columns, err: err}
	}
}
