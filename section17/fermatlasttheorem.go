package main

import (
	"fmt"
	"math"
	"strconv"
)

type Pair3 struct {
	X int
	Y int
	Z int
}

func intPow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

func fermatLastTheoremV1(maxNum int, squireNum int) []Pair3 {
	var result []Pair3
	if squireNum < 2 {
		return result
	}

	maxZ := int(math.Pow(float64((maxNum-1)*(maxNum-1)+maxNum*maxNum), 1.0/float64(squireNum)))
	for x := 1; x < maxNum+1; x++ {
		for y := x + 1; y < maxNum+1; y++ {
			for z := y + 1; z < maxZ; z++ {
				if intPow(x, squireNum)+intPow(y, squireNum) == intPow(z, squireNum) {
					result = append(result, Pair3{X: x, Y: y, Z: z})
				}
			}
		}
	}
	return result
}

func fermatLastTheoremV2(maxNum int, squireNum int) []Pair3 {
	var result []Pair3
	if squireNum < 2 {
		return result
	}

	maxInt := int(^uint(0) >> 1)

	for x := 1; x < maxNum+1; x++ {
		for y := x + 1; y < maxNum+1; y++ {
			powSum := intPow(x, squireNum) + intPow(y, squireNum)

			if powSum > maxInt {
				panic("Overflow: " + strconv.Itoa(x) + ", " + strconv.Itoa(y) + ", squireNum=" + strconv.Itoa(squireNum) + ", powSum=" + strconv.Itoa(powSum))
			}

			zFloat := math.Pow(float64(powSum), 1.0/float64(squireNum))
			if math.Floor(zFloat) != zFloat {
				continue
			}
			z := int(zFloat)
			if intPow(z, squireNum) == powSum {
				result = append(result, Pair3{X: x, Y: y, Z: z})
			}
		}
	}
	return result
}

func main() {
	fmt.Println("v1", fermatLastTheoremV1(10, 2))
	fmt.Println("v2", fermatLastTheoremV2(10, 2))
}
