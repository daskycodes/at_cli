package main

import (
	"fmt"
	"os"

	"github.com/daskycodes/at_cli/cli"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(cli.InitialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
