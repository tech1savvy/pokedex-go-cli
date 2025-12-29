package main

import (
	"fmt"
	"strings"
)

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

