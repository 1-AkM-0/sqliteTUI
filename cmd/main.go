package main

import (
	"log"
	"os"

	"github.com/1-AkM-0/sqliteTUI/internal/db"
	"github.com/1-AkM-0/sqliteTUI/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	client, err := db.Open(os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	p := tea.NewProgram(tui.New(client))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
