package tree

import "fmt"

type BinaryTree struct {
	Root  int
	Left  *BinaryTree
	Right *BinaryTree
}

func NewBinaryTree(root int) *BinaryTree {
	return &BinaryTree{
		Root:  root,
		Left:  nil,
		Right: nil,
	}
}

func (bt *BinaryTree) String() string {
	return fmt.Sprintf("I am a binary tree with root %d", bt.Root)
}

func (bt *BinaryTree) Insert() {
	fmt.Println("Insert")
}

func (bt *BinaryTree) InOrder() {
	fmt.Println("InOrder")
}

func (bt *BinaryTree) PreOrder() {
	fmt.Println("PreOrder")
}

func (bt *BinaryTree) PostOrder() {
	fmt.Println("PostOrder")
}
