package main

import "fmt"

type IStack interface {
	Push(data any)
	Pop() (any, bool)
	Print()
}

type Stack struct {
	stack []any
}

func NewStack() IStack {
	return &Stack{stack: []any{}}
}

func (s *Stack) Push(data any) {
	s.stack = append(s.stack, data)
}

func (s *Stack) Pop() (any, bool) {
	if len(s.stack) == 0 {
		return nil, false
	}
	lastIndex := len(s.stack) - 1
	data := s.stack[lastIndex]
	s.stack = s.stack[:lastIndex]
	return data, true
}

func (s *Stack) Print() {
	fmt.Println("Stack:", s.stack)
}

// func main() {
// 	var stack IStack = NewStack()

// 	stack.Print()

// 	stack.Push(1)
// 	stack.Print()

// 	stack.Push(2)
// 	stack.Print()

// 	value, _ := stack.Pop()
// 	fmt.Println("Popped:", value)
// 	stack.Print()
// }
