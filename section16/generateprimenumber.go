package main

// import "fmt"

func generatePrimesV1(number int) []int {
	var primes []int
	for x := 2; x < number+1; x++ {
		isPrime := true
		for y := 2; y < x; y++ {
			if x%y == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, x)
		}
	}
	return primes
}

func generatePrimesV2(number int) []int {
	primes := []int{}
	cache := make(map[int]bool)

	for x := 2; x <= number; x++ {
		if val, ok := cache[x]; ok && val == false {
			continue
		}
		primes = append(primes, x)
		cache[x] = true
		for y := x * 2; y <= number; y += x {
			cache[y] = false
		}
	}
	return primes
}

func generatePrimesV3(number int) <-chan int {
	cache := make(map[int]bool)
	ch := make(chan int)
	go func() {
		for x := 2; x < number+1; x++ {
			if val, ok := cache[x]; ok && val == false {
				continue
			}
			ch <- x
			cache[x] = true
			for y := x * 2; y < number+1; y += x {
				cache[y] = false
			}
		}
		close(ch)
	}()
	return ch

}

// func main() {
// 	nums := generatePrimesV3(50)
// 	for num := range nums {
// 		fmt.Printf("%d ", num)
// 	}
// }
