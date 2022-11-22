package main

import "testing"

type testCase[T any] struct {
	name    string
	getData func() (root, n1, n2, want *node[T])
}

var testCases = []testCase[int]{
	{"nil", func() (root *node[int], n1 *node[int], n2 *node[int], want *node[int]) {
		return nil, nil, nil, nil
	}},
	{"standard", func() (*node[int], *node[int], *node[int], *node[int]) {
		root := n(10).
			l(n(5).
				l(n(2)).
				r(n(7))).
			r(n(15))

		return root, root.left.left, root.right, root
	}},
	{"n1 = n2", func() (*node[int], *node[int], *node[int], *node[int]) {
		root := n(10).
			l(n(5).
				l(n(2)).
				r(n(7))).
			r(n(15))

		return root, root.right, root.right, root.right
	}},
	{"all roots", func() (*node[int], *node[int], *node[int], *node[int]) {
		root := n(10).
			l(n(5).
				l(n(2)).
				r(n(7))).
			r(n(15))

		return root, root, root, root
	}},
}

func getVal[T any](n *node[T]) T {
	if n == nil {
		var res T
		return res
	}
	return n.val
}

func TestNCA(t *testing.T) {
	for _, tc := range testCases {
		root, n1, n2, want := tc.getData()
		if got := NCA(root, n1, n2); got != want {
			t.Errorf("got = %v, want = %v", getVal(got), getVal(want))
		}
	}
}
