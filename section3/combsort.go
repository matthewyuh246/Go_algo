package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

func comb_sort(numbers []int) []int {
	len_numbers := len(numbers)
	gap := len_numbers
	swapped := true

	for gap != 1 || swapped {
		gap = int(float64(gap) / 1.3)
		if gap < 1 {
			gap = 1
		}

		swapped = false

		for i := range len_numbers - gap {
			if numbers[i] > numbers[i+gap] {
				numbers[i], numbers[i+gap] = numbers[i+gap], numbers[i]
				swapped = true
			}
		}
	}
	return numbers
}

// func main() {
// 	numbers := make([]int, 10)
// 	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	for i := range numbers {
// 		numbers[i] = rng.Intn(1001)
// 	}
// 	fmt.Println(comb_sort(numbers))
// }
