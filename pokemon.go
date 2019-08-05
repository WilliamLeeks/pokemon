package main

import (
	"os"

	"github.com/WilliamLeeks/pokemon/internal/commands"
)

func main() {
	var cmd string

	if len(os.Args) > 1 {
		cmd = os.Args[1]
	} else {
		cmd = "help"
	}

	commands.Process(cmd)
}
