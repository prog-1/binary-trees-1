package main

import "fmt"

type node[T any] struct {
	v           T
	left, right *node[T]
}

// Outputs the way from target to root
// *Return value is for internal use only
func getPath[T any](root, target *node[T], sink func(n *node[T])) bool {
	if root == nil {
		return false
	}
	if root == target || getPath(root.left, target, sink) || getPath(root.right, target, sink) {
		sink(root)
		return true
	}
	return false
}

// NCA stands for Nearest Common Ancestor
func NCA[T any](root, n1, n2 *node[T]) *node[T] {
	// traverse tree to n1 and save node path in slice s1
	// traverse tree to n2 and save node path in slice s2
	// compare these slices and return last common node

	return nil
}

func n[T any](v T) *node[T]              { return &node[T]{v: v} }
func (n *node[T]) l(l *node[T]) *node[T] { n.left = l; return n }
func (n *node[T]) r(r *node[T]) *node[T] { n.right = r; return n }

func main() {

	it := n(10).
		l(n(5).
			l(n(2)).
			r(n(7))).
		r(n(15))

	var s []*node[int]
	getPath(it, it.left.right, func(n *node[int]) { s = append(s, n) })
	for _, el := range s {
		fmt.Println(el.v)
	}
}
