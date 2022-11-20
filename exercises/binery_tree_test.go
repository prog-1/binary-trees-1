package main

import "testing"

func large() *node {
	return &node{val: 2,
		left: &node{val: 1},
		right: &node{val: 4,
			left:  &node{val: 3},
			right: &node{val: 5},
		}}
}

func large2() *node {
	return &node{val: 2,
		left: &node{val: 1},
		right: &node{val: 4,
			left: &node{val: 3},
			right: &node{val: 5,
				left: &node{val: 9},
			},
		}}
}

func TestNodeCount(t *testing.T) {
	for _, tc := range []struct {
		name string
		tree *node
		want int
	}{
		{"large bin tree", large(), 5},
		{"large 2 bin tree", large2(), 6},
		{"NIL tree", nil, 0},
	} {
		if got := nodeCount(tc.tree); got != tc.want {
			t.Errorf("%v; got = %v, want = %v", tc.name, got, tc.want)
		}
	}
}

func TestHeight(t *testing.T) {
	for _, tc := range []struct {
		name string
		tree *node
		want int
	}{
		{"large bin tree", large(), 2},
		{"large 2 bin tree", large2(), 3},
		{"NIL tree", nil, 0},
	} {
		if got := height(tc.tree); got != tc.want {
			t.Errorf("%v; got = %v, want = %v", tc.name, got, tc.want)
		}
	}
}
