package main

import (
	// "fmt"
	"strings"
)

func snakeStringV1(chars string) []string {
	var result [3]strings.Builder

	for i, r := range chars {
		var insertIndex int
		if i%4 == 1 {
			insertIndex = 0
		} else if i%2 == 0 {
			insertIndex = 1
		} else if i%4 == 3 {
			insertIndex = 2
		}

		result[insertIndex].WriteRune(r)

		for j := 0; j < 3; j++ {
			if j != insertIndex {
				result[j].WriteRune(' ')
			}
		}
	}

	return []string{
		result[0].String(),
		result[1].String(),
		result[2].String(),
	}
}

func snakeStringV2(chars string, depth int) []string {
	builders := make([]strings.Builder, depth)

	insertIndex := depth / 2

	pos := func(i int) int { return i + 1 }

	neg := func(i int) int { return i - 1 }

	op := neg

	for _, r := range chars {
		builders[insertIndex].WriteRune(r)

		for i := 0; i < depth; i++ {
			if i != insertIndex {
				builders[i].WriteRune(' ')
			}
		}

		if insertIndex <= 0 {
			op = pos
		}

		if insertIndex >= depth-1 {
			op = neg
		}

		insertIndex = op(insertIndex)
	}

	result := make([]string, depth)
	for i, b := range builders {
		result[i] = b.String()
	}
	return result
}

// func main() {
// 	var sb strings.Builder
// 	for j := 0; j < 5; j++ {
// 		for i := 0; i < 10; i++ {
// 			sb.WriteString(strconv.Itoa(i))
// 		}
// 	}
// 	s := sb.String()

// 	resultLines := snakeStringV1(s)

// 	for _, line := range resultLines {
// 		fmt.Println(line)
// 	}

// 	alphabet := "abcdefghijklmnopqrstuvwxyz"
// 	s := alphabet + alphabet

// 	resultLines := snakeStringV2(s, 6)

// 	for _, line := range resultLines {
// 		fmt.Println(line)
// 	}
// }
