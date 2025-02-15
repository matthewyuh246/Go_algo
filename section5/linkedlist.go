package main

// import "fmt"

// // ILinkedList インターフェース
// type ILinkedList interface {
// 	Append(data int)
// 	Insert(data int)
// 	Remove(data int)
// 	Print()
// 	ReverseIterative()
// 	ReverseRecursive()
// 	ReverseEven()
// }

// // Node 構造体 (連結リストのノード)
// type Node struct {
// 	Data int
// 	Next *Node
// }

// // LinkedList 構造体 (単方向連結リスト)
// type LinkedList struct {
// 	Head *Node
// }

// // NewLinkedList は新しい LinkedList を作成
// func NewLinkedList() ILinkedList {
// 	return &LinkedList{Head: nil}
// }

// // Append は末尾にノードを追加
// func (l *LinkedList) Append(data int) {
// 	newNode := &Node{Data: data}

// 	if l.Head == nil {
// 		l.Head = newNode
// 		return
// 	}

// 	lastNode := l.Head
// 	for lastNode.Next != nil {
// 		lastNode = lastNode.Next
// 	}
// 	lastNode.Next = newNode
// }

// // Insert は先頭にノードを挿入
// func (l *LinkedList) Insert(data int) {
// 	newNode := &Node{Data: data, Next: l.Head}
// 	l.Head = newNode
// }

// // Remove は指定した値のノードを削除
// func (l *LinkedList) Remove(data int) {
// 	currentNode := l.Head

// 	if currentNode != nil && currentNode.Data == data {
// 		l.Head = currentNode.Next
// 		return
// 	}

// 	var previousNode *Node
// 	for currentNode != nil && currentNode.Data != data {
// 		previousNode = currentNode
// 		currentNode = currentNode.Next
// 	}

// 	if currentNode == nil {
// 		return
// 	}

// 	previousNode.Next = currentNode.Next
// }

// // Print はリストを出力
// func (l *LinkedList) Print() {
// 	currentNode := l.Head
// 	for currentNode != nil {
// 		fmt.Println(currentNode.Data)
// 		currentNode = currentNode.Next
// 	}
// }

// // ReverseIterative はリストを反転 (イテレーティブ)
// func (l *LinkedList) ReverseIterative() {
// 	var previousNode *Node
// 	currentNode := l.Head

// 	for currentNode != nil {
// 		nextNode := currentNode.Next
// 		currentNode.Next = previousNode
// 		previousNode = currentNode
// 		currentNode = nextNode
// 	}
// 	l.Head = previousNode
// }

// // ReverseRecursive はリストを反転 (再帰的)
// func (l *LinkedList) ReverseRecursive() {
// 	l.Head = l.reverseRecursiveHelper(l.Head, nil)
// }

// // reverseRecursiveHelper は再帰的にリストを反転
// func (l *LinkedList) reverseRecursiveHelper(currentNode, previousNode *Node) *Node {
// 	if currentNode == nil {
// 		return previousNode
// 	}
// 	nextNode := currentNode.Next
// 	currentNode.Next = previousNode
// 	return l.reverseRecursiveHelper(nextNode, currentNode)
// }

// func (l *LinkedList) ReverseEven() {
// 	l.Head = l.reverseEvenHelper(l.Head, nil)
// }

// func (l *LinkedList) reverseEvenHelper(head, prev *Node) *Node {
// 	if head == nil {
// 		return nil
// 	}

// 	current := head
// 	var next *Node

// 	for current != nil && current.Data%2 == 0 {
// 		next = current.Next
// 		current.Next = prev
// 		prev = current
// 		current = next
// 	}

// 	if current != head {
// 		head.Next = current
// 		l.reverseEvenHelper(current, nil)
// 		return prev
// 	} else {
// 		head.Next = l.reverseEvenHelper(head.Next, head)
// 		return head
// 	}
// }

// func main() {
// 	// ILinkedList インターフェースとして LinkedList を利用
// 	var l ILinkedList = NewLinkedList()

// 	l.Append(1)
// 	l.Append(2)
// 	l.Append(4)
// 	l.Append(6)
// 	l.Append(3)
// 	l.Append(8)
// 	l.Append(10)
// 	l.Append(12)
// 	l.Append(5)

// 	fmt.Println("Original List:")
// 	l.Print()

// 	fmt.Println("######################")
// 	l.ReverseEven()
// 	fmt.Println("After Iterative Reverse:")
// 	l.Print()

// 	fmt.Println("#######################")
// 	l.ReverseRecursive()
// 	fmt.Println("After Recursive Reverse:")
// 	l.Print()
// }
