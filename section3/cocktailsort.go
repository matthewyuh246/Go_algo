package main

func cocktail_sort(numbers []int) []int {
	len_numbers := len(numbers)
	swapped := true
	start := 0
	end := len_numbers - 1
	for swapped {
		swapped = false
		for i := start; i < end; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		swapped = false
		end -= 1
		for i := end; i > start; i-- {
			if numbers[i] < numbers[i-1] {
				numbers[i], numbers[i-1] = numbers[i-1], numbers[i]
				swapped = true
			}
		}

		start += 1
	}
	return numbers
}

// func main() {
// 	numbers := make([]int, 10)
// 	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	for i := range numbers {
// 		numbers[i] = rng.Intn(1001)
// 	}

// 	fmt.Print(cocktail_sort(numbers))
// }
