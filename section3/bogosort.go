package main

import (
	"math/rand"
	"time"
)

func inOrder(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}
	return true
}

func bogoSort(numbers []int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for !inOrder(numbers) {
		rng.Shuffle(len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})
	}
}

// func main() {
// 	numbers := make([]int, 10)
// 	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	for i := range numbers {
// 		numbers[i] = rng.Intn(1001)
// 	}

// 	fmt.Println("Before sorting:", numbers)
// 	bogoSort(numbers)
// 	fmt.Println("After sorting:", numbers)
// }
