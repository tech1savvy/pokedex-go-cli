package pokeapi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetPokemon(t *testing.T) {
	tests := map[string]struct {
		url  string
		want Pokemon
	}{
		"pikachu": {
			"https://pokeapi.co/api/v2/pokemon/pikachu",
			Pokemon{
				"pikachu",
				112,
				4,
				60,
				[]PokemonStat{
					{
						35,
						Stat{
							"hp",
						},
					},
					{
						55,
						Stat{
							"attack",
						},
					},
					{
						40,
						Stat{
							"defense",
						},
					},
					{
						50,
						Stat{
							"special-attack",
						},
					},
					{
						50,
						Stat{
							"special-defense",
						},
					},
					{
						90,
						Stat{
							"speed",
						},
					},
				},
				[]PokemonType{
					{
						Type{
							"electric",
						},
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := NewClient()
			got, err := client.GetPokemon(tc.url)
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
