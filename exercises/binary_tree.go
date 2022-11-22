package main

import "fmt"

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
	if root.left != nil {
		inorderTraversal(root.left, sink)
	}
	sink(root.val)
	if root.right != nil {
		inorderTraversal(root.right, sink)
	}
}

// equal checks if two trees are equivalent.
func equal(a, b *node) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b == nil || a == nil && b != nil || a.val != b.val {
		return false
	}
	// if a.left == nil && b.left != nil || a.left != nil && b.left == nil || a != nil && b != nil && !equal(a.left, b.left) {
	// 	return false
	// }
	// if a.right == nil && b.right != nil || a.right != nil && b.right == nil || a != nil && b != nil && !equal(a.right, b.right) {
	// 	return false
	// }
	return equal(a.left, b.left) && equal(a.right, b.right)
	//inorderTraversal(a, func(v int) { aNode.val = v })
	//inorderTraversal(b, func(v int) { bNode.val = v })
}

// invertTree inverts a tree, so that the left and the right subtrees are swapped.
func invertTree(root *node) {
	//|| root.left == nil && root.right == nil
	if root == nil {
		return
	}
	root.left, root.right = root.right, root.left
	invertTree(root.left)
	invertTree(root.right)
}

func main() {
	//tree := tree1()
	inorderTraversal(tree2(), func(v int) { fmt.Print(v, " ") })
	//fmt.Println(equal(tree1(), tree1()))
	//fmt.Println()
	//invertTree(tree)
	//inorderTraversal(tree, func(v int) { fmt.Print(v, " ") })
}

/*
Tree 1
        _______1______
       /              \
    ___2__          ___3__
   /      \        /      \
  _4_      _5      6        7
 /   \    /
 8   9   10

*/

func tree1() *node {
	return &node{val: 1,
		left: &node{val: 2,
			left: &node{val: 4,
				left:  &node{val: 8},
				right: &node{val: 9}},
			right: &node{val: 5,
				left: &node{val: 10}}},
		right: &node{val: 3,
			left:  &node{val: 6},
			right: &node{val: 7}}}
}

/*
Tree 2
        _______4______
       /              \
    ___2__          ___6__
   /      \        /      \
   1      3       5        7

*/

func tree2() *node {
	return &node{val: 4, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 6, left: &node{val: 5}, right: &node{val: 7}}}
}

/*
Tree 3
        _______4______
       /              \
    ___2__          ___6
   /      \        /
   1      3       5

*/

func tree3() *node {
	return &node{val: 4, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 6, left: &node{val: 5}}}
}

/*
Tree 4
         1

*/

func tree4() *node {
	return &node{val: 1}
}

/*
Tree 5
        nil

*/

func tree5() *node {
	return &node{}
}

func tree1inv() *node {
	return &node{val: 1,
		left: &node{val: 3,
			left:  &node{val: 7},
			right: &node{val: 6}},
		right: &node{val: 2,
			left: &node{val: 5,
				right: &node{val: 10}},
			right: &node{val: 4,
				left:  &node{val: 9},
				right: &node{val: 8}}}}
}

func tree2inv() *node {
	return &node{val: 4, left: &node{val: 6, left: &node{val: 7}, right: &node{val: 5}}, right: &node{val: 2, left: &node{val: 3}, right: &node{val: 1}}}
}

func tree3inv() *node {
	return &node{val: 4, left: &node{val: 6, right: &node{val: 5}}, right: &node{val: 2, left: &node{val: 3}, right: &node{val: 1}}}
}

