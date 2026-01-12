package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tech1savvy/pokedex-go-cli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, params ...string) error
}

type config struct {
	client *pokeapi.Client
	next   string
	prev   string
}

func cleanInput(input string) []string {
	words := strings.Split(input, " ")
	result := []string{}
	for _, w := range words {
		if w == "" {
			continue
		}

		lower := strings.ToLower(w)
		trim := strings.TrimSpace(lower)
		result = append(result, trim)
	}
	return result
}

func commandExit(c *config, params ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, params ...string) error {
	// TODO: Dynamically generate the "usage" section by iterating over my registry of commands.
	// NOTE: How to access the registry/map of cmds inside help cmmand callback?
	msg := `
Welcome to the Pokedex!
Usage:
help: Displays a help message
exit: Exit the Pokedex
		`
	fmt.Println(msg)
	return nil
}

func commandMap(c *config, params ...string) error {
	url := "https://pokeapi.co/api/v2/location-area" // GET first 20
	if c.next != "" {
		url = c.next // Get next 20
	}

	c.prev = url
	locationAreas, err := c.client.GetLocationAreas(url)
	c.next = locationAreas.Next
	c.prev = locationAreas.Previous
	if err != nil {
		return err
	}

	for i := range len(locationAreas.Results) {
		fmt.Println(locationAreas.Results[i].Name)
	}

	return nil
}

func commandMapB(c *config, params ...string) error {
	if c.prev == "" {
		fmt.Println("You're on the first page!")
		return nil
	}
	url := c.prev

	locationAreas, err := c.client.GetLocationAreas(url)
	c.next = locationAreas.Next
	c.prev = locationAreas.Previous
	if err != nil {
		return err
	}

	for i := range len(locationAreas.Results) {
		fmt.Println(locationAreas.Results[i].Name)
	}

	return nil
}

func commandExplore(c *config, params ...string) error {
	// FIX: index out of range if len(params) == 0
	location := params[0]
	if location == "" {
		fmt.Println("Please proved a location or area name")
		return nil
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
	// TODO: use `pokeapiLocationArea` identifier instead
	locationArea, err := c.client.GetLocationArea(url)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Printf("Found Pokemon:\n")

	for _, pokemonEncounter := range locationArea.PokemonEncounters {
		fmt.Printf("- %v\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
