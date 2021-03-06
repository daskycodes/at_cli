package cli

import (
	"fmt"

	"github.com/daskycodes/at_cli/at_commands"
	"github.com/daskycodes/at_cli/serial_ports"
)

func (m Model) View() string {
	header := "AT Command CLI\n\n"
	information := fmt.Sprintf("Selected Serial Port: %s\n\n", m.SelectedPort.Name)
	mainView := m.MainView.Render(m)
	inputs := "\nPress ctrl+k for custom input\n"
	if m.SelectedPort.Name == "No Port Selected" {
		inputs = "\n\n"
	}
	footer := "\nPress ctrl+c to quit.\n"

	return header + information + mainView + inputs + footer
}

type MainView struct {
	Name       string
	Action     func(Model) Model
	ItemLength int
	Render     func(Model) string
}

var SerialPortView = MainView{
	Name: "SerialPortView",
	Render: func(m Model) string {
		var portListString string
		for i, port := range serial_ports.GetSerialPorts() {
			cursor := " "
			if m.Cursor == i {
				cursor = "👉"
			}
			portListString += fmt.Sprintf("%s %s\n", cursor, port.Name)
		}
		return portListString
	},
	Action: func(m Model) Model {
		m.SelectedPort = serial_ports.GetSerialPorts()[m.Cursor]
		m.MainView = AtCommandView
		m.Cursor = 0
		return m
	},
	ItemLength: len(serial_ports.GetSerialPorts()),
}

var CustomInputView = MainView{
	Name: "CustomInputView",
	Render: func(m Model) string {
		return fmt.Sprintf(
			"Custom AT command: %s\n", m.textInput.View(),
		)
	}, Action: func(m Model) Model {
		command := at_commands.AtCommand{Command: m.textInput.Value()}
		info := command.Run(m.SelectedPort)
		m.AtCommandResult = info
		m.MainView = AtCommandResultView
		m.textInput.Blur()
		return m
	},
}

var AtCommandView = MainView{
	Name: "AtCommandView",
	Render: func(m Model) string {
		var commandListString string
		for i, command := range at_commands.GetModemInfoCommands() {
			cursor := " "
			if m.Cursor == i {
				cursor = "👉"
			}
			commandListString += fmt.Sprintf("%s %s - %s\n", cursor, command.Command, command.Description)
		}
		return commandListString
	},
	Action: func(m Model) Model {
		command := at_commands.GetModemInfoCommands()[m.Cursor]
		info := command.Run(m.SelectedPort)
		m.AtCommandResult = info
		m.MainView = AtCommandResultView
		return m
	},
	ItemLength: len(at_commands.GetModemInfoCommands()),
}

var AtCommandResultView = MainView{
	Name: "AtCommandResultView",
	Render: func(m Model) string {
		var resultString string
		for _, l := range m.AtCommandResult {
			resultString += fmt.Sprintf("%s\n", l)
		}
		return resultString
	},
	Action: func(m Model) Model {
		return m
	},
	ItemLength: 0,
}
