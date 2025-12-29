// Package pokeapi
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tech1savvy/pokedex-go-cli/internal/pokecache"
)

type PokeAPIResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Client struct {
	cache *pokecache.Cache
}

// NewClient returns pointer to new Client struct initialized
func NewClient() *Client {
	client := &Client{}
	cache := pokecache.NewCache(2 * time.Millisecond)
	client.cache = cache
	return client
}

func (c *Client) GetLocationAreas(url string) (next string, prev string, err error) {
	// TODO: Check cache first
	_, ok := c.cache.Get(url)
	if ok {
		// TODO: Get values form byte cache
		return "", "", nil
	}

	res, err := http.Get(url)
	if err != nil {
		return "", "", fmt.Errorf("error making api request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", fmt.Errorf("error parsing response body: %w", err)
	}
	if res.StatusCode > 299 {
		return "", "", fmt.Errorf("response failed with  status code: %d and\nbody: %s", res.StatusCode, body)
	}

	data := PokeAPIResponse{}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", "", fmt.Errorf("failed to decode res.Body to json: %w", err)
	}

	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	if data.Next != nil {
		next = *data.Next
	}
	if data.Previous != nil {
		prev = *data.Previous
	}

	return next, prev, nil
}
