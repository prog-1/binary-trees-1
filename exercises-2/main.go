package main

type node struct {
	val         int
	left, right *node
}

func findNearestCommonAncestor(root, a, b *node) *node {
	if root == nil {
		return nil
	}
	var node *node
	if root == a || root == b {
		node = root
	}

	pathA := findNearestCommonAncestor(root.left, a, b)
	pathB := findNearestCommonAncestor(root.right, a, b)

	if node != nil {
		return node
	} else if pathA != nil && pathB != nil {
		return root
	} else if pathA != nil {
		return pathA
	}
	return pathB
}

func main() {
}
