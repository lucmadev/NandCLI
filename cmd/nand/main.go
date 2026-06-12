package main

import (
	"log"
	"nandcli/internal/config"
	"nandcli/internal/tui"

	tea "charm.land/bubbletea/v2"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	config.Current = cfg

	p := tea.NewProgram(
		tui.NewApp(),
	)

	p.Run()
}
