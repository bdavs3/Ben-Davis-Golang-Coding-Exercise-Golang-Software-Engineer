package main

import (
	"ayoconnect-golang-challenge/robber"
	"ayoconnect-golang-challenge/tree"
	"fmt"
)

func main() {
	bst := tree.NewBinarySearchTree()

	err := bst.Insert(8)
	if err != nil {
		fmt.Println(err)
	}
	bst.Insert(4)
	bst.Insert(10)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(9)
	bst.Insert(11)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(5)
	bst.Insert(7)

	fmt.Println("In-order traversal:", bst.InOrder())
	bst.PreOrder()
	bst.PostOrder()

	fmt.Println("")

	robber.HouseRobber()
}
