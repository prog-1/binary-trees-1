package main

import "testing"

func tree1() (tree, a_pointer, b_pointer, nearestCommonAncestor_pointer *node) {
	a_pointer = &node{val: 5, left: &node{val: 9}, right: &node{val: 10}}
	b_pointer = &node{val: 8}
	nearestCommonAncestor_pointer = &node{val: 2, left: &node{val: 4, right: b_pointer}, right: a_pointer}
	tree = &node{val: 1, left: nearestCommonAncestor_pointer, right: &node{val: 3}}
	return
}

func tree2() (tree, a_pointer, b_pointer, nearestCommonAncestor_pointer *node) {
	a_pointer = &node{val: 5, left: &node{val: 9}, right: &node{val: 10}}
	b_pointer = &node{val: 8, left: &node{val: 7, left: a_pointer, right: &node{val: 7}}, right: &node{val: 3}}
	nearestCommonAncestor_pointer = b_pointer
	tree = &node{val: 1, left: &node{val: 1, left: &node{val: 9}, right: nearestCommonAncestor_pointer}, right: &node{val: 3}}
	return
}

func TestNearestCommonAncestor(t *testing.T) {
	for _, tc := range []struct {
		name   string
		tcFunc func() (*node, *node, *node, *node)
	}{
		{"a and b in different branches", tree1},
		{"a and b in one branch", tree2},
	} {
		tree, a, b, want := tc.tcFunc()
		if got := nearestCommonAncestor(tree, a, b); got != want {
			t.Errorf("Error in test case %v", tc.name)
		}
	}
}
