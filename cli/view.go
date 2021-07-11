package cli

import (
	"fmt"

	"git.coco.study/dkhaapam/at_cli/at_commands"
	"git.coco.study/dkhaapam/at_cli/serial_ports"
)

func (m Model) View() string {
	header := "AT Command CLI\n\n"
	information := fmt.Sprintf("Selected Serial Port: %s\n\n", m.SelectedPort.Name)
	mainView := m.MainView.Render(m)
	footer := "\nPress q to quit.\n"

	return header + information + mainView + footer
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

var AtCommandView = MainView{
	Name: "AtCommndView",
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
		command.Run(m.SelectedPort)
		return m
	},
	ItemLength: len(at_commands.GetModemInfoCommands()),
}
