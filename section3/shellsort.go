package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

func shellSort(numbers []int) []int {
	lenNumbers := len(numbers)
	gap := lenNumbers / 2
	for gap > 0 {
		for i := gap; i < lenNumbers; i++ {
			tmp := numbers[i]
			j := i
			for j >= gap && numbers[j-gap] > tmp {
				numbers[j] = numbers[j-gap]
				j -= gap
			}
			numbers[j] = tmp
		}
		gap /= 2
	}

	return numbers
}

// func main() {
// 	numbers := make([]int, 10)
// 	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	for i := range numbers {
// 		numbers[i] = rng.Intn(1001)
// 	}
// 	fmt.Println(shellSort(numbers))
// }
