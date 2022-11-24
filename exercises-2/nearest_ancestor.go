package main

type node struct {
	// We don't really care about values here.
	// Unless you want to print them out.
	left, right *node
}

func FindPath(root, el *node, path []*node) (bool, []*node) {
	if root == el {
		return true, path
	}
	if root.left != nil {
		if ok, path := FindPath(root.left, el, path); ok {
			return true, append(path, root.left)
		}
	}
	if root.right != nil {
		if ok, path := FindPath(root.right, el, path); ok {
			return true, append(path, root.right)
		}
	}
	return false, path
}

func NearestAncestor(root, a, b *node) *node {
	_, path1 := FindPath(root, a, nil)
	_, path2 := FindPath(root, b, nil)
	for i, j := len(path1)-1, len(path2)-1; i >= 0 && j >= 0; i, j = i-1, j-1 { // because the slice contains the reversed version of the actual path
		if path1[i] == path2[j] {
			root = path1[i]
		}
	}
	return root
}

func main() {

}
