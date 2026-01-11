package pokeapi

import (
	"encoding/json"
	"fmt"
)

type Pokemon struct {
	Name string `json:"name"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocationArea struct {
	Name              string              `json:"name"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

func (c *Client) GetLocationArea(url string) (LocationArea, error) {
	body, err := c.Get(url)
	if err != nil {
		return LocationArea{}, err
	}

	var jsonBody LocationArea
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		return LocationArea{}, fmt.Errorf("failed to decode res.Body to json: %w", err)
	}

	return jsonBody, nil
}
