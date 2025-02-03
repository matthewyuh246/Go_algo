package main

import "fmt"

type IQueue interface {
	Enqueue(data any)
	Dequeue() (any, bool)
	Print()
}

type Queue struct {
	queue []any
}

func NewQueue() IQueue {
	return &Queue{
		queue: []any{},
	}
}

func (q *Queue) Enqueue(data any) {
	q.queue = append(q.queue, data)
}

func (q *Queue) Dequeue() (any, bool) {
	if len(q.queue) == 0 {
		return nil, false
	}

	firstIndex := 0
	data := q.queue[firstIndex]
	q.queue = q.queue[firstIndex+1:]
	return data, true
}

func (q *Queue) Print() {
	fmt.Println("Queue:", q.queue)
}

func main() {
	var q IQueue = NewQueue()

	q.Print()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Print()

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Print()
}
