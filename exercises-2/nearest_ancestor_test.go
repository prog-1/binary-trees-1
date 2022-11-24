package main

import (
	"testing"
)

func fullTree(depth int) *node {
	if depth <= 0 {
		return nil
	}
	return &node{left: fullTree(depth - 1), right: fullTree(depth - 1)}
}

func tree(depth int) func() *node {
	return func() *node { return fullTree(depth) }
}

func root(t *node) *node { return t }

func child(dirs ...func(*node) *node) func(*node) *node {
	return func(t *node) *node {
		for _, d := range dirs {
			t = d(t)
		}
		return t
	}
}

func left(t *node) *node  { return t.left }
func right(t *node) *node { return t.right }

func TestNearestAncestor(t *testing.T) {
	for _, tc := range []struct {
		name string
		t    func() *node
		a, b func(t *node) *node
		want func(t *node) *node
	}{
		{name: "empty", t: tree(0),
			a:    root,
			b:    root,
			want: root},
		{name: "single", t: tree(1),
			a:    root,
			b:    root,
			want: root},
		{name: "direct parent 1", t: tree(2),
			a:    child(left),
			b:    child(right),
			want: root},
		{name: "direct parent 2", t: tree(3),
			a:    child(left, left),
			b:    child(left, right),
			want: child(left)},
		{name: "direct parent 3", t: tree(3),
			a:    child(right, right),
			b:    child(right, left),
			want: child(right)},
		{name: "direct parent 4", t: tree(4),
			a:    child(right, right, left),
			b:    child(right, right, right),
			want: child(right, right)},
		{name: "direct parent 5", t: tree(5),
			a:    child(right, right, left, right),
			b:    child(left, right, left, right),
			want: child(root)},
		{name: "direct parent 6", t: tree(5),
			a:    child(right, right, left, right),
			b:    child(right, left, left, right),
			want: child(right)},
		{name: "direct parent 7", t: tree(5),
			a:    child(right, right, left, right),
			b:    child(right, right, right, right),
			want: child(right, right)},
		{name: "direct parent 8", t: tree(5),
			a:    child(right, right, left, right),
			b:    child(right, right, left, left),
			want: child(right, right, left)},
		{name: "direct parent 9", t: tree(5),
			a:    child(right, right, left, right),
			b:    child(right, right, left, right),
			want: child(right, right, left, right)},
		{name: "direct parent 10", t: tree(8),
			a:    child(right, right, left, right, left, left, right),
			b:    child(right, right, left, left),
			want: child(right, right, left)},
		{name: "direct parent 11", t: tree(9),
			a:    child(right, right, left, left, left, left, right),
			b:    child(right, right, left, left, right, left, right, left),
			want: child(right, right, left, left)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := tc.t()
			if got := NearestAncestor(root, tc.a(root), tc.b(root)); got != tc.want(root) {
				t.Error("returned unexpected node")
			}
		})
	}
}
