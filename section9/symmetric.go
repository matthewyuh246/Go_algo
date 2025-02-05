package main

import "fmt"

type Pair struct {
	First  int
	Second int
}

func findPair(pairs []Pair) []Pair {
	cache := make(map[int]int)
	var result []Pair

	for _, pair := range pairs {
		first, second := pair.First, pair.Second
		value, exists := cache[second]

		if !exists {
			cache[first] = second
		} else if value == first {
			result = append(result, pair)
		}
	}
	return result
}

func main() {
	pairs := []Pair{
		{1, 2}, {3, 5}, {4, 7}, {5, 3}, {7, 4},
	}

	for _, r := range findPair(pairs) {
		fmt.Println(r)
	}
}
