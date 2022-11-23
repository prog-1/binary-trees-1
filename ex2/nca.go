package main

type node struct {
	val         int
	left, right *node
}

func nearestCommonAncestor(root *node, a, b *node) *node {
	if root == nil || root == a || root == b {
		return root
	}
	l := nearestCommonAncestor(root.left, a, b)
	r := nearestCommonAncestor(root.right, a, b)
	if l != nil && r != nil {
		return root
	}
	if l == nil {
		return r
	}

	return l
}
