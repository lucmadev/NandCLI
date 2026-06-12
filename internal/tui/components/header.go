package components

type Menu struct {
	Cursor int
	Items []string
}

func (m Menu) View() string {

	var s string

	for i, item := range m.Items {

		cursor := " "

		if i == m.Cursor {
			cursor = ">"
		}

		s += cursor + " " + item + "\n"
	}

	return s
}