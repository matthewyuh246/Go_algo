package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertion_sort(numbers []int) []int {
	len_numbers := len(numbers)
	for i := 1; i < len_numbers; i++ {
		temp := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > temp {
			numbers[j+1] = numbers[j]
			j -= 1
		}

		numbers[j+1] = temp
	}
	return numbers
}

func bucket_sort(numbers []int) []int {
	if len(numbers) == 0 {
		return numbers
	}

	max_numbers := max(numbers)
	len_numbers := len(numbers)
	size := max_numbers / len_numbers

	buckets := make([][]int, size)
	for i := range buckets {
		buckets[i] = []int{}
	}

	for _, num := range numbers {
		i := num / size
		if i != size {
			buckets[i] = append(buckets[i], num)
		} else {
			buckets[size-1] = append(buckets[size-1], num)
		}
	}

	for i := range buckets {
		insertion_sort(buckets[i])
	}

	result := []int{}
	for i := range buckets {
		result = append(result, buckets[i]...)
	}
	return result
}

func max(numbers []int) int {
	maxNum := numbers[0]
	for _, num := range numbers {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func main() {
	numbers := make([]int, 10)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range numbers {
		numbers[i] = rng.Intn(1001)
	}
	fmt.Println(bucket_sort(numbers))
}
