package at_commands

import (
	"fmt"
	"io"
	"time"

	"git.coco.study/dkhaapam/at_cli/serial_ports"

	"github.com/warthog618/modem/at"
	"github.com/warthog618/modem/serial"
)

type AtCommand struct {
	Command     string
	Description string
}

func GetModemInfoCommands() []AtCommand {
	return []AtCommand{
		{"I", "Display product identification information."},
	}
}

func (command AtCommand) Run(port serial_ports.SerialPort) []string {
	baud := 115200
	timeout := 400 * time.Millisecond
	modem, err := serial.New(serial.WithPort(port.Name), serial.WithBaud(baud))
	if err != nil {
		return []string{fmt.Sprintf("Error: %s", err)}
	}
	defer modem.Close()
	var mio io.ReadWriter = modem
	a := at.New(mio, at.WithTimeout(timeout))
	err = a.Init()
	if err != nil {
		return []string{fmt.Sprintf("Error: %s", err)}
	}

	info, err := a.Command(command.Command)
	if err != nil {
		return []string{fmt.Sprintf("Error: %e", err)}
	}
	return info
}
