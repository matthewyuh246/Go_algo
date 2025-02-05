package main

import "fmt"

// IBinarySearchTree インターフェース
type IBinarySearchTree interface {
	Insert(value int)
	Search(value int) bool
	Remove(value int)
	Inorder()
}

// Node 構造体 (二分探索木のノード)
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BinarySearchTree 構造体 (二分探索木)
type BinarySearchTree struct {
	Root *Node
}

// NewBinarySearchTree は新しい BST を作成
func NewBinarySearchTree() IBinarySearchTree {
	return &BinarySearchTree{Root: nil}
}

// Insert は値を挿入
func (bst *BinarySearchTree) Insert(value int) {
	bst.Root = bst.insertRec(bst.Root, value)
}

// insertRec は再帰的にノードを挿入
func (bst *BinarySearchTree) insertRec(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}

	if value < node.Value {
		node.Left = bst.insertRec(node.Left, value)
	} else {
		node.Right = bst.insertRec(node.Right, value)
	}
	return node
}

// Search は指定した値が存在するかを探索
func (bst *BinarySearchTree) Search(value int) bool {
	return bst.searchRec(bst.Root, value)
}

// searchRec は再帰的に探索
func (bst *BinarySearchTree) searchRec(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	} else if value < node.Value {
		return bst.searchRec(node.Left, value)
	} else {
		return bst.searchRec(node.Right, value)
	}
}

// minValue は最小値のノードを見つける
func (bst *BinarySearchTree) minValue(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Remove は値を削除
func (bst *BinarySearchTree) Remove(value int) {
	bst.Root = bst.removeRec(bst.Root, value)
}

// removeRec は再帰的にノードを削除
func (bst *BinarySearchTree) removeRec(node *Node, value int) *Node {
	if node == nil {
		return node
	}

	if value < node.Value {
		node.Left = bst.removeRec(node.Left, value)
	} else if value > node.Value {
		node.Right = bst.removeRec(node.Right, value)
	} else {
		// 片方または両方が NULL の場合
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// 2つの子ノードを持つ場合
		temp := bst.minValue(node.Right)
		node.Value = temp.Value
		node.Right = bst.removeRec(node.Right, temp.Value)
	}
	return node
}

// Inorder は中順 (Inorder Traversal) を実行
func (bst *BinarySearchTree) Inorder() {
	bst.inorderRec(bst.Root)
	fmt.Println()
}

// inorderRec は再帰的に中順探索
func (bst *BinarySearchTree) inorderRec(node *Node) {
	if node != nil {
		bst.inorderRec(node.Left)
		fmt.Print(node.Value, " ")
		bst.inorderRec(node.Right)
	}
}

func main() {
	var bst IBinarySearchTree = NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(5)
	bst.Insert(7)
	bst.Insert(1)
	bst.Insert(10)
	bst.Insert(2)

	fmt.Println("Inorder Traversal:")
	bst.Inorder()

	fmt.Println("Search 6:", bst.Search(6))

	bst.Remove(6)
	fmt.Println("########### Remove 6")
	bst.Inorder()
}
