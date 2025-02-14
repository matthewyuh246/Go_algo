package main

import (
	// "fmt"
	"math"
)

func removeZero(numbers *[]int) {
	if len(*numbers) > 0 && (*numbers)[0] == 0 {
		*numbers = (*numbers)[1:]
		removeZero(numbers)
	}
}

func listToInt(numbers []int) int {
	sumNumbers := 0
	length := len(numbers)

	for i, num := range numbers {
		exp := length - 1 - i
		sumNumbers += num * int(math.Pow(10, float64(exp)))
	}
	return sumNumbers
}

func listToIntPlusOne(numbers []int) int {
	i := len(numbers) - 1
	numbers[i] += 1

	for i > 0 {
		if numbers[i] != 10 {
			removeZero(&numbers)
			break
		}

		numbers[i] = 0
		numbers[i-1] += 1
		i--
	}

	if numbers[0] == 10 {
		numbers[0] = 1
		numbers = append(numbers, 0)
	}

	return listToInt(numbers)
}

// func main() {
// 	fmt.Println(listToIntPlusOne([]int{1, 2, 3, 9})) // 1240
// 	fmt.Println(listToIntPlusOne([]int{9, 9, 9}))    // 1000
// 	fmt.Println(listToIntPlusOne([]int{0, 1, 2}))    // 13
// }
