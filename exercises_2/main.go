package main

import "fmt"

type node struct {
	val         int
	left, right *node
}

func find(t, a, b *node) *node {
	if t == nil {
		return nil
	}
	if t == a || t == b {
		return t
	}
	lefta := find(t.left, a, b)
	righta := find(t.right, a, b)

	if lefta != nil && righta != nil {
		return t
	}
	if lefta != nil {
		return lefta
	}
	return righta
}

func main() {
	a := &node{val: 4, left: &node{val: 5}, right: &node{val: 6}}
	b := &node{val: 3}
	tree := find(&node{val: 1, left: &node{val: 2, left: b, right: a}}, a, b)
	fmt.Println(tree.val)
}
