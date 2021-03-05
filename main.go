package main

import (
	"ayoconnect-golang-challenge/robber"
	"ayoconnect-golang-challenge/tree"
	"fmt"
)

func main() {
	bst := tree.NewBinarySearchTree()
	bst.Insert()
	bst.InOrder()
	bst.PreOrder()
	bst.PostOrder()

	fmt.Println("")

	robber.HouseRobber()
}
