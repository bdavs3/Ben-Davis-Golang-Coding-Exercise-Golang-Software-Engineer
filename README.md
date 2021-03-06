# Ayoconnect Golang Challenge

This is a coding challenge for the Golang Software Engineer position at [Ayoconnect](https://ayoconnect.id).

Specifications for the challenge can be found in the [Challenge](https://github.com/bdavs3/Ben-Davis-Golang-Coding-Exercise-Golang-Software-Engineer/blob/master/Challenge.pdf) document at the root of the repository.

## System Requirements

You'll need to have [Go](https://golang.org/dl/) installed on your system in order to run the code.

## Usage

This repository contains packages for two different assignments: **Binary Tree** and **Robber**. The functionality of each package is made available through the command line and the use of a few necessary arguments:

```sh
$ go run main.go tree 5 1 17 8 2 9 # Build a binary search tree with the provided values (and automatically print the possible traversals)
In-order traversal: [1 2 5 8 9 17]
Pre-order traversal: [5 1 2 17 8 9]
Post-order traversal: [2 1 9 8 17 5]
$ go run main.go robber 1 2 3 1 # Provide house values to compute the best possible robbing route
4
```