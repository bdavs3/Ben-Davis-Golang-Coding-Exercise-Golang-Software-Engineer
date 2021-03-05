package main

import (
	"ayoconnect-golang-challenge/robber"
	"ayoconnect-golang-challenge/tree"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	pkg := os.Args[1]
	switch pkg {
	case "tree":
		bst := tree.NewBinarySearchTree()
		values := os.Args[2:]

		for _, val := range values {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("Binary search tree may only accept integer values.")
			}

			err = bst.Insert(intVal)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println("In-order traversal:", bst.InOrder())
		fmt.Println("Pre-order traversal:", bst.PreOrder())
		fmt.Println("Post-order traversal:", bst.PostOrder())
	case "robber":
		robber.HouseRobber()
	default:
		log.Fatal("Please enter either 'tree' or 'robber' as the first program argument.")
	}
}
