package main

type node struct {
	val   int
	left  *node
	right *node
}

// nodeCount returns the total number of nodes in a tree.
func nodeCount(root *node) int {
	if root == nil {
		return 0
	}
	return nodeCount(root.left) + nodeCount(root.right) + 1
}

// height returns the height of the tree.
func height(root *node) int {
	if root == nil || (root.left == nil && root.right == nil) {
		return 0
	}

	if height(root.left) > height(root.right) {
		return height(root.left) + 1
	} else {
		return height(root.right) + 1
	}
}

// inorderTraversal performs an inorder traversal, i.e.
// 1. visits the left subtree
// 2. visits the current node
// 3. visits the right subtree.
func inorderTraversal(root *node, sink func(v int)) {
	// TODO
}

// equal checks if two trees are equivalent.
func equal(a, b *node) bool {
	// TODO
	return false
}

// invertTree inverts a tree, so that the left and the right subtrees are swapped.
func invertTree(root *node) {
	// TODO
}

func main() {
}
