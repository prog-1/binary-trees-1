package main

import "testing"

var (
	pointer1 = &node{v: 4}
	pointer2 = &node{v: 3}
	pointer3 = &node{v: 5}
	pointer4 = &node{v: 2, left: pointer1, right: pointer3}
	pointer5 = &node{v: 1, left: pointer4, right: pointer2}
	root     = pointer5
)

func TestFindNearestCommonAncestor(t *testing.T) {
	for _, tc := range []struct {
		name       string
		a, b, want *node
	}{
		{"1", pointer5, pointer2, pointer5},
		{"2", pointer1, pointer2, pointer5},
		{"3", pointer1, pointer3, pointer4},
		{"4", pointer1, pointer4, pointer4},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := findAncestor(root, tc.a, tc.b); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
