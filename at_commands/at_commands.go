package at_commands

import (
	"fmt"
	"io"
	"log"
	"main/serial_ports"
	"time"

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

func (command AtCommand) Run(port serial_ports.SerialPort) {
	baud := 115200
	timeout := 400 * time.Millisecond
	modem, err := serial.New(serial.WithPort(port.Name), serial.WithBaud(baud))
	if err != nil {
		log.Println(err)
		return
	}
	defer modem.Close()
	var mio io.ReadWriter = modem
	a := at.New(mio, at.WithTimeout(timeout))
	err = a.Init()
	if err != nil {
		log.Println(err)
		return
	}

	info, err := a.Command(command.Command)
	fmt.Println("AT" + command.Command)
	if err != nil {
		fmt.Printf("AT%s: %s\n", command.Command, err)
	}
	for _, l := range info {
		fmt.Printf("AT%s: %s\n", command.Command, l)
	}
}
