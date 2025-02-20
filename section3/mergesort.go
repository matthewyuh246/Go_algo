package main

import "fmt"

func mergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	center := len(numbers) / 2
	left := mergeSort(numbers[:center])
	right := mergeSort(numbers[center:])

	merged := make([]int, 0, len(numbers))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}

func main() {
	numbers := []int{5, 4, 1, 8, 7, 3, 2, 9}
	sortNum := mergeSort(numbers)
	fmt.Println(sortNum)
}
