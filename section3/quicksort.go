package main

// import "fmt"

func partition(numbers []int, low, high int) int {
	pivot := numbers[high]
	i := low - 1

	for j := low; j < high; j++ {
		if numbers[j] <= pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	return i + 1
}

func quickSort(numbers []int) []int {
	var _quickSort func(numbers []int, low, high int)
	_quickSort = func(numbers []int, low, high int) {
		if low < high {
			partitionIndex := partition(numbers, low, high)
			_quickSort(numbers, low, partitionIndex-1)
			_quickSort(numbers, partitionIndex+1, high)
		}
	}
	_quickSort(numbers, 0, len(numbers)-1)
	return numbers
}

// func main() {
// 	nums := []int{8, 1, 4, 2, 9, 8, 3}
// 	sortedNums := quickSort(nums)
// 	fmt.Println(sortedNums)
// }
