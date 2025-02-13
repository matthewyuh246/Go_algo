package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

func gnome_sort(numbers []int) []int {
	len_numbers := len(numbers)
	index := 0
	for index < len_numbers {
		if index == 0 {
			index += 1
		}
		if numbers[index] >= numbers[index-1] {
			index += 1
		} else {
			numbers[index], numbers[index-1] = numbers[index-1], numbers[index]
			index -= 1
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
// 	fmt.Println(gnome_sort(numbers))
// }
