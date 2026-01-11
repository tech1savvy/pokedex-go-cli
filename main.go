package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tech1savvy/pokedex-go-cli/internal/pokeapi"
)

func main() {
	cfg := &config{}

	cmds := map[string]cliCommand{
		"exit": {
			"exit",
			"Exit the Pokedex",
			commandExit,
		},
		"help": {
			"help",
			"Displays a help message",
			commandHelp,
		},
		"map": {
			"map",
			"Display Next Locations",
			commandMap,
		},
		"mapb": {
			"mapb",
			"Display Previous Locations",
			commandMapB,
		},
	}
	client := pokeapi.NewClient()
	cfg.client = client

	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Pokedex > ")
		scanner.Scan()
		rawInput := scanner.Text()
		input := cleanInput(rawInput)

		if len(input) == 0 {
			continue
		}

		if _, ok := cmds[input[0]]; !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmds[input[0]].callback(cfg); err != nil {
			fmt.Println(err)
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Input scanner failed: %s", err)
		}
	}
}
