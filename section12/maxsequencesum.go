package main

// import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxSequenceSum(numbers []int) int {
	resultSequence, sumSequence := 0, 0
	for _, num := range numbers {
		sumSequence = max(num, sumSequence+num)
		resultSequence = max(resultSequence, sumSequence)
	}
	return resultSequence
}

func findMaxCircularSequenceSum(numbers []int) int {
	maxSequenceSum := getMaxSequenceSum(numbers)
	var invertNumber []int
	allSum := 0
	for _, num := range numbers {
		allSum += num
		invertNumber = append(invertNumber, -num)
	}

	maxWrapSequence := allSum - (-getMaxSequenceSum(invertNumber))
	return max(maxSequenceSum, maxWrapSequence)
}

// func main() {
// 	numbers := []int{1, -2, 3, 6, -1, 2, 4, -5, 2}
// 	fmt.Println(findMaxCircularSequenceSum(numbers))
// }
