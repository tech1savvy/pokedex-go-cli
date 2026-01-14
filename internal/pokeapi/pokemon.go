package pokeapi

import (
	"encoding/json"
	"fmt"
)

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int  `json:"base_stat"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}

func (c *Client) GetPokemon(url string) (Pokemon, error) {
	body, err := c.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	var jsonBody Pokemon
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		return Pokemon{}, fmt.Errorf("failed to decode res.Body to json: %w", err)
	}

	return jsonBody, nil
}
