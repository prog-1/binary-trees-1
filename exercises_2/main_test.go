package main

import "testing"

var (
	tree  = &node{val: 1, left: &node{val: 2, left: &node{val: 3}, right: &node{val: 4, left: &node{val: 5}, right: &node{val: 6}}}}
	a     = tree.left.right
	b     = tree.left.left
	anc   = tree.left
	tree2 = &node{val: 1, left: &node{val: 2, left: &node{val: 3}, right: &node{val: 4}}}
	a2    = tree2.left
	b2    = tree2.left.right
	anc2  = tree2.left
)

func TestFind(t *testing.T) {
	for _, tc := range []struct {
		name          string
		t, a, b, want *node
	}{
		{"nil", nil, nil, nil, nil},
		{"different branches", tree, a, b, anc},
		{"from the same branch", tree2, a2, b2, anc2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := find(tc.t, tc.a, tc.b); got != tc.want {
				t.Errorf("TEST: %v. Got = %v, want = %v", tc.name, got, tc.want)
			}
		})
	}
}
