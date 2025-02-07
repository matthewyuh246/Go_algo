package main

import (
	// "fmt"
	"sync"
	// "time"
)

type IMemorize interface {
	Call(n int) int
}

type Memorize struct {
	cache map[int]int
	mu    sync.Mutex
	fn    func(int) int
}

func NewMemorize(f func(int) int) IMemorize {
	return &Memorize{
		cache: make(map[int]int),
		fn:    f,
	}
}

func (m *Memorize) Call(n int) int {
	m.mu.Lock()
	defer m.mu.Unlock()

	if result, found := m.cache[n]; found {
		return result
	}

	result := m.fn(n)
	m.cache[n] = result
	return result
}

func longFunc(num int) int {
	r := 0
	for i := 0; i < 10000000; i++ {
		r += num * i
	}
	return r
}

// func main() {
// 	var memorizedFunc IMemorize = NewMemorize(longFunc)

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(memorizedFunc.Call(i))
// 	}

// 	start := time.Now()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(memorizedFunc.Call(i))
// 	}
// 	fmt.Println("Execution time:", time.Since(start))
// }
