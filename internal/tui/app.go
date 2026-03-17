package tui

import (
	"github.com/1-AkM-0/sqliteTUI/internal/db"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Panel int
type Mode int

const (
	SideBar Panel = iota
	EditorPanel
	ResultsPanel
)

const (
	Normal Mode = iota
	Insert
	Visual
)

type Model struct {
	client        *db.Client
	tables        []db.Table
	tablesColumns []db.Column
	activateTable int
	query         string
	result        *db.Result
	focused       Panel
	mode          Mode
	err           error
	width         int
	height        int
}

func New(client *db.Client) Model {
	return Model{
		client:  client,
		focused: SideBar,
		mode:    Normal,
	}
}

func (m Model) Init() tea.Cmd {
	return loadTables(m.client)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String():
			return m, tea.Quit
		case "j", tea.KeyDown.String():
			if m.activateTable < len(m.tables)-1 {
				m.activateTable++
			}
		case "k", tea.KeyUp.String():
			if m.activateTable > 0 {
				m.activateTable--
			}
		case tea.KeyEnter.String():
			return m, loadColumns(m.client, m.tables[m.activateTable].Name)

		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tablesLoadedMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}
		m.tables = msg.tables
		return m, nil

	case columnsLoadedMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}
		m.tablesColumns = msg.columns
		return m, nil

	}

	return m, nil
}

func (m Model) View() string {

	sidebar := Sidebar(m)
	editor := Editor(m)
	results := Results(m)
	vertical := lipgloss.JoinVertical(lipgloss.Top, editor, results)
	horizontal := lipgloss.JoinHorizontal(lipgloss.Left, sidebar, vertical)

	return horizontal
}
