package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

func maxNum1(numbers []int) int {
	max := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max
}

func countingSort1(numbers []int, place int) []int {
	n := len(numbers)
	counts := make([]int, 10)
	result := make([]int, n)

	// 各数値の出現回数をカウントする
	for _, num := range numbers {
		index := int(num/place) % 10
		counts[index] += 1
	}

	// 累積カウントを計算する
	for i := 1; i < 10; i++ {
		counts[i] += counts[i-1]
	}

	// ソートされた結果を生成する
	for i := n - 1; i >= 0; i-- {
		index := int(numbers[i]/place) % 10
		result[counts[index]-1] = numbers[i]
		counts[index] -= 1
	}
	return result
}

func radixSort(numbers []int) []int {
	maxNum := maxNum1(numbers)
	place := 1
	for maxNum > place {
		numbers = countingSort1(numbers, place)
		place *= 10
	}
	return numbers
}

// func main() {
// 	numbers := make([]int, 10)
// 	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	for i := range numbers {
// 		numbers[i] = rng.Intn(1001)
// 	}
// 	fmt.Println(radixSort(numbers))
// }
