package main

import (
	// "fmt"
	"math"
)

type Pair struct {
	X int
	Y int
}

type TaxiCabResult struct {
	Answer int
	Pairs  []Pair
}

func taxiCabNumber(maxAnswerNum int, matchAnswerNum int, startAnswer int) []TaxiCabResult {
	results := []TaxiCabResult{}
	gotAnswerCount := 0
	answer := startAnswer

	for gotAnswerCount < maxAnswerNum {
		matchAnswerCount := 0
		pairs := []Pair{}

		maxParam := int(math.Pow(float64(answer), 1.0/3.0)) + 1
		for x := 1; x < maxParam; x++ {
			for y := x + 1; y < maxParam; y++ {
				if x*x*x+y*y*y == answer {
					matchAnswerCount++
					pairs = append(pairs, Pair{X: x, Y: y})
				}
			}
		}

		if matchAnswerCount == matchAnswerNum {
			results = append(results, TaxiCabResult{Answer: answer, Pairs: pairs})
			gotAnswerCount++
		}
		answer++
	}

	return results
}

// func main() {
// 	results := taxiCabNumber(2, 1, 1)

// 	for _, res := range results {
// 		fmt.Printf("Answer: %d, Pairs: %v\n", res.Answer, res.Pairs)
// 	}
// }
