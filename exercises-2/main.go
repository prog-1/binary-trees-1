package main

type node struct {
	val   int
	left  *node
	right *node
}

func findNearestCommonAncestor(root, a, b *node) *node {
	var aPath, bPath []*node
	findPath(root, a, &aPath)
	findPath(root, b, &bPath)
	var i int
	for i < len(aPath) && i < len(bPath) && aPath[i] == bPath[i] {
		i++
	}
	return aPath[i-1]
}

func findPath(root, x *node, path *[]*node) bool {
	*path = append(*path, root)
	if root == x || (root.left != nil && findPath(root.left, x, path)) || (root.right != nil && findPath(root.right, x, path)) {
		return true
	}
	*path = (*path)[:len(*path)-1]
	return false
}
