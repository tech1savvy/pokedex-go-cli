package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"trailing spaces": {
			input: " hello world  ",
			want:  []string{"hello", "world"},
		},
		"uppercase characters": {
			input: "Charmander Bulbasaur PIKACHU",
			want:  []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Error(diff)
			}
		})
	}
}
