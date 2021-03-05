package main

import (
	"ayoconnect-golang-challenge/robber"
	"ayoconnect-golang-challenge/tree"
	"fmt"
)

func main() {
	binaryTree := tree.NewBinaryTree(5)
	fmt.Println(binaryTree)
	binaryTree.Insert()
	binaryTree.InOrder()
	binaryTree.PreOrder()
	binaryTree.PostOrder()

	fmt.Println("")

	robber.HouseRobber()
}
