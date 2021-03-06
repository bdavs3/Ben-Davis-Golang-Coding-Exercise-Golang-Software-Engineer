# Ayoconnect Golang Challenge

This is a coding challenge for the Golang Software Engineer position at [Ayoconnect](https://ayoconnect.id).

Specifications for the challenge can be found in the [Challenge](https://github.com/bdavs3/Ben-Davis-Golang-Coding-Exercise-Golang-Software-Engineer/blob/master/Challenge.pdf) document at the root of the repository.

## System Requirements

You'll need to have [Go](https://golang.org/dl/) installed on your system in order to run the code.

## Usage

This repository contains packages for two different assignments: **Binary Tree** and **Robber**. The functionality of each package is made available through the command line:

```sh
$ go run main.go tree 5 1 17 8 2 9 # Build a binary search tree with the provided values (and automatically print the possible traversals)
In-order traversal: 1 2 5 8 9 17
Pre-order traversal: 5 1 2 17 8 9
Post-order traversal: 2 1 9 8 17 5
$ go run main.go robber 1 2 3 1 # Provide house values to compute the best possible robbing route
4
```

## Testing

To run tests on the packages, use the built-in Go CLI tool:

```sh
$ go test ./... # Run each group of tests to see what passed or failed
$ go test -v ./... # Use the "verbose" flag to see descriptions of each test
```

## Assumptions

These are the main assumptions I made for this assignment:
- Instead of creating a more general **binary tree** structure, I created a **binary search tree** structure so that there would be strict rules for the placement of new nodes.
- I allowed only integers (not floats) for the binary search tree.
- I did not allow duplicate values for the binary search tree.
- For program usage, I decided to allow the user to enter custom input rather than hard-coding input given in the challenge spec. I reserved hard-coded inputs for the testing modules.