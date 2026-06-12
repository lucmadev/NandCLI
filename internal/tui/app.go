package tui

import (
	"nandcli/internal/tui/screens/home"
	tea "charm.land/bubbletea/v2"
)

type App struct {
	screen tea.Model
}

func NewApp() *App {
	return &App{
		screen: home.NewHomeScreen(),
	}
}

func (a App) Init() tea.Cmd {
	return a.screen.Init()
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	a.screen, cmd = a.screen.Update(msg)

	return a, cmd
}

func (a App) View() tea.View {
	return a.screen.View()
}