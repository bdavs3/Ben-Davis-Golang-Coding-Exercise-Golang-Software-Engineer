package tree

import (
	"errors"
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

func (bst *BinarySearchTree) Insert(val int) error {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	insert := NewNode(val)
	if bst.root == nil {
		bst.root = insert
		return nil
	} else {
		return traverseAndInsert(bst.root, insert)
	}
}

func traverseAndInsert(root, insert *Node) error {
	if insert.val < root.val {
		if insert.left == nil {
			root.left = insert
		} else {
			traverseAndInsert(root.left, insert)
		}
	} else if insert.val > root.val {
		if insert.right == nil {
			root.right = insert
		} else {
			traverseAndInsert(root.right, insert)
		}
	} else { // As a simplification, prevent the tree from containing duplicate values.
		return errors.New(fmt.Sprintf("The value %d is already in the binary search tree!", insert.val))
	}

	return nil
}

func (bst *BinarySearchTree) InOrder() {
	fmt.Println("InOrder")
}

func (bst *BinarySearchTree) PreOrder() {
	fmt.Println("PreOrder")
}

func (bst *BinarySearchTree) PostOrder() {
	fmt.Println("PostOrder")
}
