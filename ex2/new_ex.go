package main

type node[T any] struct {
	val         T
	left, right *node[T]
}

func n[T any](v T) *node[T]              { return &node[T]{val: v} }
func (n *node[T]) l(l *node[T]) *node[T] { n.left = l; return n }
func (n *node[T]) r(r *node[T]) *node[T] { n.right = r; return n }

type nodeInt = node[int]

func path(root *node[int], p *node[int], k *[]*node[int]) bool {
	if root == nil {
		return false
	}
	*k = append(*k, root)
	if root == p || root.left != nil && path(root.left, p, k) || root.right != nil && path(root.right, p, k) {
		return true
	}
	(*k) = (*k)[:len(*k)-1]
	return false
}

func nca(root, a, b *node[int]) *node[int] {
	var a1, b1 []*node[int]
	i := 0
	if !path(root, a, &a1) || !path(root, b, &b1) {
		return nil
	}
	for ; i < len(a1) && a1[i] == b1[i]; i++ {
	}
	return a1[i-1]

}

func Tree() *nodeInt {
	return n(365).l(n(52).l(n(78)).r(n(69))).r(n(100).l(n(420)).r(n(999)))
}

func main() {

}
