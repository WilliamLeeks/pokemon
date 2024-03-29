package potd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/WilliamLeeks/pokemon/internal/api"
	"github.com/WilliamLeeks/pokemon/internal/file"
)

const (
	wikiURL    string = "https://bulbapedia.bulbagarden.net/wiki/%s"
	numPokemon int    = 807
)

// Run executes the command.
func Run() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed getting Working Directory: %s", err)
	}

	path := filepath.Join(wd, "pokelog.txt")

	ids, err := file.ReadLog(path)
	if err != nil {
		fmt.Printf("Failed reading log file: %s", err)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	var r int

	for {
		r = rand.Intn(807)

		if contains(r, ids) {
			continue
		}
		break
	}

	p := api.GetPokemon(r)

	URL := fmt.Sprintf(wikiURL, p.Name)

	fmt.Printf("Today’s Pokémon is: %s\nLink: %s\n", strings.Title(p.Name), URL)

	line := strconv.Itoa(r) + "," + strings.Title(p.Name) + "\n"

	err = file.WriteLog(path, line)
	if err != nil {
		log.Fatalf("Error writing to log: %s\n", err)
	}

	os.Exit(0)
}

func contains(num int, set []int) bool {
	for _, a := range set {
		if a == num {
			return true
		}
	}
	return false
}
