package main

import (
	// "fmt"
	"math"
)

func listnumplus(numbers []int) int {
	len_numbers := len(numbers)
	var result int
	for i := range len_numbers {
		result += numbers[i] * int(math.Pow(10, float64(len_numbers-1-i)))
	}
	result += 1
	return result
}

// func main() {
// 	numbers := []int{7, 8, 9}
// 	fmt.Println(listnumplus(numbers))
// }
