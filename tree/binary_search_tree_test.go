package tree

import (
	"testing"
)

func TestBSTInsert(t *testing.T) {
	tests := []struct {
		comment string
		values  []int
		want    bool // True if insert results in an error, false otherwise.
	}{
		{
			comment: "Valid input (1)",
			values:  []int{1},
			want:    false,
		},
		{
			comment: "Valid input (1, 2, 3, 4, 5)",
			values:  []int{1},
			want:    false,
		},
		{
			comment: "Valid input (71, 4, 54, 1, 99)",
			values:  []int{71, 4, 54, 1, 99},
			want:    false,
		},
		{
			comment: "Duplicate input",
			values:  []int{1, 2, 3, 1},
			want:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.comment, func(t *testing.T) {
			bst := NewBinarySearchTree()
			got := false

			for _, v := range test.values {
				err := bst.Insert(v)
				if err != nil {
					got = true
				}
			}

			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestBSTTraversal(t *testing.T) {
	bst1 := NewBinarySearchTree()
	values1 := []int{1, 2, 3, 4, 5}
	for _, i := range values1 {
		bst1.Insert(i)
	}

	bst2 := NewBinarySearchTree()
	values2 := []int{71, 8, 1, 3, -19, 4, 6}
	for _, i := range values2 {
		bst2.Insert(i)
	}

	bst3 := NewBinarySearchTree()
	values3 := []int{-10, 2, 1, 9, 0}
	for _, i := range values3 {
		bst3.Insert(i)
	}

	tests := []struct {
		comment         string
		traversalMethod string
		tree            *BinarySearchTree
		want            []int
	}{
		{
			comment:         "In-order traversal of (1, 2, 3, 4, 5)",
			traversalMethod: "in",
			tree:            bst1,
			want:            []int{1, 2, 3, 4, 5},
		},
		{
			comment:         "Pre-order traversal of (1, 2, 3, 4, 5)",
			traversalMethod: "pre",
			tree:            bst1,
			want:            []int{1, 2, 3, 4, 5},
		},
		{
			comment:         "Post-order traversal of (1, 2, 3, 4, 5)",
			traversalMethod: "post",
			tree:            bst1,
			want:            []int{5, 4, 3, 2, 1},
		},
		{
			comment:         "In-order traversal of (71, 8, 1, 3, -19, 4, 6)",
			traversalMethod: "in",
			tree:            bst2,
			want:            []int{-19, 1, 3, 4, 6, 8, 71},
		},
		{
			comment:         "Pre-order traversal of (71, 8, 1, 3, -19, 4, 6)",
			traversalMethod: "pre",
			tree:            bst2,
			want:            []int{71, 8, 1, -19, 3, 4, 6},
		},
		{
			comment:         "Post-order traversal of (71, 8, 1, 3, -19, 4, 6)",
			traversalMethod: "post",
			tree:            bst2,
			want:            []int{-19, 6, 4, 3, 1, 8, 71},
		},
		{
			comment:         "In-order traversal of (-10, 2, 1, 9, 0)",
			traversalMethod: "in",
			tree:            bst3,
			want:            []int{-10, 0, 1, 2, 9},
		},
		{
			comment:         "Pre-order traversal of (-10, 2, 1, 9, 0)",
			traversalMethod: "pre",
			tree:            bst3,
			want:            []int{-10, 2, 1, 0, 9},
		},
		{
			comment:         "Post-order traversal of (-10, 2, 1, 9, 0)",
			traversalMethod: "post",
			tree:            bst3,
			want:            []int{0, 1, 9, 2, -10},
		},
	}

	for _, test := range tests {
		t.Run(test.comment, func(t *testing.T) {
			var got []int

			switch test.traversalMethod {
			case "in":
				got = test.tree.InOrder()
			case "pre":
				got = test.tree.PreOrder()
			case "post":
				got = test.tree.PostOrder()
			}

			if !sliceEq(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

// sliceEq tests the equality of two integer slices.
func sliceEq(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
