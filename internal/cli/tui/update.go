package tui

import tea "github.com/charmbracelet/bubbletea"

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < m.optionCount()-1 {
				m.cursor++
			}
		case "enter":
			switch m.current {
			case menuMain:
				switch m.cursor {
				case 0:
					m.current = menuBitable
				case 1:
					m.current = menuTable
				case 2:
					m.current = menuField
				case 3:
					m.current = menuRecord
				case 4:
					m.current = menuConfig
				case 5:
					return m, tea.Quit
				}
				m.cursor = 0
			default:
				m.current = menuMain
				m.cursor = 0
			}
		case "esc", "left", "h":
			m.current = menuMain
			m.cursor = 0
		}
	}
	return m, nil
}
