package main

type node[T any] struct {
	val         T
	left, right *node[T]
}

func n[T any](v T) *node[T]              { return &node[T]{val: v} }
func (n *node[T]) l(l *node[T]) *node[T] { n.left = l; return n }
func (n *node[T]) r(r *node[T]) *node[T] { n.right = r; return n }

// NCA stands for Nearest Common Ancestor
func NCA[T any](root, n1, n2 *node[T]) *node[T] {
	if root == nil || n1 == nil || n2 == nil {
		return nil
	}
	// 1. traverse tree to n1 and save node path in slice s1
	// 2. traverse tree to n2 and save node path in slice s2
	// 3. compare these slices and return last common node

	// Outputs the way from target to root
	// *Return value is for internal use only
	var getPath func(root, target *node[T], sink func(n *node[T])) bool
	getPath = func(root, target *node[T], sink func(n *node[T])) bool {
		if root == nil {
			return false
		}
		if root == target || getPath(root.left, target, sink) || getPath(root.right, target, sink) {
			sink(root)
			return true
		}
		return false
	}

	var s1, s2 []*node[T]
	getPath(root, n1, func(n *node[T]) { s1 = append(s1, n) })
	getPath(root, n2, func(n *node[T]) { s2 = append(s2, n) })

	for i := 1; ; i++ {
		if len(s1)-i < 0 {
			return s1[0]
		} else if len(s2)-i < 0 {
			return s2[0]
		}
		if s1[len(s1)-i] != s2[len(s2)-i] {
			return s1[len(s1)-i+1]
		}
	}
}

func main() {

}
