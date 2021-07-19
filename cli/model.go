package cli

import (
	"git.coco.study/dkhaapam/at_cli/serial_ports"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	Cursor          int
	SelectedPort    serial_ports.SerialPort
	MainView        MainView
	AtCommandResult []string
	textInput       textinput.Model
}

func InitialModel() Model {
	ti := textinput.NewModel()
	ti.Placeholder = "+CPIN?"
	ti.CharLimit = 156
	ti.Width = 20
	ti.Prompt = "> AT"
	return Model{
		textInput:    ti,
		MainView:     SerialPortView,
		SelectedPort: serial_ports.SerialPort{Name: "No Port Selected"},
	}
}
func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
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

		case "esc":
			m.MainView = AtCommandView
			m.textInput.Blur()

		case "ctrl+k":
			m.MainView = CustomInputView
			m.textInput.Focus()
		}

	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}
