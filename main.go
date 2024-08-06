package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a text to calculate word frequency and check palindrome: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	frequency := CalculateFrequency(text)
	fmt.Println("Word Frequency:", frequency)

	isPalindrome := IsPalindrome(text)
	fmt.Println("Is Palindrome:", isPalindrome)
}
