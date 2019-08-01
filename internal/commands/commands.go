package commands

import (
	"github.com/WilliamLeeks/pokemon/internal/commands/help"
	"github.com/WilliamLeeks/pokemon/internal/commands/potd"
)

// Process chooses which command to execute.
func Process(cmd string) {
	switch cmd {
	case "potd":
		potd.Run()
	case "help":
		help.Run()
	default:
		help.Run()
	}
}
