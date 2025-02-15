package main

import "fmt"

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type IDoublyLinkedList interface {
	Append(data int)
	Insert(data int)
	Remove(data int)
	Print()
	ReverseIterative()
	ReverseRecursive()
	Sort()
}

type DoublyLinkedList struct {
	Head *Node
}

func NewDoublyLinkedList() IDoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
	}
}

func (dl *DoublyLinkedList) Append(data int) {
	newNode := &Node{Data: data}
	if dl.Head == nil {
		dl.Head = newNode
		return
	}

	currentNode := dl.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode
	newNode.Prev = currentNode
}

func (dl *DoublyLinkedList) Insert(data int) {
	newNode := &Node{Data: data}
	if dl.Head == nil {
		dl.Head = newNode
		return
	}

	currentNode := dl.Head
	currentNode.Prev = newNode
	newNode.Next = currentNode
	dl.Head = newNode
}

func (dl *DoublyLinkedList) Remove(data int) {
	if dl.Head == nil {
		return
	}

	currentNode := dl.Head
	if currentNode != nil && currentNode.Data == data {
		if currentNode.Next != nil {
			currentNode = nil
			dl.Head = nil
			return
		} else {
			nextNode := currentNode.Next
			nextNode.Prev = nil
			currentNode = nil
			dl.Head = nextNode
			return
		}
	}

	for currentNode != nil && currentNode.Data != data {
		currentNode = currentNode.Next
	}

	if currentNode == nil {
		return
	}

	if currentNode.Next == nil {
		prevNode := currentNode.Prev
		prevNode.Next = nil
		currentNode = nil
		return
	} else {
		nextNode := currentNode.Next
		prevNode := currentNode.Prev
		nextNode.Prev = prevNode
		prevNode.Next = nextNode
		currentNode = nil
		return
	}
}

func (dl *DoublyLinkedList) Print() {
	currentNode := dl.Head
	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}
}

func (dl *DoublyLinkedList) ReverseIterative() {
	var previousNode *Node
	currentNode := dl.Head
	for currentNode != nil {
		previousNode = currentNode.Prev
		currentNode.Prev = currentNode.Next
		currentNode.Next = previousNode

		currentNode = currentNode.Prev
	}

	if previousNode != nil {
		dl.Head = previousNode.Prev
	}
}

func (dl *DoublyLinkedList) ReverseRecursive() {
	var _reverseRecursive func(currentNode *Node) *Node
	_reverseRecursive = func(currentNode *Node) *Node {
		if currentNode == nil {
			return nil
		}

		previousNode := currentNode.Prev
		currentNode.Prev = currentNode.Next
		currentNode.Next = previousNode

		if currentNode.Prev == nil {
			return currentNode
		}
		return _reverseRecursive(currentNode.Prev)
	}
	dl.Head = _reverseRecursive(dl.Head)
}

func (dl *DoublyLinkedList) Sort() {
	if dl.Head == nil {
		return
	}

	currentNode := dl.Head
	for currentNode.Next != nil {
		nextNode := currentNode.Next
		for nextNode != nil {
			if currentNode.Data > nextNode.Data {
				currentNode.Data, nextNode.Data = nextNode.Data, currentNode.Data
			}
			nextNode = nextNode.Next
		}
		currentNode = currentNode.Next
	}
}

func main() {
	var dl IDoublyLinkedList = NewDoublyLinkedList()
	dl.Append(1)
	dl.Append(5)
	dl.Append(2)
	dl.Append(9)
	dl.Print()
	fmt.Println("############ Sort")
	dl.Sort()
	dl.Print()
}
