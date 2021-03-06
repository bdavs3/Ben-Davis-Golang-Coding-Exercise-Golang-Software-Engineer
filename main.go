package main

import (
	"ayoconnect-golang-challenge/robber"
	"ayoconnect-golang-challenge/tree"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

const usageInstructions = "Please enter either 'tree' or 'robber' as the first program argument."

func main() {
	if len(os.Args) == 1 {
		log.Fatal(usageInstructions)
	}

	pkg := os.Args[1]

	switch pkg {
	case "tree":
		bst := tree.NewBinarySearchTree()
		values := os.Args[2:]

		if len(values) == 0 {
			log.Fatal("Please supply at least 1 integer to the tree.")
		}

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

		fmt.Println("In-order traversal:", printSlice(bst.InOrder()))
		fmt.Println("Pre-order traversal:", printSlice(bst.PreOrder()))
		fmt.Println("Post-order traversal:", printSlice(bst.PostOrder()))
	case "robber":
		values := os.Args[2:]
		var houses []int

		for _, val := range values {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("Houses may only be represented by integer cash values.")
			}

			houses = append(houses, intVal)
		}

		street := robber.NewStreet(houses)
		fmt.Println(street.Rob())
	default:
		log.Fatal(usageInstructions)
	}
}

func printSlice(s []int) string {
	var b bytes.Buffer
	for _, i := range s {
		fmt.Fprint(&b, i, " ")
	}
	return b.String()
}
