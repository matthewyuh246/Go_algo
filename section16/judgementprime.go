package main

import (
	"fmt"
	"time"
)

func isPrimeV1(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func isPrimeV2(num int) bool {
	if num <= 1 {
		return false
	}

	i := 2
	for i*i < num {
		if num%i == 0 {
			return false
		}
		i++
	}

	return true
}

func isPrimeV3(num int) bool {
	if num <= 1 {
		return false
	}

	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	i := 3
	for i*i < num {
		if num%i == 0 {
			return false
		}
		i++
	}
	return true
}

func isPrimeV4(num int) bool {
	if num <= 1 {
		return false
	}

	if num == 2 || num == 3 {
		return true
	}

	if num%2 == 0 || num%3 == 0 {
		return false
	}

	i := 5
	for i*i < num {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	start1 := time.Now()
	fmt.Println(isPrimeV1(11))
	elapsed1 := time.Since(start1)
	fmt.Println("v1", elapsed1)

	start2 := time.Now()
	fmt.Println(isPrimeV2(11))
	elapsed2 := time.Since(start2)
	fmt.Println("v2", elapsed2)

	start3 := time.Now()
	fmt.Println(isPrimeV3(11))
	elapsed3 := time.Since(start3)
	fmt.Println("v3", elapsed3)

	start4 := time.Now()
	fmt.Println(isPrimeV4(11))
	elapsed4 := time.Since(start4)
	fmt.Println("v4", elapsed4)
}
