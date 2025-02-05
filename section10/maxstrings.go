package main

import (
	"fmt"
	"strings"
)

func countChars(s string) (rune, int) {
	s = strings.ToLower(s)
	charCount := make(map[rune]int)

	for _, char := range s {
		if char != ' ' {
			charCount[char]++
		}
	}

	var maxChar rune
	maxCount := 0
	for char, count := range charCount {
		if count > maxCount {
			maxChar = char
			maxCount = count
		}
	}

	return maxChar, maxCount
}

func main() {
	str := "This is a pen. This is an apple. Applepen."
	maxChar, maxCount := countChars(str)
	fmt.Printf("(%c, %d)", maxChar, maxCount)
}
