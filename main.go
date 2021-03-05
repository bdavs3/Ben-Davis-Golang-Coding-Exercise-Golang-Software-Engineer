package main

import (
	"ayoconnect-golang-challenge/robber"
	"ayoconnect-golang-challenge/tree"
	"fmt"
)

func main() {
	bst := tree.NewBinarySearchTree()

	err := bst.Insert(5)
	if err != nil {
		fmt.Println(err)
	}

	err = bst.Insert(4)
	if err != nil {
		fmt.Println(err)
	}

	bst.InOrder()
	bst.PreOrder()
	bst.PostOrder()

	fmt.Println("")

	robber.HouseRobber()
}
