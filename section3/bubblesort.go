package main

func bubble_sort(numbers []int) []int {
	len_numbers := len(numbers)
	for i := range len_numbers {
		for j := range len_numbers - 1 - i {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
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
// 	fmt.Print(bubble_sort(numbers))
// }
