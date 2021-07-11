package serial_ports

import (
	"log"

	"go.bug.st/serial"
)

type SerialPort struct {
	Name string
}

func GetSerialPorts() []SerialPort {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	var serialPorts []SerialPort
	for _, port := range ports {
		serialPorts = append(serialPorts, SerialPort{Name: port})
	}
	return serialPorts
}
