// Package pokeapi
package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tech1savvy/pokedex-go-cli/internal/pokecache"
)

type Client struct {
	cache *pokecache.Cache
}

// NewClient creats a new client to communicate with pokeapi
func NewClient() *Client {
	client := &Client{}
	cache := pokecache.NewCache(2 * time.Millisecond)
	client.cache = cache
	return client
}

func (c *Client) Get(url string) ([]byte, error) {
	var body []byte

	cachedBody, ok := c.cache.Get(url)
	if !ok {
		// Cache Miss
		res, err := http.Get(url)
		if err != nil {
			return []byte{}, fmt.Errorf("error making api request: %w", err)
		}
		defer res.Body.Close()

		resbody, err := io.ReadAll(res.Body)
		if err != nil {
			return []byte{}, fmt.Errorf("error parsing response body: %w", err)
		}
		if res.StatusCode > 299 {
			return []byte{}, fmt.Errorf("response failed with  status code: %d and\nbody: %s", res.StatusCode, resbody)
		}
		body = resbody
	} else {
		// Cache Hit
		body = cachedBody
	}

	// Create or Update cache
	c.cache.Add(url, body)

	return body, nil
}
