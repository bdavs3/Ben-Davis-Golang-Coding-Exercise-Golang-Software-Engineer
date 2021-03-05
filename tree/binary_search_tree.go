package tree

import (
	"fmt"
	"sync"
)

type BinarySearchTree struct {
	root *Node
	mu   sync.RWMutex // Synchronize access to the BST to avoid race conditions.
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

func NewNode(val int) *Node {
	return &Node{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func (bt *BinarySearchTree) Insert() {
	fmt.Println("Insert")
}

func (bt *BinarySearchTree) InOrder() {
	fmt.Println("InOrder")
}

func (bt *BinarySearchTree) PreOrder() {
	fmt.Println("PreOrder")
}

func (bt *BinarySearchTree) PostOrder() {
	fmt.Println("PostOrder")
}
