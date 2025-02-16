package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

func maxNum(numbers []int) int {
	max := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max
}

func countingSort(numbers []int) []int {
	n := len(numbers)
	max := maxNum(numbers)
	countIndex := max + 1
	counts := make([]int, countIndex)
	result := make([]int, n)

	// 各数値の出現回数をカウントする
	for _, num := range numbers {
		counts[num] += 1
	}

	// 累積カウントを計算する
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	// ソートされた結果を生成する
	for i := n - 1; i >= 0; i-- {
		index := numbers[i]
		result[counts[index]-1] = numbers[i]
		counts[index] -= 1
	}
	return result
}

// func main() {
// 	numbers := make([]int, 10)
// 	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	for i := range numbers {
// 		numbers[i] = rng.Intn(1001)
// 	}
// 	fmt.Println(countingSort(numbers))
// }
