// Playing with struct data. Implementing a tree
package main

import "fmt"

type tree struct {
	value int
	left, right *tree
}

// appendValues appends the elements of t to values in order and returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}


// Sort sorts values in place.
func Sort(values  []int) {
	var root *tree
	// From slice to tree
	for _,v := range values {
		root = add(root, v)
	}
	// from tree to array (ordered already)
	appendValues(values[:0], root)
}

func main() {
	array := []int{2,1,3,5}
	Sort(array)
	fmt.Println(array)

	var t *tree
	t = add(t,5)
	t = add(t,1)
	t = add(t,2)

	fmt.Println(appendValues(array[:0], t))
}