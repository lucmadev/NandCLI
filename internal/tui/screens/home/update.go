package home

import (
	tea "charm.land/bubbletea/v2"
)

func (m HomeScreen) Update(
	msg tea.Msg,
) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "up":

			if m.cursor > 0 {
				m.cursor--
			}

		case "down":

			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case "enter":

			switch m.items[m.cursor].Route {

			case "apt":
				//return NewAptScreen(), nil

			case "docker":
				//return NewDockerScreen(), nil

			case "stats":
				//return NewStatsScreen(), nil

			case "services":
				//return NewServicesScreen(), nil

			case "settings":
				//return NewSettingsScreen(), nil

			case "exit":
				return m, tea.Quit
			}
		}
	}

	return m, nil
}