package tree

import (
	"errors"
	"fmt"
	"sync"
)

// A BinarySearchTree is a binary tree with the following properties:
//
// - The left subtree of a node contains only nodes with keys lesser than the node’s key.
//
// - The right subtree of a node contains only nodes with keys greater than the node’s key.
//
// - The left and right subtree each must also be a binary search tree.
//
// In this implementation, BinarySearchTrees are not allowed duplicate values.
type BinarySearchTree struct {
	root *Node
	mu   sync.RWMutex // Synchronize access to the BST to avoid race conditions.
}

// NewBinarySearchTree initializes a BinarySearchTree with a nil root.
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
}

// A Node contains an integer value and has at most two children.
type Node struct {
	val         int
	left, right *Node
}

// NewNode initializes a Node with the given value and nil children.
func NewNode(val int) *Node {
	return &Node{
		val:   val,
		left:  nil,
		right: nil,
	}
}

// Insert adds a node in the proper place within the BinarySearchTree.
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
		if root.left == nil {
			root.left = insert
		} else {
			traverseAndInsert(root.left, insert)
		}
	} else if insert.val > root.val {
		if root.right == nil {
			root.right = insert
		} else {
			traverseAndInsert(root.right, insert)
		}
	} else { // Prevent the tree from inserting duplicate values.
		return errors.New(fmt.Sprintf("The value %d is already in the binary search tree!", insert.val))
	}

	return nil
}

// InOrder returns a list of integers resulting from in-order traversal of the BinarySearchTree.
func (bst *BinarySearchTree) InOrder() []int {
	bst.mu.RLock()
	defer bst.mu.RUnlock()

	visited := []int{}
	traverseInOrder(&visited, bst.root)
	return visited
}

func traverseInOrder(visited *[]int, root *Node) {
	if root != nil {
		traverseInOrder(visited, root.left)
		*visited = append(*visited, root.val)
		traverseInOrder(visited, root.right)
	}
}

// PreOrder returns a list of integers resulting from pre-order traversal of the BinarySearchTree.
func (bst *BinarySearchTree) PreOrder() []int {
	bst.mu.RLock()
	defer bst.mu.RUnlock()

	visited := []int{}
	traversePreOrder(&visited, bst.root)
	return visited
}

func traversePreOrder(visited *[]int, root *Node) {
	if root != nil {
		*visited = append(*visited, root.val)
		traversePreOrder(visited, root.left)
		traversePreOrder(visited, root.right)
	}
}

// PostOrder returns a list of integers resulting from post-order traversal of the BinarySearchTree.
func (bst *BinarySearchTree) PostOrder() []int {
	bst.mu.RLock()
	defer bst.mu.RUnlock()

	visited := []int{}
	traversePostOrder(&visited, bst.root)
	return visited
}

func traversePostOrder(visited *[]int, root *Node) {
	if root != nil {
		traversePostOrder(visited, root.left)
		traversePostOrder(visited, root.right)
		*visited = append(*visited, root.val)
	}
}
