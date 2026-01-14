package pokeapi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetLocationArea(t *testing.T) {
	tests := map[string]struct {
		url  string
		want LocationArea
	}{
		"canalave-city-area": {
			"https://pokeapi.co/api/v2/location-area/canalave-city-area",
			LocationArea{
				"canalave-city-area",
				[]PokemonEncounters{
					{Pokemon: EncounteredPokemon{Name: "tentacool"}},
					{Pokemon: EncounteredPokemon{Name: "tentacruel"}},
					{Pokemon: EncounteredPokemon{Name: "staryu"}},
					{Pokemon: EncounteredPokemon{Name: "magikarp"}},
					{Pokemon: EncounteredPokemon{Name: "gyarados"}},
					{Pokemon: EncounteredPokemon{Name: "wingull"}},
					{Pokemon: EncounteredPokemon{Name: "pelipper"}},
					{Pokemon: EncounteredPokemon{Name: "shellos"}},
					{Pokemon: EncounteredPokemon{Name: "gastrodon"}},
					{Pokemon: EncounteredPokemon{Name: "finneon"}},
					{Pokemon: EncounteredPokemon{Name: "lumineon"}},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := NewClient()
			got, err := client.GetLocationArea(tc.url)
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
