package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func generateTriangleList(depth, maxNum int) [][]int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	triangle := make([][]int, depth)
	for i := 0; i < depth; i++ {
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			row[j] = rng.Intn(maxNum + 1)
		}
		triangle[i] = row
	}
	return triangle
}

func centerString1(s string, width int, fill string) string {
	if len(s) >= width {
		return s
	}
	totalPadding := width - len(s)
	left := totalPadding / 2
	right := totalPadding - left
	return strings.Repeat(fill, left) + s + strings.Repeat(fill, right)
}

func printTriangle(data [][]int) {
	maxVal := 0
	for _, row := range data {
		for _, num := range row {
			if num > maxVal {
				maxVal = num
			}
		}
	}
	maxDigit := len(strconv.Itoa(maxVal))
	width := maxDigit + (maxDigit % 2) + 2

	depth := len(data)
	for index, row := range data {
		indentSpaces := (depth - index) * (width / 2)
		indent := strings.Repeat(" ", indentSpaces)
		var numbers []string
		for _, num := range row {
			numStr := strconv.Itoa(num)
			numbers = append(numbers, centerString1(numStr, width, " "))
		}
		fmt.Println(indent + strings.Join(numbers, ""))
	}
}

func sumMinPath(triangle [][]int) (int, bool) {
	if len(triangle) == 0 {
		return 0, false
	}

	n := len(triangle)
	treeSum := make([][]int, n)
	treeSum[0] = make([]int, len(triangle[0]))
	copy(treeSum[0], triangle[0])

	for j := 1; j < n; j++ {
		rowLen := len(triangle[j])
		treeSum[j] = make([]int, rowLen)
		for i, value := range triangle[j] {
			if i == 0 {
				treeSum[j][i] = value + treeSum[j-1][0]
			} else if i == rowLen-1 {
				treeSum[j][i] = value + treeSum[j-1][i-1]
			} else {
				left := treeSum[j-1][i-1]
				right := treeSum[j-1][i]
				if left < right {
					treeSum[j][i] = value + left
				} else {
					treeSum[j][i] = value + right
				}
			}
		}
	}
	lastRow := treeSum[n-1]
	minVal := lastRow[0]
	for _, v := range lastRow {
		if v < minVal {
			minVal = v
		}
	}
	return minVal, true
}

func main() {
	data := generateTriangleList(5, 20)
	fmt.Println("Triangle data:", data)
	printTriangle(data)
	if minPath, ok := sumMinPath(data); ok {
		fmt.Println("min path =", minPath)
	} else {
		fmt.Println("min path = nil")
	}
}
