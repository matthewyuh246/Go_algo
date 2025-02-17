package main

import (
	"fmt"
	"strings"
)

var numAlphabetMapping = map[int]string{
	0: "+",
	1: "@",
	2: "ABC",
	3: "DEF",
	4: "GHI",
	5: "JKL",
	6: "MNO",
	7: "PQRS",
	8: "TUV",
	9: "WXYZ",
}

func phoneMnemonic(phoneNumber string) []string {
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	digits := make([]int, len(phoneNumber))
	for i, ch := range phoneNumber {
		digits[i] = int(ch - '0')
	}

	var candidate []string
	stack := []string{""}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(current) == len(digits) {
			candidate = append(candidate, current)
		} else {
			nextIndex := len(current)
			digit := digits[nextIndex]
			for _, ch := range numAlphabetMapping[digit] {
				newCandidate := current + string(ch)
				stack = append(stack, newCandidate)
			}
		}
	}

	return candidate
}

func main() {
	candidates := phoneMnemonic("568-379-8466")
	for _, s := range candidates {
		if strings.Contains(s, "LOVEPYTHON") {
			fmt.Println(s)
		}
	}
}
