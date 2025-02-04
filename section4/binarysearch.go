package main

import "fmt"

// func binary_search(numbers []int, value int) int {
// 	left := 0
// 	right := len(numbers) - 1
// 	for left <= right {
// 		mid := (left + right) / 2
// 		if numbers[mid] == value {
// 			return mid
// 		} else if numbers[mid] < value {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
// 	return -1
// }

// 再帰的な関数
func binary_search(numbers []int, value int) int {
	return _binary_search(numbers, value, 0, len(numbers))
}

func _binary_search(numbers []int, value int, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if numbers[mid] == value {
		return mid
	} else if numbers[mid] < value {
		return _binary_search(numbers, value, mid+1, right)
	} else {
		return _binary_search(numbers, value, left, mid-1)
	}
}

func main() {
	nums := []int{0, 1, 5, 7, 9, 11, 15, 20, 24}
	fmt.Print(binary_search(nums, 20))
}
