// 1. Input: [11, 2, 5, 9, 10, 3], 12 => Output: (2, 10) or None
// 2. Input: [11, 2, 5, 9, 10, 3]     => Output: (11, 9) or None ex) 11 + 9 = 2 + 5 + 10 + 3
package main

import "fmt"

func getPair(numbers []int, target int) (int, int, bool) {
	cache := make(map[int]bool)

	for _, num := range numbers {
		val := target - num
		if cache[val] {
			return val, num, true
		}
		cache[num] = true
	}
	return 0, 0, false
}

func get_pair_half_sum(numbers []int) (int, int, bool) {
	var sum_numbers, half_sum int
	sum_numbers = 0
	for _, num := range numbers {
		sum_numbers += num
	}
	if sum_numbers%2 != 0 {
		return 0, 0, false
	} else {
		half_sum = int(sum_numbers / 2)
	}

	cache := make(map[int]bool)
	for _, num := range numbers {
		val := half_sum - num
		if cache[val] {
			return val, num, true
		}
		cache[num] = true
	}
	return 0, 0, false
}

func main() {
	numbers := []int{11, 2, 5, 9, 10, 3}
	target := 12

	if val, num, found := getPair(numbers, target); found {
		fmt.Printf("Pair found: (%d, %d)\n", val, num)
	} else {
		fmt.Println("No pair found")
	}
	numbers2 := []int{11, 2, 5, 9, 10, 3}
	if val, num, found := get_pair_half_sum(numbers2); found {
		fmt.Printf("Pair found: (%d, %d)\n", val, num)
	} else {
		fmt.Println("No pair found")
	}
}
