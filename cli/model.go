package cli

import (
	"main/serial_ports"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Cursor       int
	SelectedPort serial_ports.SerialPort
	MainView     MainView
}

var InitialModel = Model{
	MainView: SerialPortView,
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down", "j":
			if m.Cursor < m.MainView.ItemLength-1 {
				m.Cursor++
			}

		case "enter", " ":
			m = m.MainView.Action(m)
		}
	}

	return m, nil
}
