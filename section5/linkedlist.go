package main

import "fmt"

// ILinkedList インターフェース
type ILinkedList interface {
	Append(data any)
	Insert(data any)
	Remove(data any)
	Print()
	ReverseIterative()
	ReverseRecursive()
}

// Node 構造体 (連結リストのノード)
type Node struct {
	Data any
	Next *Node
}

// LinkedList 構造体 (単方向連結リスト)
type LinkedList struct {
	Head *Node
}

// NewLinkedList は新しい LinkedList を作成
func NewLinkedList() ILinkedList {
	return &LinkedList{Head: nil}
}

// Append は末尾にノードを追加
func (l *LinkedList) Append(data any) {
	newNode := &Node{Data: data}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	lastNode := l.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}
	lastNode.Next = newNode
}

// Insert は先頭にノードを挿入
func (l *LinkedList) Insert(data any) {
	newNode := &Node{Data: data, Next: l.Head}
	l.Head = newNode
}

// Remove は指定した値のノードを削除
func (l *LinkedList) Remove(data any) {
	currentNode := l.Head

	if currentNode != nil && currentNode.Data == data {
		l.Head = currentNode.Next
		return
	}

	var previousNode *Node
	for currentNode != nil && currentNode.Data != data {
		previousNode = currentNode
		currentNode = currentNode.Next
	}

	if currentNode == nil {
		return
	}

	previousNode.Next = currentNode.Next
}

// Print はリストを出力
func (l *LinkedList) Print() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}
}

// ReverseIterative はリストを反転 (イテレーティブ)
func (l *LinkedList) ReverseIterative() {
	var previousNode *Node
	currentNode := l.Head

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	l.Head = previousNode
}

// ReverseRecursive はリストを反転 (再帰的)
func (l *LinkedList) ReverseRecursive() {
	l.Head = l.reverseRecursiveHelper(l.Head, nil)
}

// reverseRecursiveHelper は再帰的にリストを反転
func (l *LinkedList) reverseRecursiveHelper(currentNode, previousNode *Node) *Node {
	if currentNode == nil {
		return previousNode
	}
	nextNode := currentNode.Next
	currentNode.Next = previousNode
	return l.reverseRecursiveHelper(nextNode, currentNode)
}

func main() {
	// ILinkedList インターフェースとして LinkedList を利用
	var l ILinkedList = NewLinkedList()

	l.Append(0)
	l.Append(1)
	l.Append(2)
	l.Append(3)

	fmt.Println("Original List:")
	l.Print()

	fmt.Println("######################")
	l.ReverseIterative()
	fmt.Println("After Iterative Reverse:")
	l.Print()

	fmt.Println("#######################")
	l.ReverseRecursive()
	fmt.Println("After Recursive Reverse:")
	l.Print()
}
