package main

func linear_search(numbers []int, value int) int {
	for i, kv := range numbers {
		if kv == value {
			return i
		}
	}
	return -1
}

// func main() {
// 	nums := []int{0, 1, 5, 7, 9, 11, 15, 20, 24}
// 	fmt.Print(linear_search(nums, 2))
// }
