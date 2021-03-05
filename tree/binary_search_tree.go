package tree

import "fmt"

type BinarySearchTree struct {
	Root  int
	Left  *BinarySearchTree
	Right *BinarySearchTree
}

func NewBinarySearchTree(root int) *BinarySearchTree {
	return &BinarySearchTree{
		Root:  root,
		Left:  nil,
		Right: nil,
	}
}

func (bt *BinarySearchTree) String() string {
	return fmt.Sprintf("I am a binary tree with root %d", bt.Root)
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
