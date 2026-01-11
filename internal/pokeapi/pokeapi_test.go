package pokeapi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetLocationAreas(t *testing.T) {
	tests := map[string]struct {
		url  string
		want LocationAreas
	}{
		"first 20 locations": {
			"https://pokeapi.co/api/v2/location-area",
			LocationAreas{
				Next: "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
				Results: []struct {
					Name string `json:"name"`
				}{
					{"canalave-city-area"},
					{"eterna-city-area"},
					{"pastoria-city-area"},
					{"sunyshore-city-area"},
					{"sinnoh-pokemon-league-area"},
					{"oreburgh-mine-1f"},
					{"oreburgh-mine-b1f"},
					{"valley-windworks-area"},
					{"eterna-forest-area"},
					{"fuego-ironworks-area"},
					{"mt-coronet-1f-route-207"},
					{"mt-coronet-2f"},
					{"mt-coronet-3f"},
					{"mt-coronet-exterior-snowfall"},
					{"mt-coronet-exterior-blizzard"},
					{"mt-coronet-4f"},
					{"mt-coronet-4f-small-room"},
					{"mt-coronet-5f"},
					{"mt-coronet-6f"},
					{"mt-coronet-1f-from-exterior"},
				},
			},
		},
		"first 20-40 locations": {
			"https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
			LocationAreas{
				Next:     "https://pokeapi.co/api/v2/location-area?offset=40&limit=20",
				Previous: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
				Results: []struct {
					Name string `json:"name"`
				}{
					{"mt-coronet-1f-route-216"},
					{"mt-coronet-1f-route-211"},
					{"mt-coronet-b1f"},
					{"great-marsh-area-1"},
					{"great-marsh-area-2"},
					{"great-marsh-area-3"},
					{"great-marsh-area-4"},
					{"great-marsh-area-5"},
					{"great-marsh-area-6"},
					{"solaceon-ruins-2f"},
					{"solaceon-ruins-1f"},
					{"solaceon-ruins-b1f-a"},
					{"solaceon-ruins-b1f-b"},
					{"solaceon-ruins-b1f-c"},
					{"solaceon-ruins-b2f-a"},
					{"solaceon-ruins-b2f-b"},
					{"solaceon-ruins-b2f-c"},
					{"solaceon-ruins-b3f-a"},
					{"solaceon-ruins-b3f-b"},
					{"solaceon-ruins-b3f-c"},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := NewClient()
			got, err := client.GetLocationAreas(tc.url)
			if err != nil {
				t.Fatal(err)
			}
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Error(diff)
			}
		})
	}
}
