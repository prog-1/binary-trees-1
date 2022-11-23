package main

import "testing"

func tree1() (root, a, b, nca *node) {
	a = nil
	b = nil
	nca = nil
	root = nil
	return
}
func tree2() (root, a, b, nca *node) {
	a = &node{val: 5, left: &node{val: 9, left: &node{val: 7}}, right: &node{val: 10}}
	b = &node{val: 8, left: &node{val: 123, left: &node{val: 23}, right: &node{val: 10}}, right: &node{val: 10}}
	nca = &node{val: 2, left: &node{val: 4, right: b}, right: a}
	root = &node{val: 1, left: nca, right: &node{val: 3}}
	return
}

func tree3() (root, a, b, nca *node) {
	a = &node{val: 5, left: &node{val: 9}, right: &node{val: 10}}
	b = &node{val: 8, left: &node{val: 7, left: a, right: &node{val: 7}}, right: &node{val: 3}}
	nca = b
	root = &node{val: 1, left: &node{val: 1, left: &node{val: 9}, right: nca}, right: &node{val: 3}}
	return
}

func TestNearestCommonAncestor(t *testing.T) {
	for _, tc := range []struct {
		name string
		fnc  func() (*node, *node, *node, *node)
	}{
		{"test1", tree1},
		{"test2", tree2},
		{"test3", tree3},
	} {

		t.Run(tc.name, func(t *testing.T) {
			root, a, b, want := tc.fnc()
			if got := nearestCommonAncestor(root, a, b); got != want {
				t.Errorf("got = %v, want = %v", got, want)
			}
		})
	}
}
