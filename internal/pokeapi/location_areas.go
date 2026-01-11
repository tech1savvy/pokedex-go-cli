package pokeapi

import (
	"encoding/json"
	"fmt"
)

type LocationAreas struct {
	Next     string `json:"next"`
	Previous string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(url string) (LocationAreas, error) {
	body, err := c.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}

	var jsonBody LocationAreas
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		return LocationAreas{}, fmt.Errorf("failed to decode res.Body to json: %w", err)
	}

	return jsonBody, nil
}
