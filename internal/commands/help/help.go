package help

import (
	"fmt"
	"os"
)

// Run executes the command.
func Run() {
	helpText := `Welcome to Pokémon CLI v1.0

Available Commands:
potd - Random Pokémon of the Day.
help - Help information.
`

	fmt.Println(helpText)
	os.Exit(0)
}
