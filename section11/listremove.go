package main

import "fmt"

func countFrequencies(arr []int) map[int]int {
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}
	return freq
}

func minCountRemove(x *[]int, y *[]int) {
	counterX := countFrequencies(*x)
	counterY := countFrequencies(*y)

	newX := []int{}
	newY := []int{}

	for _, num := range *x {
		if counterY[num] == 0 || counterX[num] >= counterY[num] {
			newX = append(newX, num)
		}
	}

	for _, num := range *y {
		if counterX[num] == 0 || counterY[num] >= counterX[num] {
			newY = append(newY, num)
		}
	}

	*x = newX
	*y = newY
}

func main() {
	x := []int{1, 2, 3, 4, 4, 5, 5, 8, 10}
	y := []int{4, 5, 5, 5, 6, 7, 8, 8, 10}

	fmt.Println("Before:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)

	minCountRemove(&x, &y)

	fmt.Println("After:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
}
