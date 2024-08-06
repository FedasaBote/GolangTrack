package main

import (
	"strings"
	"unicode"
)

func CalculateFrequency(s string) map[string]int {
	newS := ""
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsSpace(char) {
			newS += string(unicode.ToLower(char))
		}
	}

	words := strings.Fields(newS)

	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}

	return frequency
}
