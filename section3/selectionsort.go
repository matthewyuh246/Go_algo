package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

func selection_sort(numbers []int) []int {
	len_numbers := len(numbers)
	for i := range len_numbers {
		min_idx := i
		for j := i + 1; j < len_numbers; j++ {
			if numbers[min_idx] > numbers[j] {
				min_idx = j
			}
		}
		numbers[i], numbers[min_idx] = numbers[min_idx], numbers[i]
	}
	return numbers
}

// func main() {
// 	numbers := make([]int, 10)
// 	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	for i := range numbers {
// 		numbers[i] = rng.Intn(1001)
// 	}
// 	fmt.Println(selection_sort(numbers))
// }
