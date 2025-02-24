package main

import (
	"fmt"
	"strconv"
	"strings"
)

func generatePascalTriangle(depth int) [][]int {
	data := make([][]int, depth)
	for i := 0; i < depth; i++ {
		data[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			data[i][j] = 1
		}
	}

	for i := 2; i < depth; i++ {
		for j := 1; j < i; j++ {
			data[i][j] = data[i-1][j-1] + data[i-1][j]
		}
	}
	return data
}

func centerString(s string, width int, fill string) string {
	if len(s) >= width {
		return s
	}
	totalPadding := width - len(s)
	leftPadding := totalPadding / 2
	rightPadding := totalPadding - leftPadding
	return strings.Repeat(fill, leftPadding) + s + strings.Repeat(fill, rightPadding)
}

func printPascal(data [][]int) {
	lastRow := data[len(data)-1]
	maxVal := lastRow[0]
	for _, v := range lastRow {
		if v > maxVal {
			maxVal = v
		}
	}
	maxDigit := len(strconv.Itoa(maxVal))
	width := maxDigit + (maxDigit % 2) + 2
	depth := len(data)
	for index, line := range data {
		indentSpaces := (depth - index) * (width / 2)
		indent := strings.Repeat(" ", indentSpaces)
		var numbers []string
		for _, num := range line {
			numStr := strconv.Itoa(num)
			numbers = append(numbers, centerString(numStr, width, " "))
		}
		fmt.Println(indent + strings.Join(numbers, ""))
	}
}

func main() {
	triangle := generatePascalTriangle(10)
	printPascal(triangle)
}
