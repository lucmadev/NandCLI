package home

import (
	"fmt"
	tea "charm.land/bubbletea/v2"
)

func (m HomeScreen) Init() tea.Cmd {
	return nil
}

func NewHomeScreen() HomeScreen {

	return HomeScreen{
		items: []MenuItem{
		{"Install / Update Packages", "apt"},
		{"Manage Docker", "docker"},
		{"System Info", "stats"},
		{"Manage Services", "services"},
		{"Settings", "settings"},
		{"Open Shell", "shell"},
		{"Exit", "exit"},
		},
	}
}

func (m HomeScreen) View() tea.View {

	s := "NandCLI\n\n"

	for i, item := range m.items {

		cursor := " "

		if i == m.cursor {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, item.Title)
	}

	return tea.NewView(s)
}