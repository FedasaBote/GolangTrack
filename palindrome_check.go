package main

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	newString := ""
	for _, char := range s {
		if unicode.IsLetter(char) {
			newString += string(unicode.ToLower(char))
		}
	}

	left, right := 0, len(newString)-1
	for left < right {
		if newString[left] != newString[right] {
			return false
		}
		left++
		right--
	}

	return true
}
